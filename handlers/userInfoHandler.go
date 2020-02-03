package handlers

import (
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func UserInfoHandler(r *mux.Router) {
	r.HandleFunc("/user_info", CheckSecurity(UserInfoHandlerFunc))
}

func UserInfoHandlerFunc(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len(""):]
	p := loadPageUserInfo(title)
	t, _ := template.ParseFiles("./views/user_info.html")
	t.Execute(w, p)
}
func loadPageUserInfo(title string) *UserInfoPage {
	filename := title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	return &UserInfoPage{Title: title, Body: body}
}

type UserInfoPage struct {
	Title string
	Body  []byte
}
