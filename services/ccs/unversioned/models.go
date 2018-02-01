package ccs

import (
	"github.com/zqfan/tencentcloud-sdk-go/common"
)

type DescribeClusterRequest struct {
	*common.BaseRequest
	ClusterIds  []*string `name:"clusterIds" list`
	ClusterName *string   `name:"clusterName"`
	Status      *string   `name:"status"`
	OrderField  *string   `name:"orderField"`
	OrderType   *string   `name:"orderType"`
	Offset      *int      `name:"offset"`
	Limit       *int      `name:"limit"`
}

type Cluster struct {
	ClusterCIDR             *string `json:"clusterCIDR"`
	ClusterExternalEndpoint *string `json:"clusterExternalEndpoint"`
	ClusterId               *string `json:"clusterId"`
	ClusterName             *string `json:"clusterName"`
	CreatedAt               *string `json:"createdAt"`
	Description             *string `json:"description"`
	K8sVersion              *string `json:"k8sVersion"`
	MasterLbSubnetId        *string `json:"masterLbSubnetId"`
	NodeNum                 *int    `json:"nodeNum"`
	NodeStatus              *string `json:"nodeStatus"`
	OpenHttps               *int    `json:"openHttps"`
	OS                      *string `json:"os"`
	ProjectId               *int    `json:"projectId"`
	Region                  *string `json:"region"`
	RegionId                *int    `json:"regionId"`
	Status                  *string `json:"status"`
	TotalCPU                *int    `json:"totalCpu"`
	TotalMem                *int    `json:"totalMem"`
	UnVpcId                 *string `json:"unVpcId"`
	UpdatedAt               *string `json:"updatedAt"`
	VpcId                   *int    `json:"vpcId"`
}

type DescribeClusterResponse struct {
	*common.BaseResponse
	Code     *int    `json:"code"`
	CodeDesc *string `json:"codeDesc"`
	Message  *string `json:"message"`
	Data     *struct {
		TotalCount *int       `json:"totalCount"`
		Clusters   []*Cluster `json:"clusters"`
	} `json:"data"`
}

type CreateClusterRequest struct {
	*common.BaseRequest
	ClusterName               *string `name:"clusterName"`
	ClusterDesc               *string `name:"clusterDesc"`
	ClusterCIDR               *string `name:"clusterCIDR"`
	IgnoreClusterCIDRConflict *int    `name:"ignoreClusterCIDRConflict"`
	ZoneId                    *string `name:"zoneId"`
	GoodsNum                  *int    `name:"goodsNum"`
	CPU                       *int    `name:"cpu"`
	Mem                       *int    `name:"mem"`
	OSName                    *string `name:"osName"`
	InstanceType              *string `name:"instanceType"`
	CVMType                   *string `name:"cvmType"`
	BandwidthType             *string `name:"bandwidthType"`
	Bandwidth                 *int    `name:"bandwidth"`
	WanIp                     *int    `name:"wanIp"`
	VpcId                     *string `name:"vpcId"`
	SubnetId                  *string `name:"subnetId"`
	IsVpcGateway              *int    `name:"isVpcGateway"`
	RootSize                  *int    `name:"rootSize"`
	StorageSize               *int    `name:"storageSize"`
	Password                  *string `name:"password"`
	KeyId                     *string `name:"keyId"`
	Period                    *int    `name:"period"`
	SgId                      *string `name:"sgId"`
}

type CreateClusterResponse struct {
	*common.BaseResponse
	Code     *int    `json:"code"`
	CodeDesc *string `json:"codeDesc"`
	Message  *string `json:"message"`
	Data     *struct {
		RequestId *int    `json:"requestId"`
		ClusterId *string `json:"clusterId"`
	} `json:"data"`
}

type DeleteClusterRequest struct {
	*common.BaseRequest
	ClusterId      *string `name:"clusterId"`
	NodeDeleteMode *string `name:"nodeDeleteMode"`
}

type DeleteClusterResponse struct {
	*common.BaseResponse
	Code     *int    `json:"code"`
	CodeDesc *string `json:"codeDesc"`
	Message  *string `json:"message"`
	Data     *struct {
		RequestId *int `json:"requestId"`
	} `json:"data"`
}

type Request struct {
}

type Response struct {
}
