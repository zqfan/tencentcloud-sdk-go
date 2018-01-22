package common

import (
	"io"
	"log"
	"math/rand"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	POST = "POST"
	GET  = "GET"

	RootDomain = "api.qcloud.com"
	V2Path     = "/v2/index.php"
)

type Request interface {
	GetAction() string
	GetBodyReader() io.Reader
	GetDomain() string
	GetHttpMethod() string
	GetParams() map[string]string
	GetPath() string
	GetService() string
	GetUrl() string
	GetVersion() string
	SetDomain(string)
	SetHttpMethod(string)
}

type BaseRequest struct {
	httpMethod string
	domain     string
	path       string
	params     map[string]string
	FormParams map[string]string

	service string
	version string
	action  string
}

func (r *BaseRequest) GetAction() string {
	return r.action
}

func (r *BaseRequest) GetHttpMethod() string {
	return r.httpMethod
}

func (r *BaseRequest) GetParams() map[string]string {
	return r.params
}

func (r *BaseRequest) GetPath() string {
	return r.path
}

func (r *BaseRequest) GetDomain() string {
	return r.domain
}

func (r *BaseRequest) SetDomain(domain string) {
	r.domain = domain
}

func (r *BaseRequest) SetHttpMethod(method string) {
	switch strings.ToUpper(method) {
	case POST:
		{
			r.httpMethod = POST
		}
	case GET:
		{
			r.httpMethod = GET
		}
	default:
		{
			r.httpMethod = GET
		}
	}
}

func (r *BaseRequest) GetService() string {
	return r.service
}

func (r *BaseRequest) GetUrl() string {
	if r.httpMethod == GET {
		return "https://" + r.domain + r.path + "?" + getUrlQueriesEncoded(r.params)
	} else if r.httpMethod == POST {
		return "https://" + r.domain + r.path
	} else {
		return ""
	}
}

func (r *BaseRequest) GetVersion() string {
	return r.version
}

func getUrlQueriesEncoded(params map[string]string) string {
	values := url.Values{}
	for key, value := range params {
		if value != "" {
			values.Add(key, value)
		}
	}
	return values.Encode()
}

func (r *BaseRequest) GetBodyReader() io.Reader {
	if r.httpMethod == POST {
		s := getUrlQueriesEncoded(r.params)
		log.Printf("[DEBUG] body: %s", s)
		return strings.NewReader(s)
	} else {
		return strings.NewReader("")
	}
}

func (r *BaseRequest) Init() *BaseRequest {
	r.httpMethod = GET
	r.domain = ""
	r.path = V2Path
	r.params = make(map[string]string)
	r.FormParams = make(map[string]string)
	return r
}

func (r *BaseRequest) WithApiInfo(service, version, action string) *BaseRequest {
	r.service = service
	r.version = version
	r.action = action
	return r
}

func GetServiceDomain(service string) (domain string) {
	domain = service + "." + RootDomain
	return
}

func CompleteCommonParams(request Request, c *Client) {
	params := request.GetParams()
	params["Region"] = c.GetRegion()
	params["Version"] = request.GetVersion()
	params["Action"] = request.GetAction()
	params["Timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	params["Nonce"] = strconv.Itoa(rand.Int())
}

func ConstructParams(req Request) (err error) {
	value := reflect.ValueOf(req).Elem()
	err = flatStructure(value, req, "")
	log.Printf("[DEBUG] params=%s", req.GetParams())
	return
}

func flatStructure(value reflect.Value, request Request, prefix string) (err error) {
	valueType := value.Type()
	for i := 0; i < valueType.NumField(); i++ {
		tag := valueType.Field(i).Tag
		log.Printf("[DEBUG] tag=%s, type=%s, value=%s", tag, valueType.Field(i), value.Field(i))
		nameTag, hasNameTag := tag.Lookup("name")
		if !hasNameTag {
			continue
		}
		_, hasListTag := tag.Lookup("list")
		if !hasListTag {
			key := prefix + nameTag
			v := value.Field(i)
			var vs string
			if v.Type().String() == "string" {
				vs = v.String()
			} else if v.Type().String() == "int" {
				vs = strconv.FormatInt(v.Int(), 10)
			}
			request.GetParams()[key] = vs
		} else {
			list := value.Field(i)
			if list.Kind() != reflect.Slice {
				list = list.Elem()
			}
			if !list.IsValid() || list.IsNil() {
				continue
			}
			for j := 0; j < list.Len(); j++ {
				vj := list.Index(j)
				key := prefix + nameTag + "." + strconv.Itoa(j)
				if vj.Type().String() == "string" {
					request.GetParams()[key] = vj.String()
				} else {
					err = flatStructure(vj, request, key+".")
					if err != nil {
						return
					}
				}
			}
		}
	}
	return
}
