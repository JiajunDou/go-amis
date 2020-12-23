package component

import (
	"fmt"
	"strings"
)

type Color struct {
	JsonStr []string
}

func (p Color) ClassName(s string) Color {
	return p._KeyVal("className", s)
}

// 显示的颜色值
func (p Color) Value(s string) Color {
	return p._KeyVal("value", s)
}

// 默认颜色值
func (p Color) DefaultColor(s string) Color {
	return p._KeyVal("defaultColor", s)
}

// 是否显示右边的颜色值
func (p Color) ShowValue(s bool) Color {
	return p._KeyVal("showValue", s)
}

func (p Color) Build() string {
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"type":"%s"`, "color"))
	return strings.Join(p.JsonStr, ",")
}

// Form 中静态展示
func (p Color) BuildStatic() string {
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"type":"%s"`, "static-color"))
	return strings.Join(p.JsonStr, ",")
}

func (p Color) _KeyVal(k string, v interface{}) Color {
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

func (p Color) _NotEmpty(py, v interface{}) {
	switch v := v.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			panic(fmt.Errorf("Property:%s,Cannot hold empty values", py))
		}
	}
}

func NewColor() Color {
	var c Color
	return c
}
