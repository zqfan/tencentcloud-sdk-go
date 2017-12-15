# qcloud-sdk-go

## Example

```
package main

import (
    "github.com/zqfan/qcloudapi-sdk-go/client"
)

func main() {
    client := client.NewClient("YOUR_SECRET_ID", "YOUR_SECRET_KEY", "REGION_NAME")
    params := map[string]string {
        "Action": "DescribeInstances",
    }
    // get raw text response body from server
    response, err := client.SendRequest("cvm", params)
}
```

If you want to print debug info, set ``Client`` object's attribute ``Debug`` to true like: ``client.Debug = true``
