package component

import (
	"fmt"
	"strings"
)

type Image struct {
	JsonStr []string
}

func (p *Image) ClassName(s string) *Image {
	return p._KeyVal("className", s)
}

// *Image CSS
func (p Image) ImageClassName(s string) *Image {
	return p._KeyVal("imageClassName", s)
}

// title
func (p *Image) Title(s string) *Image {
	return p._KeyVal("title", s)
}

// imageCaption
// 就是图片描述
func (p *Image) ImageCaption(s string) *Image {
	return p._KeyVal("imageCaption", s)
}

// Placeholder
// 占位符
func (p *Image) Placeholder(s string) *Image {
	return p._KeyVal("placeholder", s)
}

// Default*ImageURL
// 默认显示的图片地址，应该是图片没有或者找不到的时候
// TODO 这个要测试一下
func (p *Image) DefaultImage(s string) *Image {
	return p._KeyVal("defaultImage", s)
}

// Thumb URL
// 图片缩略图地址
func (p *Image) Src(s string) *Image {
	return p._KeyVal("src", s)
}

// Origin URL
// 图片原图地址
func (p *Image) OriginalSrc(s string) *Image {
	return p._KeyVal("originalSrc", s)
}

// 放大预览
func (p *Image) EnlargeAble(s bool) *Image {
	return p._KeyVal("enlargeAble", s)
}

// 放大预览时候的标题
func (p *Image) EnlargeTitle(s string) *Image {
	return p._KeyVal("enlargeTitle", s)
}

// 放大预览时候的描述
func (p *Image) EnlargeCaption(s string) *Image {
	return p._KeyVal("enlargeCaption", s)
}

const (
	ImageThumbModeWidthFull = "w-full"
	ImageThumbModeHighFull  = "h-full"
	ImageThumbModeContain   = "contain"
	ImageThumbModeCover     = "cover"
)

// 缩略图模式
func (p *Image) ThumbMode(imageThumbMode string) *Image {
	return p._KeyVal("thumbMode", imageThumbMode)
}

var (
	ImageThumbRatio1v1  = "1:1"
	ImageThumbRatio4v3  = "4:3"
	ImageThumbRatio16v9 = "16:9"
)

// 缩略图模式
func (p *Image) ThumbRatio(imageThumbMode string) *Image {
	return p._KeyVal("thumbRatio", imageThumbMode)
}

func (p *Image) Build() string {
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"type":"%s"`, "image"))
	return strings.Join(p.JsonStr, ",")
}

// Form 中静态展示
func (p *Image) BuildStatic() string {
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"type":"%s"`, "static-image"))
	return strings.Join(p.JsonStr, ",")
}

func (p *Image) _KeyVal(k string, v interface{}) *Image {
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

func (p *Image) _NotEmpty(py, v interface{}) {
	switch v := v.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			panic(fmt.Errorf("Property:%s,Cannot hold empty values", py))
		}
	}
}

func NewImage() *Image {
	c := new(Image)
	return c
}
