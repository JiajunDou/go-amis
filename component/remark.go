package component

import (
	"fmt"
	"strings"
)

var (
	RemarkPlacementTop    = "top"
	RemarkPlacementRight  = "right"
	RemarkPlacementLeft   = "left"
	RemarkPlacementBottom = "bottom"

	RemarkTriggerHover = "hover"
	RemarkTriggerClick = "click"
	RemarkTriggerFocus = "focus"
)

type Remark struct {
	JsonStr []string
}

func (r Remark) ClassName(s string) Remark {
	return r._KeyVal("className", s)
}

func (r Remark) Title(s string) Remark {
	return r._KeyVal("title", s)
}

func (r Remark) Content(s string) Remark {
	return r._KeyVal("content", s)
}

func (r Remark) Placement(s string) Remark {
	return r._KeyVal("placement", s)
}

func (r Remark) Trigger(s string) Remark {
	return r._KeyVal("trigger", s)
}

func (r Remark) Icon(s string) Remark {
	return r._KeyVal("icon", s)
}

func (r Remark) RootClose(s bool) Remark {
	return r._KeyVal("rootClose", s)
}

func (r Remark) Build() string {
	r.JsonStr = append(r.JsonStr, fmt.Sprintf(`"type":"%s"`, "remark"))
	return "{" + strings.Join(r.JsonStr, ",") + "}"
}

func (r Remark) NoTypeBuild() string {
	return "{" + strings.Join(r.JsonStr, ",") + "}"
}

func (r Remark) _KeyVal(k string, v interface{}) Remark {
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

func (r Remark) _NotEmpty(py, v interface{}) {
	switch v := v.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			panic(fmt.Errorf("Property:%s,Cannot hold empty values", py))
		}
	}
}

func NewRemark() Remark {
	var c Remark
	return c
}
