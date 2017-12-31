package main

import (
	"fmt"
	"github.com/zqfan/qcloudapi-sdk-go/client"
	"os"
)

func main() {
	client := client.NewClient(os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
		"ap-guangzhou")
	client.Debug = true

	params := map[string]string{
		"Action": "DescribeInstances",
	}

	_, err := client.SendRequest("cvm", params)
	if err != nil {
		fmt.Print("Error.", err)
	}

	return
}
