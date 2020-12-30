package component

import (
	"fmt"
	"strings"
)

type Progress struct {
	JsonStr []string
}

func (p *Progress) ClassName(s string) *Progress {
	return p._KeyVal("className", s)
}

// 进度条CSS
func (p *Progress) ProgressClassName(s string) *Progress {
	return p._KeyVal("progressClassName", s)
}

// Used in static mode
// 进度值
func (p *Progress) Value(s string) *Progress {
	return p._KeyVal("value", s)
}

// Placeholder:-
// 占位符:-
func (p *Progress) Placeholder(s string) *Progress {
	return p._KeyVal("placeholder", s)
}

//是否展示进度文本
func (p *Progress) ShowLabel(s bool) *Progress {
	return p._KeyVal("showLabel", s)
}

// 进度条颜色 最多5等分
// 默认值是：'bg-danger', 'bg-warning', 'bg-info', 'bg-success', 'bg-success'
// TODO 这里可以var几个值
func (p *Progress) Map(progressMap ...string) *Progress {
	ss := strings.Join(progressMap, ",")
	p._NotEmpty("map", ss)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"map":[%s]`, ss))
	return p
}

func (p *Progress) Build() string {
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"type":"%s"`, "progress"))
	return strings.Join(p.JsonStr, ",")
}

// Form 中静态展示
func (p *Progress) BuildStatic() string {
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"type":"%s"`, "static-progress"))
	return strings.Join(p.JsonStr, ",")
}

func (p *Progress) _KeyVal(k string, v interface{}) *Progress {
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

func (p *Progress) _NotEmpty(py, v interface{}) {
	switch v := v.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			panic(fmt.Errorf("Property:%s,Cannot hold empty values", py))
		}
	}
}

func NewProgress() *Progress {
	c := new(Progress)
	return c
}
