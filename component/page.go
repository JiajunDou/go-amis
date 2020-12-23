package component

import (
	"fmt"
	"strings"
)

type Page struct {
	JsonStr []string
}

func (p Page) Title(s ...string) Page {
	ss := strings.Join(s, ",")
	p._NotEmpty("title", ss)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"title":[%s]`, ss))
	return p
}

func (p Page) SubTitle(s ...string) Page {
	ss := strings.Join(s, ",")
	p._NotEmpty("subTitle", ss)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"subTitle":[%s]`, ss))
	return p
}

func (p Page) Remark(s string) Page {
	p._NotEmpty("remark", s)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"remark":%s`, s))
	return p
}

func (p Page) Aside(s ...string) Page {
	ss := strings.Join(s, ",")
	p._NotEmpty("aside", ss)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"aside":[%s]`, ss))
	return p
}

func (p Page) Toolbar(s ...string) Page {
	ss := strings.Join(s, ",")
	p._NotEmpty("toolbar", ss)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"toolbar":[%s]`, ss))
	return p
}

func (p Page) Body(s ...string) Page {
	ss := strings.Join(s, ",")
	p._NotEmpty("body", ss)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"body":[%s]`, ss))
	return p
}

func (p Page) ClassName(s string) Page {
	return p._KeyVal("className", s)
}

func (p Page) ToolbarClassName(s string) Page {
	return p._KeyVal("toolbarClassName", s)
}

func (p Page) BodyClassName(s string) Page {
	return p._KeyVal("bodyClassName", s)
}

func (p Page) AsideClassName(s string) Page {
	return p._KeyVal("asideClassName", s)
}

func (p Page) HeaderClassName(s string) Page {
	return p._KeyVal("headerClassName", s)
}

func (p Page) InitApi(s string) Page {
	return p._KeyVal("initApi", s)
}

func (p Page) InitFetch(s bool) Page {
	return p._KeyVal("initFetch", s)
}

func (p Page) InitFetchOn(s string) Page {
	return p._KeyVal("initFetchOn", s)
}

func (p Page) Interval(s int) Page {
	return p._KeyVal("interval", s)
}

func (p Page) SilentPolling(s bool) Page {
	return p._KeyVal("silentPolling", s)
}

func (p Page) StopAutoRefreshWhen(s bool) Page {
	return p._KeyVal("stopAutoRefreshWhen", s)
}

func (p Page) Build() string {
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"type":"%s"`, "page"))
	return "{" + strings.Join(p.JsonStr, ",") + "}"
}

func (p Page) _KeyVal(k string, v interface{}) Page {
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

func (p Page) _NotEmpty(py, v interface{}) {
	switch v := v.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			panic(fmt.Errorf("Property:%s,Cannot hold empty values", py))
		}
	}
}

func NewPage() Page {
	var c Page
	return c
}
