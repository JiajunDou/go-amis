package component

import (
	"fmt"
	"strings"
)

type Service struct {
	JsonStr []string
}

func (r *Service) ClassName(s string) *Service {
	return r._KeyVal("className", s)
}

func (r *Service) Title(s string) *Service {
	return r._KeyVal("title", s)
}

func (r *Service) Api(s string) *Service {
	return r._KeyVal("api", s)
}

func (p *Service) Body(s ...string) *Service {
	ss := strings.Join(s, ",")
	p._NotEmpty("body", ss)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"body":[%s]`, ss))
	return p
}

func (r *Service) Build() string {
	r.JsonStr = append(r.JsonStr, fmt.Sprintf(`"type":"%s"`, "service"))
	return "{" + strings.Join(r.JsonStr, ",") + "}"
}

func (r *Service) NoTypeBuild() string {
	return "{" + strings.Join(r.JsonStr, ",") + "}"
}

func (r *Service) _KeyVal(k string, v interface{}) *Service {
	r._NotEmpty(k, v)
	switch v := v.(type) {
	case string:
		r.JsonStr = append(r.JsonStr, fmt.Sprintf(`"%s":"%s"`, k, v))
	case int:
		r.JsonStr = append(r.JsonStr, fmt.Sprintf(`"%s":%d`, k, v))
	case bool:
		r.JsonStr = append(r.JsonStr, fmt.Sprintf(`"%s":%t`, k, v))
	}
	return r
}

func (r *Service) _NotEmpty(py, v interface{}) {
	switch v := v.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			panic(fmt.Errorf("Property:%s,Cannot hold empty values", py))
		}
	}
}

func NewService() *Service {
	c := new(Service)
	return c
}
