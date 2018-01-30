package vpc

import (
	"encoding/json"
	"github.com/zqfan/tencentcloud-sdk-go/common"
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
	// create
	createReq := NewCreateNatGatewayRequest()
	createReq.VpcId = vpcDescResp.Data[0].UnVpcId
	createReq.NatName = common.StringPtr("nat-test-xyz")
	createReq.MaxConcurrent = common.IntPtr(1000)
	createReq.AutoAllocEipNum = common.IntPtr(1)
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
	deleteReq := NewDeleteNatGatewayRequest()
	deleteReq.VpcId = descResp.Data[0].UnVpcId
	deleteReq.NatId = descResp.Data[0].NatId
	for {
		deleteResp, err := c.DeleteNatGateway(deleteReq)
		b, _ = json.Marshal(deleteResp)
		t.Logf("resp=%s", b)
		if _, ok := err.(*common.APIError); ok {
			t.Errorf("Fail err=%v, resp=%v", err, descResp)
			return
		}
		taskReq := NewDescribeVpcTaskResultRequest()
		taskReq.TaskId = deleteResp.TaskId
		for {
			taskResp, err := c.DescribeVpcTaskResult(taskReq)
			b, _ = json.Marshal(taskResp)
			t.Logf("resp=%s", b)
			if _, ok := err.(*common.APIError); ok {
				t.Errorf("Fail err=%v, resp=%v", err, descResp)
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
