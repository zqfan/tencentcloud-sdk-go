package lb

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

func TestLBCRUD(t *testing.T) {
	c, _ := newClient()
	descReq := NewDescribeLoadBalancersRequest()
	descResp, err := c.DescribeLoadBalancers(descReq)
	if _, ok := err.(*common.APIError); ok {
		t.Errorf("[ERROR] err=%v", err)
		return
	}
	b, _ := json.Marshal(descResp)
	t.Logf("lb desc resp=%s", b)
}

func TestListenerCRUD(t *testing.T) {
	c, _ := newClient()
	lbDescReq := NewDescribeLoadBalancersRequest()
	lbDescReq.Forward = common.IntPtr(LBForwardTypeApplication)
	lbDescResp, err := c.DescribeLoadBalancers(lbDescReq)
	lbDescJson, _ := json.Marshal(lbDescResp)
	t.Logf("lb desc resp=%s", lbDescJson)
	if len(lbDescResp.LoadBalancerSet) == 0 {
		t.Errorf("[ERROR] Application LB not found")
		return
	}

	descReq := NewDescribeForwardLBListenersRequest()
	descReq.LoadBalancerId = lbDescResp.LoadBalancerSet[0].UnLoadBalancerId
	descResp, err := c.DescribeForwardLBListeners(descReq)
	if _, ok := err.(*common.APIError); ok {
		t.Errorf("[ERROR] err=%v", err)
		return
	}
	descJson, _ := json.Marshal(descResp)
	t.Logf("listener desc resp=%s", descJson)
}

func filterServerWithPublicIp() *string {
	c, _ := cvm.NewClientWithSecretId(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
		"ap-guangzhou",
	)
	descReq := cvm.NewDescribeInstancesRequest()
	descResp, _ := c.DescribeInstances(descReq)
	for _, ins := range descResp.Response.InstanceSet {
		if len(ins.PublicIpAddresses) > 0 {
			return ins.InstanceId
		}
	}
	return nil
}

func TestServerAttachDetach(t *testing.T) {
	c, _ := newClient()

	lbDescReq := NewDescribeLoadBalancersRequest()
	lbDescReq.Forward = common.IntPtr(LBForwardTypeApplication)
	lbDescResp, err := c.DescribeLoadBalancers(lbDescReq)
	if len(lbDescResp.LoadBalancerSet) == 0 {
		t.Errorf("[ERROR] Application LB not found")
		return
	}

	lblDescReq := NewDescribeForwardLBListenersRequest()
	lblDescReq.LoadBalancerId = lbDescResp.LoadBalancerSet[0].UnLoadBalancerId
	lblDescReq.Protocol = common.IntPtr(LBListenerProtocolHTTPS)
	lblDescResp, err := c.DescribeForwardLBListeners(lblDescReq)
	if len(lblDescResp.ListenerSet) == 0 {
		t.Errorf("[ERROR] TCP Listener not found")
		return
	}

	insId := filterServerWithPublicIp()
	if insId == nil {
		t.Errorf("[ERROR] Instance with public ip not found")
		return
	}

	attachReq := NewRegisterInstancesWithForwardLBSeventhListenerRequest()
	attachReq.LoadBalancerId = lbDescResp.LoadBalancerSet[0].UnLoadBalancerId
	attachReq.ListenerId = lblDescResp.ListenerSet[0].ListenerId
	attachReq.Backends = []*Backend{
		&Backend{
			InstanceId: insId,
			Port:       common.IntPtr(2333),
		},
	}
	attachResp, err := c.RegisterInstancesWithForwardLBSeventhListener(attachReq)
	if _, ok := err.(*common.APIError); ok {
		t.Errorf("[ERROR] err=%v", err)
		return
	}
	attachJson, _ := json.Marshal(attachResp)
	t.Logf("attach resp=%s", attachJson)

	taskReq := NewDescribeLoadBalancersTaskResultRequest()
	taskReq.RequestId = attachResp.RequestId
	for {
		taskResp, err := c.DescribeLoadBalancersTaskResult(taskReq)
		if _, ok := err.(*common.APIError); ok {
			t.Errorf("[ERROR] err=%v", err)
			return
		}
		taskJson, _ := json.Marshal(taskResp)
		t.Logf("attach resp=%s", taskJson)
		if *taskResp.Data.Status == LBTaskSuccess {
			break
		} else if *taskResp.Data.Status == LBTaskFail {
			t.Errorf("[ERROR] attach server to LB failed")
			return
		}
		time.Sleep(10 * time.Second)
	}

	backendReq := NewDescribeForwardLBBackendsRequest()
	backendReq.LoadBalancerId = lbDescResp.LoadBalancerSet[0].UnLoadBalancerId
	backendReq.ListenerIds = []*string{lblDescResp.ListenerSet[0].ListenerId}
	backendResp, err := c.DescribeForwardLBBackends(backendReq)
	backendJson, _ := json.Marshal(backendResp)
	t.Logf("backend desc resp=%s", backendJson)

	detachReq := NewDeregisterInstancesFromForwardLBRequest()
	detachReq.LoadBalancerId = lbDescResp.LoadBalancerSet[0].UnLoadBalancerId
	detachReq.ListenerId = lblDescResp.ListenerSet[0].ListenerId
	detachReq.Backends = []*Backend{
		&Backend{
			InstanceId: insId,
			Port:       common.IntPtr(2333),
		},
	}
	detachResp, err := c.DeregisterInstancesFromForwardLB(detachReq)
	if _, ok := err.(*common.APIError); ok {
		t.Errorf("[ERROR] err=%v", err)
		return
	}
	detachJson, _ := json.Marshal(detachResp)
	t.Logf("detach resp=%s", detachJson)

	taskReq.RequestId = detachResp.RequestId
	for {
		taskResp, err := c.DescribeLoadBalancersTaskResult(taskReq)
		if _, ok := err.(*common.APIError); ok {
			t.Errorf("[ERROR] err=%v", err)
			return
		}
		taskJson, _ := json.Marshal(taskResp)
		t.Logf("detach resp=%s", taskJson)
		if *taskResp.Data.Status == LBTaskSuccess {
			break
		} else if *taskResp.Data.Status == LBTaskFail {
			t.Errorf("[ERROR] detach server from LB failed")
			return
		}
		time.Sleep(10 * time.Second)
	}
}
