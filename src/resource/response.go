package resource

import (
	"github.com/go2rust1/cherry/src/trait"
	"unsafe"
)

type _response struct {
	request trait.Request
	body    []byte
}

func (r *_response) Request() trait.Request {
	return r.request
}

func (r *_response) Body() []byte {
	return r.body
}

func (r *_response) Text() string {
	return *(*string)(unsafe.Pointer(&r.body))
}

func (r *_response) Meta() trait.Meta {
	return r.request.Meta()
}

func NewResponse(request trait.Request, body []byte) trait.Response {
	return &_response{request: request, body: body}
}
