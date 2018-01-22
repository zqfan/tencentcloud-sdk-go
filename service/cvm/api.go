package cvm

import (
	"github.com/zqfan/tencentcloud-sdk-go/common"
)

func NewDescribeAddressesRequest() (request *DescribeAddressesRequest) {
	request = &DescribeAddressesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().WithApiInfo("eip", "2017-03-12", "DescribeAddresses")
	return
}

func NewDescribeAddressesResponse() (response *DescribeAddressesResponse) {
	response = &DescribeAddressesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAddresses(request *DescribeAddressesRequest) (response *DescribeAddressesResponse, err error) {
	if request == nil {
		request = NewDescribeAddressesRequest()
	}
	response = NewDescribeAddressesResponse()
	err = c.Send(request, response)
	return
}

func (c *Client) ReleaseAddresses(request *ReleaseAddressesRequest) (*ReleaseAddressesResponse, error) {
	return nil, nil
}
