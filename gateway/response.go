package gateway

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"net/http"
)

type Response struct {
	*resty.Response
}

func NewResponse(resp *resty.Response) *Response {
	return &Response{
		Response: resp,
	}
}

func (r *Response) IsSuccessful() bool {
	return r.GetStatusCode() >= http.StatusOK && r.GetStatusCode() < http.StatusMultipleChoices
}

func (r *Response) GetStatusCode() int {
	return r.StatusCode()
}

func (r *Response) GetStatus() string {
	return r.Status()
}

func (r *Response) GetProto() string {
	return r.Proto()
}

func (r *Response) GetBody() string {
	return string(r.Body())
}

func (r *Response) GetJson(v interface{}) error {
	return json.Unmarshal(r.Body(), &v)
}
