package handlers

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func AboutAppHandler(r *mux.Router) {
	r.HandleFunc("/about_app", CheckSecurity(AboutAppHandlerFunc))
}

func AboutAppHandlerFunc(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len(""):]
	p := loadPage(title)
	t, _ := template.ParseFiles("./views/aboutApp.html")
	t.Execute(w, p)
}
