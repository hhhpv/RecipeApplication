package handlers

import (
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

type Page struct {
	Title string
	Body  []byte
}

func loadPage(title string) *Page {
	filename := title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	return &Page{Title: title, Body: body}
}

func ViewHandler(r *mux.Router) {
	r.HandleFunc("/", ViewHandlerFunc)
}

func ViewHandlerFunc(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len(""):]
	p := loadPage(title)
	t, _ := template.ParseFiles("./views/HomePage.html")
	t.Execute(w, p)
}
