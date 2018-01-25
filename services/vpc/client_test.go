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
	if _, ok := err.(*common.APIError); ok {
		t.Errorf("Fail err=%v", err)
		return
	}
	b, _ = json.Marshal(createResp)
	t.Logf("resp=%s", b)
	// retrieve
	descReq := NewDescribeNatGatewayRequest()
	descReq.NatName = common.StringPtr("nat-test-xyz")
	descResp, err := c.DescribeNatGateway(descReq)
	if _, ok := err.(*common.APIError); ok {
		t.Errorf("Fail err=%v, resp=%v", err, descResp)
		return
	}
	b, _ = json.Marshal(descResp)
	t.Logf("resp=%s", b)
	deleteReq := NewDeleteNatGatewayRequest()
	deleteReq.VpcId = descResp.Data[0].UnVpcId
	deleteReq.NatId = descResp.Data[0].NatId
	var deleteResp *DeleteNatGatewayResponse
	for {
		deleteResp, err = c.DeleteNatGateway(deleteReq)
		b, _ = json.Marshal(deleteResp)
		t.Logf("resp=%s", b)
		if apiErr, ok := err.(*common.APIError); ok {
			if apiErr.Code == "InvalidAddressState" {
				time.Sleep(10 * time.Second)
				continue
			}
		}
		if err != nil {
			t.Errorf("%s", err)
		}
		return
	}
}
