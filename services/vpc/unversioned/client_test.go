package vpc

import (
	"encoding/json"
	"github.com/zqfan/tencentcloud-sdk-go/common"
	cvm "github.com/zqfan/tencentcloud-sdk-go/services/cvm/v20170312"
	"os"
	"testing"
	"time"
)

func newClient() (*Client, error) {
	return NewClientWithSecretId(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
		"ap-guangzhou",
	)
}

func TestNatGatewayCRUD(t *testing.T) {
	c, _ := newClient()
	vpcDescReq := NewDescribeVpcExRequest()
	vpcDescResp, err := c.DescribeVpcEx(vpcDescReq)
	b, _ := json.Marshal(vpcDescResp)
	t.Logf("resp=%s", b)

	cvmc, _ := cvm.NewClientWithSecretId(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
		"ap-guangzhou",
	)

	eipDescReq := cvm.NewDescribeAddressesRequest()
	eipDescReq.Filters = []*cvm.Filter{
		&cvm.Filter{
			Name:   common.StringPtr("address-status"),
			Values: []*string{common.StringPtr("UNBIND")},
		},
	}
	eipDescResp, _ := cvmc.DescribeAddresses(eipDescReq)
	b, _ = json.Marshal(eipDescResp)
	t.Logf("eip desc resp=%s", b)

	// create
	createReq := NewCreateNatGatewayRequest()
	createReq.VpcId = vpcDescResp.Data[0].UnVpcId
	createReq.NatName = common.StringPtr("nat-test-xyz")
	createReq.MaxConcurrent = common.IntPtr(1000)
	createReq.AssignedEipSet = []*string{eipDescResp.Response.AddressSet[0].AddressIp}
	createResp, err := c.CreateNatGateway(createReq)
	b, _ = json.Marshal(createResp)
	t.Logf("resp=%s", b)
	if _, ok := err.(*common.APIError); ok {
		t.Errorf("Fail err=%v", err)
		return
	}
	// retrieve
	descReq := NewDescribeNatGatewayRequest()
	descReq.NatName = common.StringPtr("nat-test-xyz")
	descResp, err := c.DescribeNatGateway(descReq)
	b, _ = json.Marshal(descResp)
	t.Logf("resp=%s", b)
	if _, ok := err.(*common.APIError); ok {
		t.Errorf("Fail err=%v, resp=%v", err, descResp)
		return
	}
	// delete
	deleteNatGateway(descResp.Data[0].UnVpcId, descResp.Data[0].NatId, t)
}

func deleteNatGateway(vpcId, natId *string, t *testing.T) {
	vpcconn, _ := NewClientWithSecretId(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
		"ap-guangzhou",
	)
	deleteReq := NewDeleteNatGatewayRequest()
	deleteReq.VpcId = vpcId
	deleteReq.NatId = natId
	for {
		deleteResp, err := vpcconn.DeleteNatGateway(deleteReq)
		b, _ := json.Marshal(deleteResp)
		t.Logf("resp=%s", b)
		if _, ok := err.(*common.APIError); ok {
			t.Errorf("Fail err=%v, resp=%v", err, deleteResp)
			return
		}
		taskReq := NewDescribeVpcTaskResultRequest()
		taskReq.TaskId = deleteResp.TaskId
		for {
			taskResp, err := vpcconn.DescribeVpcTaskResult(taskReq)
			b, _ = json.Marshal(taskResp)
			t.Logf("resp=%s", b)
			if _, ok := err.(*common.APIError); ok {
				t.Errorf("Fail err=%v, resp=%v", err, taskResp)
				return
			}
			if *taskResp.Data.Status == 0 {
				return
			} else if *taskResp.Data.Status == 1 {
				// fail, need retry delete
				break
			}
			time.Sleep(10 * time.Second)
		}
		time.Sleep(10 * time.Second)
	}
}

func TestNatGatewayBindUnbindEIP(t *testing.T) {
	c, _ := newClient()
	cvmc, _ := cvm.NewClientWithSecretId(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
		"ap-guangzhou",
	)

	vpcDescReq := NewDescribeVpcExRequest()
	vpcDescResp, _ := c.DescribeVpcEx(vpcDescReq)

	eipDescReq := cvm.NewDescribeAddressesRequest()
	eipDescReq.Filters = []*cvm.Filter{
		&cvm.Filter{
			Name:   common.StringPtr("address-status"),
			Values: []*string{common.StringPtr("UNBIND")},
		},
	}
	eipDescResp, _ := cvmc.DescribeAddresses(eipDescReq)
	b, _ := json.Marshal(eipDescResp)
	t.Logf("eip desc resp=%s", b)

	createReq := NewCreateNatGatewayRequest()
	createReq.VpcId = vpcDescResp.Data[0].UnVpcId
	createReq.NatName = common.StringPtr("nat-jngbqyfs")
	createReq.MaxConcurrent = common.IntPtr(1000)
	createReq.AssignedEipSet = []*string{eipDescResp.Response.AddressSet[0].AddressIp}
	createResp, _ := c.CreateNatGateway(createReq)
	b, _ = json.Marshal(createResp)
	t.Logf("nat create resp=%s", b)

	descReq := NewDescribeNatGatewayRequest()
	descReq.NatName = common.StringPtr("nat-jngbqyfs")
	descResp, _ := c.DescribeNatGateway(descReq)
	b, _ = json.Marshal(descResp)
	t.Logf("nat desc resp=%s", b)

	eipDescReq = cvm.NewDescribeAddressesRequest()
	eipDescReq.Filters = []*cvm.Filter{
		&cvm.Filter{
			Name:   common.StringPtr("address-status"),
			Values: []*string{common.StringPtr("UNBIND")},
		},
	}
	eipDescResp, _ = cvmc.DescribeAddresses(eipDescReq)
	b, _ = json.Marshal(eipDescResp)
	t.Logf("eip desc resp=%s", b)

	// here we must wait otherwise bind will fail because nat gateway not found
	time.Sleep(10 * time.Second)
	bindReq := NewEipBindNatGatewayRequest()
	bindReq.NatId = descResp.Data[0].NatId
	bindReq.VpcId = descResp.Data[0].UnVpcId
	// bind doesn't support duplicate, so here we must choose another eip
	bindReq.AssignedEipSet = []*string{eipDescResp.Response.AddressSet[1].AddressIp}
	bindResp, _ := c.EipBindNatGateway(bindReq)
	b, _ = json.Marshal(bindResp)
	t.Logf("eip nat bind resp=%s", b)
	taskReq := NewDescribeVpcTaskResultRequest()
	taskReq.TaskId = bindResp.TaskId
	for {
		taskResp, err := c.DescribeVpcTaskResult(taskReq)
		b, _ = json.Marshal(taskResp)
		t.Logf("task desc resp=%s", b)
		if _, ok := err.(*common.APIError); ok {
			t.Errorf("Fail err=%v, task desc resp=%v", err, taskResp)
			break
		}
		if *taskResp.Data.Status == 0 {
			break
		} else if *taskResp.Data.Status == 1 {
			// fail, need retry delete
			break
		}
		time.Sleep(10 * time.Second)
	}

	descResp, _ = c.DescribeNatGateway(descReq)
	b, _ = json.Marshal(descResp)
	t.Logf("nat desc resp=%s", b)

	// unbind
	unbindReq := NewEipUnBindNatGatewayRequest()
	unbindReq.NatId = descResp.Data[0].NatId
	unbindReq.VpcId = descResp.Data[0].UnVpcId
	unbindReq.AssignedEipSet = []*string{eipDescResp.Response.AddressSet[1].AddressIp}
	unbindResp, _ := c.EipUnBindNatGateway(unbindReq)
	b, _ = json.Marshal(unbindResp)
	t.Logf("eip nat unbind resp=%s", b)
	taskReq.TaskId = unbindResp.TaskId
	for {
		taskResp, err := c.DescribeVpcTaskResult(taskReq)
		b, _ = json.Marshal(taskResp)
		t.Logf("task desc resp=%s", b)
		if _, ok := err.(*common.APIError); ok {
			t.Errorf("Fail err=%v, task desc resp=%v", err, taskResp)
			break
		}
		if *taskResp.Data.Status == 0 {
			break
		} else if *taskResp.Data.Status == 1 {
			// fail, need retry delete
			break
		}
		time.Sleep(10 * time.Second)
	}

	descResp, _ = c.DescribeNatGateway(descReq)
	b, _ = json.Marshal(descResp)
	t.Logf("nat desc resp=%s", b)
	deleteNatGateway(descResp.Data[0].UnVpcId, descResp.Data[0].NatId, t)
}
