package util

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var funcMap = template.FuncMap{}

// MakeTemplates
func MakeTemplates(path string) *Template {
	return template.Must(template.New("*").Funcs(funcMap).ParseGlob(path))
}

//FuncMap adds a grouping of functions to the func map
func FuncMap(fnMap map[string]any) template.FuncMap {
	for k, v := range fnMap {
		funcMap[k] = v
	}
	return funcMap
}

//HandleSigInt does what it says on the tin
func HandleSigInt(fn func()) {
	log.Println("Please press ctrl+c to exit.")
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fn()
		os.Exit(1)
	}()
}

//RenderTemplate does what it says on the tin
func RenderTemplate(w http.ResponseWriter, tmpl string, page Page) {
	buffer := GetBuf()
	err := templates.ExecuteTemplate(buffer, tmpl+".html", page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	buffer.WriteTo(w)
	PutBuf(buffer)
}
