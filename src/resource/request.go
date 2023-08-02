package resource

import (
	"github.com/go2rust1/cherry/src/trait"
	"net/http"
)

type _request struct {
	topic   trait.Topic
	request *http.Request
	parser  trait.Parser
	meta    trait.Meta
}

func (r *_request) Topic() trait.Topic {
	return r.topic
}

func (r *_request) Request() *http.Request {
	return r.request
}

func (r *_request) Parser() trait.Parser {
	return r.parser
}

func (r *_request) Meta() trait.Meta {
	return r.meta
}

func (r *_request) Finger() string {
	return r.topic.Name() + "-> " + r.request.Method + " " + r.request.URL.String()
}

func NewRequest(topic trait.Topic, request *http.Request, parser trait.Parser, meta trait.Meta) trait.Request {
	return &_request{
		topic:   topic,
		request: request,
		parser:  parser,
		meta:    meta,
	}
}
