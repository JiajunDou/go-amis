package component

import (
	"fmt"
	"strings"
)

type Switch struct {
	JsonStr []string
}

func (p Switch) ClassName(s string) Switch {
	return p._KeyVal("className", s)
}

// Used in static mode
// 0 : false 1: true
func (p Switch) Value(s bool) Switch {
	return p._KeyVal("value", s)
}

// 真值，当值为该值时，开关开启
func (p Switch) TrueValue(s int) Switch {
	return p._KeyVal("trueValue", s)
}

// 右侧选项文本
func (p Switch) Option(s string) Switch {
	return p._KeyVal("option", s)
}

// Placeholder:-
// 占位符:-
func (p Switch) Placeholder(s string) Switch {
	return p._KeyVal("placeholder", s)
}

func (p Switch) Build() string {
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"type":"%s"`, "switch"))
	return strings.Join(p.JsonStr, ",")
}

// Form 中静态展示
func (p Switch) BuildStatic() string {
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"type":"%s"`, "static-switch"))
	return strings.Join(p.JsonStr, ",")
}

func (p Switch) _KeyVal(k string, v interface{}) Switch {
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

func (p Switch) _NotEmpty(py, v interface{}) {
	switch v := v.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			panic(fmt.Errorf("Property:%s,Cannot hold empty values", py))
		}
	}
}

func NewSwitch() Switch {
	var c Switch
	return c
}
