package component

import (
	"fmt"
	"strings"
)

const (
	AudioControlsRates   = "rates"
	AudioControlsPlay    = "play"
	AudioControlsTime    = "time"
	AudioControlsProcess = "process"
	AudioControlsVolume  = "volume"
)

type Audio struct {
	JsonStr []string
}

func (p *Audio) ClassName(s string) *Audio {
	return p._KeyVal("className", s)
}

func (p *Audio) Inline(s bool) *Audio {
	return p._KeyVal("inline", s)
}

func (p *Audio) Src(s string) *Audio {
	return p._KeyVal("src", s)
}

func (p *Audio) Loop(s bool) *Audio {
	return p._KeyVal("loop", s)
}

func (p *Audio) AutoPlay(s bool) *Audio {
	return p._KeyVal("autoPlay", s)
}

func (p *Audio) Rates(s ...string) *Audio {
	ss := strings.Join(s, ",")
	p._NotEmpty("rates", ss)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"rates":[%s]`, ss))
	return p
}

func (p *Audio) Controls(s ...string) *Audio {
	ss := strings.Join(s, ",")
	p._NotEmpty("controls", ss)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"controls":[%s]`, ss))
	return p
}

func (p *Audio) Build() string {
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"type":"%s"`, "audio"))
	return "{" + strings.Join(p.JsonStr, ",") + "}"
}

func (p *Audio) _KeyVal(k string, v interface{}) *Audio {
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

func (p Audio) _NotEmpty(py, v interface{}) {
	switch v := v.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			panic(fmt.Errorf("Property:%s,Cannot hold empty values", py))
		}
	}
}

func NewAudio() *Audio {
	c := new(Audio)
	return c
}
