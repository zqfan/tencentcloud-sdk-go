package main

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

// goFormat returns the Go formated string of the input.
func goFormat(s string) string {
	buf, err := format.Source([]byte(s))
	if err != nil {
		log.Panicf("Failed to format %s because %s", s, err.Error())
	}
	return string(buf)
}

func main() {
	tmpl := template.Must(template.New("api").Parse(`package {{.PkgName}}

import (
    "github.com/zqfan/tencentcloud-sdk-go/common"
)

const APIVersion = "{{.APIVersion}}"
{{range .Actions}}
func New{{.Action}}Request() (request *{{.Action}}Request) {
    request = &{{.Action}}Request{
        BaseRequest: &common.BaseRequest{},
    }
    request.Init().WithApiInfo("{{.SubDomain}}", APIVersion, "{{.Action}}")
    return
}

func New{{.Action}}Response() (response *{{.Action}}Response) {
    response = &{{.Action}}Response{
        BaseResponse: &common.BaseResponse{},
    }
    return
}

func (c *Client) {{.Action}}(request *{{.Action}}Request) (response *{{.Action}}Response, err error) {
    if request == nil {
        request = New{{.Action}}Request()
    }
    response = New{{.Action}}Response()
    err = c.Send(request, response)
    return
}
{{end}}`))
	type Action struct {
		Action    string
		SubDomain string
	}
	type Service struct {
		PkgName    string
		APIVersion string
		Actions    []Action
	}
	services := []Service{
		Service{
			PkgName:    "cvm",
			APIVersion: "2017-03-12",
			Actions: []Action{
				{"DescribeAddresses", "eip"},
				{"ReleaseAddresses", "eip"},
				{"ModifyAddressAttribute", "eip"},
				{"AllocateAddresses", "eip"},
				{"AssociateAddress", "eip"},
				{"DisassociateAddress", "eip"},
				{"RunInstances", "cvm"},
				{"DescribeInstances", "cvm"},
				{"TerminateInstances", "cvm"},
			},
		},
	}
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	serviceDir := filepath.Join(dir, "..", "services")
	for _, service := range services {
		var buffer bytes.Buffer
		tmpl.Execute(&buffer, service)
		apiFilePath := filepath.Join(serviceDir, service.PkgName, "api.go")
		err := ioutil.WriteFile(apiFilePath, []byte(goFormat(buffer.String())), 0644)
		if err != nil {
			log.Println(err)
		}
	}
}
