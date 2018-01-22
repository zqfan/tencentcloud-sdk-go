package cvm

import (
	"os"
	"testing"
)

func newClient() (*Client, error) {
	return NewClientWithSecretId(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
		"ap-guangzhou",
	)
}

func TestDescribeAddresses(t *testing.T) {
	c, _ := newClient()
	request := NewDescribeAddressesRequest()
	request.Limit = "100"
	//request.SetHttpMethod("POST")
	response, err := c.DescribeAddresses(request)
	if err != nil {
		t.Errorf("Fail: err=%v, response=%s", err, response)
	}
}
