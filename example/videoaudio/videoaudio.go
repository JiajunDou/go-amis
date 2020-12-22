package main

import (
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
	t1.Execute(w, struct{ Data string }{Data: Data()})
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
			component.NewVideo().
				Src("https://media.w3.org/2010/05/sintel/trailer_hd.mp4").
				Poster("https://video-react.js.org/assets/poster.png").
				Build(),
			component.NewAudio().
				Src("https://amis.bj.bcebos.com/amis/2019-7/1562137295708/chicane-poppiholla-original-radio-edit%20(1).mp3").
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

func main() {
	fs := http.FileServer(http.Dir(AmisPath))
	http.Handle("/amis/", http.StripPrefix("/amis/", fs))
	http.HandleFunc("/", tmpl)
	http.ListenAndServe(":8080", nil)
}
