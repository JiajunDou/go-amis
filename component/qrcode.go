package component

import (
	"fmt"
	"strings"
)

const (
	QRCodeLevelL = "L"
	QRCodeLevelM = "M"
	QRCodeLevelQ = "Q"
	QRCodeLevelH = "H"
)

type QRCode struct {
	JsonStr []string
}

func (p *QRCode) ClassName(s string) *QRCode {
	return p._KeyVal("className", s)
}

func (p *QRCode) QrcodeClassName(s string) *QRCode {
	return p._KeyVal("qrcodeClassName", s)
}

func (p *QRCode) CodeSize(s int) *QRCode {
	return p._KeyVal("codeSize", s)
}

func (p *QRCode) BackgroundColor(s string) *QRCode {
	return p._KeyVal("backgroundColor", s)
}

func (p *QRCode) ForegroundColor(s string) *QRCode {
	return p._KeyVal("foregroundColor", s)
}

func (p *QRCode) Level(s string) *QRCode {
	return p._KeyVal("level", s)
}

func (p *QRCode) Value(s string) *QRCode {
	return p._KeyVal("value", s)
}

func (p *QRCode) Build() string {
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"type":"%s"`, "qr-code"))
	return "{" + strings.Join(p.JsonStr, ",") + "}"
}

func (p *QRCode) _KeyVal(k string, v interface{}) *QRCode {
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

func (p *QRCode) _NotEmpty(py, v interface{}) {
	switch v := v.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			panic(fmt.Errorf("Property:%s,Cannot hold empty values", py))
		}
	}
}

func NewQRCode() *QRCode {
	c := new(QRCode)
	return c
}
