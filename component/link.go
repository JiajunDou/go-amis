package component

import (
	"fmt"
	"strings"
)

type Link struct {
	JsonStr []string
}

func (p Link) ClassName(s string) Link {
	return p._KeyVal("className", s)
}

// 链接文本
func (p Link) Body(s string) Link {
	return p._KeyVal("body", s)
}

// 链接地址
func (p Link) Href(s string) Link {
	return p._KeyVal("href", s)
}

// 是否在新标签页打开
func (p Link) Blank(s bool) Link {
	return p._KeyVal("blank", s)
}

// a标签的target
func (p Link) HtmlTarget(s string) Link {
	return p._KeyVal("htmlTarget", s)
}

func (p Link) Build() string {
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"type":"%s"`, "link"))
	return strings.Join(p.JsonStr, ",")
}

// Form 中静态展示
func (p Link) BuildStatic() string {
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"type":"%s"`, "static-link"))
	return strings.Join(p.JsonStr, ",")
}

func (p Link) _KeyVal(k string, v interface{}) Link {
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

func (p Link) _NotEmpty(py, v interface{}) {
	switch v := v.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			panic(fmt.Errorf("Property:%s,Cannot hold empty values", py))
		}
	}
}

func NewLink() Link {
	var c Link
	return c
}
