package cvm

import (
	"github.com/zqfan/tencentcloud-sdk-go/common"
)

type Error struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
}

type Filter struct {
	Name   string
	Values []string
}

type DescribeAddressesRequest struct {
	*common.BaseRequest
	AddressIds []string `name:"AddressIds" list`
	Filters    []Filter `name:"Filters" list`
	Offset     string   `name:"Offset" type:"int"`
	Limit      string   `name:"Limit" type:"int"`
}

type Address struct {
	AddressId             string `json:"AddressId"`
	AddressIp             string `json:"AddressIp"`
	AddressName           string `json:"AddressName"`
	AddressState          string `json:"AddressState"`
	AddressStatus         string `json:"AddressStatus"`
	BindedResourceId      string `json:"BindedResourceId"`
	CreatedTime           string `json:"CreatedTime"`
	InstanceId            string `json:"InstanceId"`
	IsArrears             bool   `json:"IsArrears"`
	IsBlocked             bool   `json:"IsBlocked"`
	IsEipDirectConnection bool   `json:"IsEipDirectConnection"`
	NetworkInterfaceId    string `json:"NetworkInterfaceId"`
	PrivateAddressIp      string `json:"PrivateAddressIp"`
}

type DescribeAddressesResponse struct {
	*common.BaseResponse
	Response struct {
		RequestId  string    `json:"RequestId"`
		TotalCount int       `json:"TotalCount"`
		AddressSet []Address `json:"AddressSet`
	} `json:"Response"`
}

type ReleaseAddressesRequest struct {
}

type ReleaseAddressesResponse struct {
}
