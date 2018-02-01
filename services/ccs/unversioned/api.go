package ccs

import (
	"github.com/zqfan/tencentcloud-sdk-go/common"
)

const APIVersion = ""

func NewCreateClusterRequest() (request *CreateClusterRequest) {
	request = &CreateClusterRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().WithApiInfo("ccs", APIVersion, "CreateCluster")
	return
}

func NewCreateClusterResponse() (response *CreateClusterResponse) {
	response = &CreateClusterResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
	if request == nil {
		request = NewCreateClusterRequest()
	}
	response = NewCreateClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterRequest() (request *DescribeClusterRequest) {
	request = &DescribeClusterRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().WithApiInfo("ccs", APIVersion, "DescribeCluster")
	return
}

func NewDescribeClusterResponse() (response *DescribeClusterResponse) {
	response = &DescribeClusterResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCluster(request *DescribeClusterRequest) (response *DescribeClusterResponse, err error) {
	if request == nil {
		request = NewDescribeClusterRequest()
	}
	response = NewDescribeClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteClusterRequest() (request *DeleteClusterRequest) {
	request = &DeleteClusterRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().WithApiInfo("ccs", APIVersion, "DeleteCluster")
	return
}

func NewDeleteClusterResponse() (response *DeleteClusterResponse) {
	response = &DeleteClusterResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DeleteCluster(request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
	if request == nil {
		request = NewDeleteClusterRequest()
	}
	response = NewDeleteClusterResponse()
	err = c.Send(request, response)
	return
}
