package sauce

import (
	"net/http"
	"github.com/go2rust1/cherry/src/channel"
	"github.com/go2rust1/cherry/src/dbms"
	"github.com/go2rust1/cherry/src/resource"
	"github.com/go2rust1/cherry/src/trait"
)

type sauce struct {
	name string
}

func (s *sauce) Name() string {
	return s.name
}

func (s *sauce) Request(url string, parser trait.Parser, meta trait.Meta) {
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	s.Requests(request, parser, meta)
}

func (s *sauce) Requests(request *http.Request, parser trait.Parser, meta trait.Meta) {
	_request := resource.NewRequest(s, request, parser, meta)
	channel.RequestIn(_request)
}

func (s *sauce) Bind(db ...trait.Database) {
	dbms.Bind(s.name, db...)
}

func (s *sauce) Send(model interface{}) {
	dbms.Send(s.name, model)
}

func New(name string) trait.Topic {
	return &sauce{name: name}
}
