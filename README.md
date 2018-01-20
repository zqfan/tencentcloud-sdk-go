# qcloud-sdk-go

## Example

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
