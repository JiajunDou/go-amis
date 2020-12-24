package main

import (
	"encoding/json"
	"github.com/JiajunDou/go-amis/component"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	AmisPath = getParentDirectory(getCurrentDirectory()) + "/amis"
	TmplFile = AmisPath + "/amis.html"
)

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func tmpl(w http.ResponseWriter, r *http.Request) {

	t1, err := template.ParseFiles(TmplFile)
	if err != nil {
		panic(err)
	}
	_ = t1.Execute(w, struct{ Data string }{Data: Data()})
}

func Data() string {
	return component.NewPage().
		Title(
			component.Tpl("Big Title"),
		).
		Remark(
			component.NewRemark().
				Title("title").
				Content("Write something here").
				Icon("question-mark").
				Placement(component.RemarkPlacementBottom).
				Trigger(component.RemarkTriggerHover).
				RootClose(true).
				Build(),
		).
		Body(
			component.NewService().
				Api("/basicTable").
				Body(
					component.NewTable().
						Title("动态数据填充的表格，分页等其他特性请参照表格的CRUD演示").
						Source("$rows").
						// 显示列筛选器
						ColumnsTogglable(true).
						Columns(
							component.NewTableColumns().
								TypeText().
								Name("id").
								Label("ID(已手动设置)").
								// 隐藏该列，请从列筛选器里面打勾显示
								Toggled(true).
								Build(),
							component.NewTableColumns().
								TypeImage().
								Name("image").
								Label("图片/列宽20%").
								WidthInPercentage("20%").
								Build(),
							component.NewTableColumns().
								TypeDate().
								Name("date").
								Label("日期/列宽200px").
								WidthInPixel(200).
								Build(),
							component.NewTableColumns().
								TypeProgress().
								Name("progress").
								Label("进度条").
								Build(),
							component.NewTableColumns().
								TypeStatus().
								Name("status").
								Label("状态").
								Build(),
							component.NewTableColumns().
								TypeSwitch().
								Name("switch").
								Label("开关").
								Build(),
							component.NewTableColumns().
								CustomTypeLink(
									component.NewLink().
										Body("这个是一个链接").
										Href("${link | raw}${id}").
										Blank(true).
										Build(),
								).
								Name("link").
								Label("链接").
								CopyableWithContent("复制的内容是:${link | raw}").
								Build(),
							component.NewTableColumns().
								TypeColor().
								Name("color").
								Label("颜色").
								Copyable(true).
								Build(),
						).
						Build(),
				).
				Build(),
		).
		Aside(
			component.Tpl("Sidebar area"),
		).
		Toolbar(
			component.Tpl("Toolbar area"),
		).
		Build()
}

type BasicTableData struct {
	Data   TableData `json:"data"`
	Msg    string    `json:"msg"`
	Status int       `json:"status"`
}

type TableData struct {
	//
	Items   []Items `json:"rows"`
	Total   int     `json:"total,omitempty"`
	HasNext bool    `json:"haNext,omitempty"`
}

type Items struct {
	ID           int         `json:"id"`
	TypeText     string      `json:"text,omitempty"`
	TypeImage    string      `json:"image,omitempty"`
	TypeDate     string      `json:"date,omitempty"`
	TypeProgress int         `json:"progress,omitempty"`
	TypeStatus   bool        `json:"status"`
	TypeSwitch   bool        `json:"switch"`
	TypeLink     string      `json:"link,omitempty"`
	TypeColor    string      `json:"color,omitempty"`
	Children     interface{} `json:"children,omitempty"`
}

func BasicTable(w http.ResponseWriter, r *http.Request) {
	tableData := BasicTableData{
		Data: TableData{
			Items: []Items{
				{
					ID:           1,
					TypeText:     "Text",
					TypeImage:    "https://ss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=3893101144,2877209892&fm=23&gp=0.jpg",
					TypeDate:     "1691270438",
					TypeProgress: 22,
					TypeStatus:   true,
					TypeSwitch:   true,
					TypeLink:     "https://www.baidu.com/",
					TypeColor:    "#108cee",
					// Children: Items{
					// 	ID:           6,
					// 	TypeText:     "Text",
					// 	TypeImage:    "https://ss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=3893101144,2877209892&fm=23&gp=0.jpg",
					// 	TypeDate:     "1691270438",
					// 	TypeProgress: 22,
					// 	TypeStatus:   true,
					// 	TypeSwitch:   true,
					// 	TypeLink:     "https://www.baidu.com/",
					// 	TypeColor:    "#108cee",
					// },
				},
				{
					ID:           2,
					TypeText:     "Text",
					TypeImage:    "https://ss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=3893101144,2877209892&fm=23&gp=0.jpg",
					TypeDate:     "1681270438",
					TypeProgress: 40,
					TypeStatus:   true,
					TypeSwitch:   false,
					TypeLink:     "https://www.baidu.com/",
					TypeColor:    "#f38900",
				},
				{
					ID:           3,
					TypeText:     "Text",
					TypeImage:    "https://ss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=3893101144,2877209892&fm=23&gp=0.jpg",
					TypeDate:     "1671270438",
					TypeProgress: 60,
					TypeStatus:   false,
					TypeSwitch:   true,
					TypeLink:     "https://www.baidu.com/",
					TypeColor:    "#04c1ba",
				},
				{
					ID:           4,
					TypeText:     "Text",
					TypeImage:    "https://ss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=3893101144,2877209892&fm=23&gp=0.jpg",
					TypeDate:     "1661270438",
					TypeProgress: 80,
					TypeStatus:   false,
					TypeSwitch:   false,
					TypeLink:     "https://www.baidu.com/",
					TypeColor:    "#108cee",
				},
				{
					ID:           5,
					TypeText:     "Text",
					TypeImage:    "https://ss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=3893101144,2877209892&fm=23&gp=0.jpg",
					TypeDate:     "1651270438",
					TypeProgress: 100,
					TypeStatus:   true,
					TypeSwitch:   true,
					TypeLink:     "https://www.baidu.com/",
					TypeColor:    "#f38900",
				},
			},
			Total: 6,
		},
		Msg:    "ok",
		Status: 0,
	}
	data, _ := json.Marshal(&tableData)
	_, _ = w.Write(data)

}

func main() {
	fs := http.FileServer(http.Dir(AmisPath))
	http.Handle("/amis/", http.StripPrefix("/amis/", fs))
	http.HandleFunc("/", tmpl)
	http.HandleFunc("/basicTable", BasicTable)
	_ = http.ListenAndServe(":8080", nil)
}
