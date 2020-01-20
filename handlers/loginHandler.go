package handlers

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func LoginHandler(r *mux.Router) {
	r.HandleFunc("/login", LoginHandlerFunc)
}

func LoginHandlerFunc(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len(""):]
	p := loadPage(title)
	t, _ := template.ParseFiles("./views/page.html")
	t.Execute(w, p)
}
