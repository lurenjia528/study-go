//go:generate statik -src=./public

package main

import (
	"strings"
	"log"
	"net/http"
	"github.com/lxn/walk"
	"github.com/rakyll/statik/fs"
	. "github.com/lxn/walk/declarative"
	_ "GoStudy/gui/statikwindows/statik"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(statikFS)))
		http.ListenAndServe(":8080", nil)
	}()
	//var le *walk.LineEdit
	var wv *walk.WebView
	MainWindow{
		//AssignTo: &mw,
		Icon:     Bind("'../img/' + icon(wv.URL) + '.ico'"),
		Title:    "Walk WebView Example'",
		//MinSize: Size{940, 300},
		Size:   Size{1050, 500},
		Layout: VBox{MarginsZero: true},
		Children: []Widget{
			//LineEdit{
			//	AssignTo: &le,
			//	Text:     Bind("wv.URL"),
			//	OnKeyDown: func(key walk.Key) {
			//		if key == walk.KeyReturn {
			//			wv.SetURL("http://127.0.0.1:8080/public/html/hw.html")
			//		}
			//	},
			//},
			WebView{
				AssignTo: &wv,
				Name:     "wv",
				URL:      "http://127.0.0.1:8080/public/html/hw.html",
			},
		},
		Functions: map[string]func(args ...interface{}) (interface{}, error){
			"icon": func(args ...interface{}) (interface{}, error) {
				if strings.HasPrefix(args[0].(string), "https") {
					return "check", nil
				}

				return "stop", nil
			},
		},
	}.Run()

}
