package handlers

import (
	"fmt"
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
	r.HandleFunc("/", CheckSecurity(ViewHandlerFunc))
}

func ViewHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Cookie)
	title := r.URL.Path[len(""):]
	p := loadPage(title)
	t, _ := template.ParseFiles("./views/HomePage.html")
	t.Execute(w, p)
}
