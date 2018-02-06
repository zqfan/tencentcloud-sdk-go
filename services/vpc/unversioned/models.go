package vpc

import (
	"github.com/zqfan/tencentcloud-sdk-go/common"
)

type CreateNatGatewayRequest struct {
	*common.BaseRequest
	NatName         *string   `name:"natName"`
	VpcId           *string   `name:"vpcId"`
	MaxConcurrent   *int      `name:"maxConcurrent"`
	Bandwidth       *int      `name:"bandwidth"`
	AssignedEipSet  []*string `name:"assignedEipSet" list`
	AutoAllocEipNum *int      `name:"autoAllocEipNum"`
}

type CreateNatGatewayResponse struct {
	*common.BaseResponse
	Code         *int    `json:"code"`
	CodeDesc     *string `json:"codeDesc"`
	Message      *string `json:"message"`
	BillId       *string `json:"billId"`
	NatGatewayId *string `json:"natGatewayId"`
}

type DescribeNatGatewayRequest struct {
	*common.BaseRequest
	NatId          *string `name:"natId"`
	NatName        *string `name:"natName"`
	VpcId          *string `name:"vpcId"`
	Offset         *int    `name:"offset"`
	Limit          *int    `name:"limit"`
	OrderField     *string `name:"orderField"`
	OrderDirection *string `name:"orderDirection"`
}

type NatGateway struct {
	AppId            *string   `json:"appId"`
	NatId            *string   `json:"natId"`
	VpcId            *int      `json:"vpcId"`
	UnVpcId          *string   `json:"unVpcID"`
	VpcName          *string   `json:"vpcName"`
	NatName          *string   `json:"natName"`
	State            *int      `json:"state"`
	MaxConcurrent    *int      `json:"maxConcurrent"`
	Bandwidth        *int      `json:"bandwidth"`
	EipCount         *int      `json:"eipCount"`
	EipSet           []*string `json:"eipSet"`
	BlockedEipSet    []*string `json:"blockedEipSet"`
	CreateTime       *string   `json:"createTime"`
	ProductionStatus *int      `json:"productionStatus"`
}

type DescribeNatGatewayResponse struct {
	*common.BaseResponse
	Code       *int          `json:"code"`
	CodeDesc   *string       `json:"codeDesc"`
	Message    *string       `json:"message"`
	TotalCount *int          `json:"totalCount"`
	Data       []*NatGateway `json:"data"`
}

type ModifyNatGatewayRequest struct {
	*common.BaseRequest
	VpcId     *string `name:"vpcId"`
	NatId     *string `name:"natId"`
	NatName   *string `name:"natName"`
	Bandwidth *int    `name:"bandwidth"`
}

type ModifyNatGatewayResponse struct {
	*common.BaseResponse
	Code    *string `json:"code"`
	Message *string `json:"message"`
}

type DeleteNatGatewayRequest struct {
	*common.BaseRequest
	VpcId *string `name:"vpcId"`
	NatId *string `name:"natId"`
}

type DeleteNatGatewayResponse struct {
	*common.BaseResponse
	Code     *int    `json:"code"`
	CodeDesc *string `json:"codeDesc"`
	Message  *string `json:"message"`
	TaskId   *int    `json:"taskId"`
}

type DescribeVpcExRequest struct {
	*common.BaseRequest
	VpcId          *string `name:"vpcId"`
	VpcName        *string `name:"vpcName"`
	Offset         *int    `name:"offset"`
	Limit          *int    `name:"limit"`
	OrderField     *string `name:"orderField"`
	OrderDirection *string `name:"orderDirection"`
}

type Vpc struct {
	VpcId          *string `json:"vpcId"`
	UnVpcId        *string `json:"unVpcId"`
	VpcName        *string `json:"vpcName"`
	CidrBlock      *string `json:"cidrBlock"`
	SubnetNum      *int    `json:"subnetNum"`
	RouteTableNum  *int    `json:"routeTableNum"`
	VpnGwNum       *int    `json:"vpnGwNum"`
	VpcPeerNum     *int    `json:"vpcPeerNum"`
	SflowNum       *int    `json:"sflowNum"`
	IsDefault      *bool   `json:"isDefault"`
	IsMulticast    *bool   `json:"isMulticast"`
	VpcDeviceNum   *int    `json:"vpcDeviceNum"`
	ClassicLinkNum *int    `json:"classicLinkNum"`
	VpgNum         *int    `json:"vpgNum"`
	NatNum         *int    `json:"natNum"`
	CreateTime     *string `json:"createTime"`
}

type DescribeVpcExResponse struct {
	*common.BaseResponse
	Code       *int    `json:"code"`
	Message    *string `json:"message"`
	TotalCount *int    `json:"totalCount"`
	Data       []*Vpc  `json:"data"`
}

type DescribeVpcTaskResultRequest struct {
	*common.BaseRequest
	TaskId *int `name:"taskId"`
}

type DescribeVpcTaskResultResponse struct {
	*common.BaseResponse
	Code     *int    `json:"code"`
	CodeDesc *string `json:"codeDesc"`
	Message  *string `json:"message"`
	Data     *struct {
		Status *int `json:"status"`
		Output *struct {
			ErrorCode *int    `json:"errorCode"`
			ErrorMsg  *string `json:"errorMsg"`
		} `json:"output"`
	} `json:"data"`
}

type EipBindNatGatewayRequest struct {
	*common.BaseRequest
	NatId           *string   `name:"natId"`
	VpcId           *string   `name:"vpcId"`
	AssignedEipSet  []*string `name:"assignedEipSet"`
	AutoAllocEipNum *int      `name:"autoAllocEipNum"`
}

type EipBindNatGatewayResponse struct {
	*common.BaseResponse
	Code     *int    `json:"code"`
	Message  *string `json:"message"`
	CodeDesc *string `json:"codeDesc"`
	TaskId   *int    `json:"taskId"`
}

type EipUnBindNatGatewayRequest struct {
	*common.BaseRequest
	NatId          *string   `name:"natId"`
	VpcId          *string   `name:"vpcId"`
	AssignedEipSet []*string `name:"assignedEipSet"`
}

type EipUnBindNatGatewayResponse struct {
	*common.BaseResponse
	Code     *int    `json:"code"`
	Message  *string `json:"message"`
	CodeDesc *string `json:"codeDesc"`
	TaskId   *int    `json:"taskId"`
}

type QueryNatGatewayProductionStatusRequest struct {
	*common.BaseRequest
	BillId *string `name:"billId"`
}

const (
	BillStatusSuccess = 0
	BillStatusFail    = 1
	BillStatusDoing   = 2
)

type QueryNatGatewayProductionStatusResponse struct {
	*common.BaseResponse
	Code     *int    `json:"code"`
	Message  *string `json:"message"`
	CodeDesc *string `json:"codeDesc"`
	Data     *struct {
		Status    *int    `json:"status"`
		ErrorCode *string `json:"errorcode"`
	} `json:"data"`
}

type GetDnaptRuleRequest struct {
	*common.BaseRequest
	VpcId *string `name:"vpcId"`
	NatId *string `name:"natId"`
}

type DnaptRule struct {
	CreateTime  *string `json:"createTime"`
	Description *string `json:"description"`
	Eip         *string `json:"eip" name:"eip"`
	Eport       *int    `json:"eport" name:"eport"`
	NatId       *int    `json:"natId"`
	Owner       *string `json:"owner"`
	Pip         *string `json:"pip"`
	PipType     *int    `json:"pipType"`
	Pport       *int    `json:"pport"`
	Proto       *string `json:"proto" name:"proto"`
	UniqNatId   *string `json:"uniqNatId"`
	UniqVpcId   *string `json:"uniqVpcId"`
	VpcId       *int    `json:"vpcId"`
}

type GetDnaptRuleResponse struct {
	*common.BaseResponse
	Code     *int    `json:"code"`
	CodeDesc *string `json:"codeDesc"`
	Data     *struct {
		TotalNum *int         `json:"totalNum"`
		Detail   []*DnaptRule `json:"detail"`
	} `json:"data"`
	Message *string `json:"message"`
}

type AddDnaptRuleRequest struct {
	*common.BaseRequest
	VpcId       *string `name:"vpcId"`
	NatId       *string `name:"natId"`
	Proto       *string `name:"proto"`
	Eip         *string `name:"eip"`
	Eport       *int    `name:"eport"`
	Pip         *string `name:"pip"`
	Pport       *int    `name:"pport"`
	Description *string `name:"description"`
}

type AddDnaptRuleResponse struct {
	*common.BaseResponse
	Code     *int    `json:"code"`
	Message  *string `json:"message"`
	CodeDesc *string `json:"codeDesc"`
	Data     []*struct {
	} `json:"data"`
}

type DeleteDnaptRuleRequest struct {
	*common.BaseRequest
	VpcId    *string      `name:"vpcId"`
	NatId    *string      `name:"natId"`
	DnatList []*DnaptRule `name:"dnatList" list`
}

type DeleteDnaptRuleResponse struct {
	*common.BaseResponse
	Code     *int    `json:"code"`
	Message  *string `json:"message"`
	CodeDesc *string `json:"codeDesc"`
	Data     []*struct {
	} `json:"data"`
}

type ModifyDnaptRuleRequest struct {
	*common.BaseRequest
	VpcId       *string `name:"vpcId"`
	NatId       *string `name:"natId"`
	OldProto    *string `name:"oldProto"`
	OldEip      *string `name:"oldEip"`
	OldEport    *int    `name:"oldEport"`
	Proto       *string `name:"proto"`
	Eip         *string `name:"eip"`
	Eport       *int    `name:"eport"`
	Pip         *string `name:"pip"`
	Pport       *int    `name:"pport"`
	Description *string `name:"description"`
}

type ModifyDnaptRuleResponse struct {
	*common.BaseResponse
	Code     *int    `json:"code"`
	Message  *string `json:"message"`
	CodeDesc *string `json:"codeDesc"`
	Data     []*struct {
	} `json:"data"`
}

type Request struct {
	*common.BaseRequest
}

type Response struct {
	*common.BaseResponse
}
