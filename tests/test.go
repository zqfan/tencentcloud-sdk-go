package main

import (
    "encoding/json"
    "fmt"
    "os"
    "github.com/zqfan/qcloudapi-sdk-go/client"
)

func main() {
    client := client.NewClient(os.Getenv("TENCENTCLOUD_SECRET_ID"),
                               os.Getenv("TENCENTCLOUD_SECRET_KEY"),
                               "ap-guangzhou")
    client.Debug = true

    params := map[string]string {
        "Action": "DescribeInstances",
    }

    response, err := client.SendRequest("cvm", params)
    if err != nil{
        fmt.Print("Error.", err)
        return
    }

    var jsonresp interface{}
    err = json.Unmarshal([]byte(response), &jsonresp)
    if err != nil {
        fmt.Println(err);
        return
    }
    jsonstr, _ := json.MarshalIndent(jsonresp, "", "  ");
    os.Stdout.Write(jsonstr)

    return
}
