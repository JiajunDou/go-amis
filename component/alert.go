package component

import (
	"fmt"
	"strings"
)

const (
	LevelInfo    = "info"
	LevelSuccess = "success"
	LevelWarning = "warning"
	LevelDanger  = "danger"
)

type Alert struct {
	JsonStr []string
}

func (p *Alert) ClassName(s string) *Alert {
	return p._KeyVal("className", s)
}

func (p *Alert) Level(s string) *Alert {
	return p._KeyVal("level", s)
}

func (p *Alert) Body(s ...string) *Alert {
	ss := strings.Join(s, ",")
	p._NotEmpty("body", ss)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"body":[%s]`, ss))
	return p
}

func (p *Alert) ShowCloseButton(s bool) *Alert {
	return p._KeyVal("showCloseButton", s)
}

func (p *Alert) Build() string {
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"type":"%s"`, "alert"))
	return "{" + strings.Join(p.JsonStr, ",") + "}"
}

func (p *Alert) _KeyVal(k string, v interface{}) *Alert {
	p._NotEmpty(k, v)
	switch v.(type) {
	case string:
		p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"%s":"%s"`, k, v))
	case int:
		p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"%s":%d`, k, v))
	case bool:
		p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"%s":%t`, k, v))
	}
	return p
}

func (p *Alert) _NotEmpty(py, v interface{}) {
	switch v := v.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			panic(fmt.Errorf("Property:%s,Cannot hold empty values", py))
		}
	}
}

func NewAlert() *Alert {
	c := new(Alert)
	return c
}
