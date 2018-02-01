package vpc

import (
	"github.com/zqfan/tencentcloud-sdk-go/common"
)

const APIVersion = ""

func NewCreateNatGatewayRequest() (request *CreateNatGatewayRequest) {
	request = &CreateNatGatewayRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "CreateNatGateway")
	return
}

func NewCreateNatGatewayResponse() (response *CreateNatGatewayResponse) {
	response = &CreateNatGatewayResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) CreateNatGateway(request *CreateNatGatewayRequest) (response *CreateNatGatewayResponse, err error) {
	if request == nil {
		request = NewCreateNatGatewayRequest()
	}
	response = NewCreateNatGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNatGatewayRequest() (request *DescribeNatGatewayRequest) {
	request = &DescribeNatGatewayRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeNatGateway")
	return
}

func NewDescribeNatGatewayResponse() (response *DescribeNatGatewayResponse) {
	response = &DescribeNatGatewayResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeNatGateway(request *DescribeNatGatewayRequest) (response *DescribeNatGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeNatGatewayRequest()
	}
	response = NewDescribeNatGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNatGatewayRequest() (request *ModifyNatGatewayRequest) {
	request = &ModifyNatGatewayRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "ModifyNatGateway")
	return
}

func NewModifyNatGatewayResponse() (response *ModifyNatGatewayResponse) {
	response = &ModifyNatGatewayResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) ModifyNatGateway(request *ModifyNatGatewayRequest) (response *ModifyNatGatewayResponse, err error) {
	if request == nil {
		request = NewModifyNatGatewayRequest()
	}
	response = NewModifyNatGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNatGatewayRequest() (request *DeleteNatGatewayRequest) {
	request = &DeleteNatGatewayRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DeleteNatGateway")
	return
}

func NewDeleteNatGatewayResponse() (response *DeleteNatGatewayResponse) {
	response = &DeleteNatGatewayResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DeleteNatGateway(request *DeleteNatGatewayRequest) (response *DeleteNatGatewayResponse, err error) {
	if request == nil {
		request = NewDeleteNatGatewayRequest()
	}
	response = NewDeleteNatGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcExRequest() (request *DescribeVpcExRequest) {
	request = &DescribeVpcExRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcEx")
	return
}

func NewDescribeVpcExResponse() (response *DescribeVpcExResponse) {
	response = &DescribeVpcExResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeVpcEx(request *DescribeVpcExRequest) (response *DescribeVpcExResponse, err error) {
	if request == nil {
		request = NewDescribeVpcExRequest()
	}
	response = NewDescribeVpcExResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcTaskResultRequest() (request *DescribeVpcTaskResultRequest) {
	request = &DescribeVpcTaskResultRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcTaskResult")
	return
}

func NewDescribeVpcTaskResultResponse() (response *DescribeVpcTaskResultResponse) {
	response = &DescribeVpcTaskResultResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeVpcTaskResult(request *DescribeVpcTaskResultRequest) (response *DescribeVpcTaskResultResponse, err error) {
	if request == nil {
		request = NewDescribeVpcTaskResultRequest()
	}
	response = NewDescribeVpcTaskResultResponse()
	err = c.Send(request, response)
	return
}

func NewEipBindNatGatewayRequest() (request *EipBindNatGatewayRequest) {
	request = &EipBindNatGatewayRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "EipBindNatGateway")
	return
}

func NewEipBindNatGatewayResponse() (response *EipBindNatGatewayResponse) {
	response = &EipBindNatGatewayResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) EipBindNatGateway(request *EipBindNatGatewayRequest) (response *EipBindNatGatewayResponse, err error) {
	if request == nil {
		request = NewEipBindNatGatewayRequest()
	}
	response = NewEipBindNatGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewEipUnBindNatGatewayRequest() (request *EipUnBindNatGatewayRequest) {
	request = &EipUnBindNatGatewayRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().WithApiInfo("vpc", APIVersion, "EipUnBindNatGateway")
	return
}

func NewEipUnBindNatGatewayResponse() (response *EipUnBindNatGatewayResponse) {
	response = &EipUnBindNatGatewayResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) EipUnBindNatGateway(request *EipUnBindNatGatewayRequest) (response *EipUnBindNatGatewayResponse, err error) {
	if request == nil {
		request = NewEipUnBindNatGatewayRequest()
	}
	response = NewEipUnBindNatGatewayResponse()
	err = c.Send(request, response)
	return
}
