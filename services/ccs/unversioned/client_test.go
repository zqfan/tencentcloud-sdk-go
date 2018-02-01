package ccs

import (
	"encoding/json"
	"github.com/zqfan/tencentcloud-sdk-go/common"
	"os"
	"testing"
	"time"
)

var eip *string

func newClient() (*Client, error) {
	return NewClientWithSecretId(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
		"ap-guangzhou",
	)
}

func TestClusterCRUD(t *testing.T) {
	c, _ := newClient()
	descReq := NewDescribeClusterRequest()
	descResp, err := c.DescribeCluster(descReq)
	if _, ok := err.(*common.APIError); ok {
		t.Errorf("Fail: err=%v", err)
		return
	}
	b, _ := json.Marshal(descResp)
	t.Logf("cluster desc resp=%s", b)
	// create
	createReq := NewCreateClusterRequest()
	createReq.ZoneId = common.StringPtr("100003")
	createReq.ClusterName = common.StringPtr("cluster-test")
	createReq.CPU = common.IntPtr(2)
	createReq.Mem = common.IntPtr(4)
	createReq.OSName = common.StringPtr("ubuntu16.04.1 LTSx86_64")
	createReq.Bandwidth = common.IntPtr(1)
	createReq.BandwidthType = common.StringPtr("PayByTraffic")
	createReq.SubnetId = common.StringPtr("subnet-q2hzxwey")
	createReq.VpcId = common.StringPtr("vpc-kg60ct5z")
	createReq.IsVpcGateway = common.IntPtr(0)
	createReq.StorageSize = common.IntPtr(0)
	createReq.RootSize = common.IntPtr(20)
	createReq.GoodsNum = common.IntPtr(0)
	createReq.ClusterCIDR = common.StringPtr("172.19.0.0/19")
	createResp, err := c.CreateCluster(createReq)
	if _, ok := err.(*common.APIError); ok {
		t.Errorf("[ERROR] err=%v", err)
		return
	}
	b, _ = json.Marshal(createResp)
	t.Logf("cluster create resp=%s", b)
	// desc
	descResp, err = c.DescribeCluster(descReq)
	if _, ok := err.(*common.APIError); ok {
		t.Errorf("[ERROR] err=%v", err)
		return
	}
	b, _ = json.Marshal(descResp)
	t.Logf("cluster desc resp=%s", b)
	// delete
	deleteReq := NewDeleteClusterRequest()
	deleteReq.ClusterId = createResp.Data.ClusterId
	for {
		deleteResp, err := c.DeleteCluster(deleteReq)
		if apiErr, ok := err.(*common.APIError); ok {
			time.Sleep(10 * time.Second)
			if apiErr.Code == "ClusterNotReadyError" {
				t.Logf("[INFO] err=%v", err)
				continue
			}
			t.Errorf("[ERROR] err=%v", err)
		}
		b, _ = json.Marshal(deleteResp)
		t.Logf("cluster delete resp=%s", b)
		break
	}
}
