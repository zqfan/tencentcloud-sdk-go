# tencentcloud-sdk-go

## Dedicated API example

```
package main

import (
    "github.com/zqfan/tencentcloud-sdk-go/service/cvm"
)

func main() {
    client, _ := cvm.NewClientWithSecretID("YOUR_SECRET_ID", "YOUR_SECRET_KEY", "REGION_NAME")
    request := NewDescribeAddressesRequest()
    request.Limit = "10"
    // get response structure
    response, err := c.DescribeAddresses(request)
    eips := response.Response.AddressSet
}
```

## Common API Example

```
package main

import (
    "github.com/zqfan/tencentcloud-sdk-go/client"
)

func main() {
    client := client.NewClient("YOUR_SECRET_ID", "YOUR_SECRET_KEY", "REGION_NAME")
    params := map[string]string {
        "Action": "DescribeInstances",
        "SignatureMethod": "HmacSHA256",
    }
    // get raw text response body from server
    response, err := client.SendRequest("cvm", params)
}
```

If you want to print debug info, set ``Client`` object's attribute ``Debug`` to true like: ``client.Debug = true``
