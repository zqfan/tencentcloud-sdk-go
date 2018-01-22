package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Response interface {
	GetHttpContentBytes() []byte
}

type BaseResponse struct {
	httpContentBytes []byte
}

func (r *BaseResponse) GetHttpContentBytes() []byte {
	return r.httpContentBytes
}

func ParseFromHttpResponse(hr *http.Response, response Response) (err error) {
	defer hr.Body.Close()
	body, err := ioutil.ReadAll(hr.Body)
	if err != nil {
		return
	}
	log.Printf("[DEBUG] Response Body=%s", body)
	err = json.Unmarshal(body, &response)
	return
}
