package component

import (
	"fmt"
	"strings"
)

type Status struct {
	JsonStr []string
}

func (p *Status) ClassName(s string) *Status {
	return p._KeyVal("className", s)
}

// Used in static mode
// 0 : false 1: true
func (p *Status) Value(s bool) *Status {
	return p._KeyVal("value", s)
}

// Placeholder:-
// 占位符:-
func (p *Status) Placeholder(s string) *Status {
	return p._KeyVal("placeholder", s)
}

func (p *Status) Build() string {
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"type":"%s"`, "status"))
	return strings.Join(p.JsonStr, ",")
}

// Form 中静态展示
func (p *Status) BuildStatic() string {
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"type":"%s"`, "static-status"))
	return strings.Join(p.JsonStr, ",")
}

func (p *Status) _KeyVal(k string, v interface{}) *Status {
	p._NotEmpty(k, v)
	switch v := v.(type) {
	case string:
		p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"%s":"%s"`, k, v))
	case int:
		p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"%s":%d`, k, v))
	case bool:
		p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"%s":%t`, k, v))
	}
	return p
}

func (p *Status) _NotEmpty(py, v interface{}) {
	switch v := v.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			panic(fmt.Errorf("Property:%s,Cannot hold empty values", py))
		}
	}
}

func NewStatus() *Status {
	c := new(Status)
	return c
}
