package component

import (
	"fmt"
	"strings"
)

var (
	VideoControlsRates   = "rates"
	VideoControlsPlay    = "play"
	VideoControlsTime    = "time"
	VideoControlsProcess = "process"
	VideoControlsVolume  = "volume"

	VideoAspectRaioAuto = "auto"
	VideoAspectRaio4v3  = "4:3"
	VideoAspectRaio16v9 = "16:9"
)

type Video struct {
	JsonStr []string
}

func (p Video) ClassName(s string) Video {
	return p._KeyVal("className", s)
}

func (p Video) Src(s string) Video {
	return p._KeyVal("src", s)
}

func (p Video) IsLive(s bool) Video {
	return p._KeyVal("isLive", s)
}

func (p Video) Poster(s string) Video {
	return p._KeyVal("poster", s)
}

func (p Video) SplitPoster(s bool) Video {
	return p._KeyVal("splitPoster", s)
}

func (p Video) Muted(s bool) Video {
	return p._KeyVal("muted", s)
}

func (p Video) AutoPlay(s bool) Video {
	return p._KeyVal("autoPlay", s)
}

func (p Video) Rates(s ...string) Video {
	ss := strings.Join(s, ",")
	p._NotEmpty("rates", s)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"rates":[%s]`, ss))
	return p
}
func (p Video) AspectRatio(s ...string) Video {
	return p._KeyVal("aspectRatio", s)
}

func (p Video) Build() string {
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"type":"%s"`, "video"))
	return "{" + strings.Join(p.JsonStr, ",") + "}"
}

func (p Video) _KeyVal(k string, v interface{}) Video {
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

func (p Video) _NotEmpty(py, v interface{}) {
	switch v := v.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			panic(fmt.Errorf("Property:%s,Cannot hold empty values", py))
		}
	}
}

func NewVideo() Video {
	var c Video
	return c
}
