package component

import (
	"fmt"
	"strings"
)

type Date struct {
	JsonStr []string
}

func (p Date) ClassName(s string) Date {
	return p._KeyVal("className", s)
}

// Format UnixTime as YYYY-MM-DD
// 格式化时间戳：YYYY-MM-DD
func (p Date) Format(s string) Date {
	return p._KeyVal("format", s)
}

// Specifies the time format
// 日期格式不是时间戳的时候，告知当前使用的日期格式
func (p Date) ValueFormat(s string) Date {
	return p._KeyVal("valueFormat", s)
}

// Used in static mode
// 静态时间数据
func (p Date) Value(s string) Date {
	return p._KeyVal("value", s)
}

// Placeholder:-
// 占位符:-
func (p Date) Placeholder(s string) Date {
	return p._KeyVal("placeholder", s)
}

// TODO How to use?
func (p Date) FromNow(s bool) Date {
	return p._KeyVal("fromNow", s)
}

// TODO How to use?
func (p Date) UpdateFrequency(s bool) Date {
	return p._KeyVal("updateFrequency", s)
}

func (p Date) Build() string {
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"type":"%s"`, "date"))
	return strings.Join(p.JsonStr, ",")
}

// Form 中静态展示
func (p Date) BuildStatic() string {
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"type":"%s"`, "static-date"))
	return strings.Join(p.JsonStr, ",")
}

func (p Date) _KeyVal(k string, v interface{}) Date {
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

func (p Date) _NotEmpty(py, v interface{}) {
	switch v := v.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			panic(fmt.Errorf("Property:%s,Cannot hold empty values", py))
		}
	}
}

func NewDate() Date {
	var c Date
	return c
}
