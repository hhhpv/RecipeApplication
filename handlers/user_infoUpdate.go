package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	Dbase "../database"
	"github.com/gorilla/mux"
)

// type RespondUpdate struct {
// 	Name   string `json:"name"`
// 	Token  string `json:"token"`
// 	Result string `json:"result"`
// }

type ReqUserUpdate struct {
	Name     string `json:"username"`
	Token    string `json:"token"`
	Location string `json:"location"`
	About    string `json:"about"`
}
type UpdateUser struct {
	Name     string `json:"username"`
	Token    string `json:"token"`
	Location string `json:"location"`
	About    string `json:"about"`
}

func UserInfoUpdateHandler(r *mux.Router) {
	r.HandleFunc("/user_update", UserInfoUpdateHandlerFunc)
}

func UserInfoUpdateHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// var res RegisterUser
	var ret ReqUserUpdate
	collection := Dbase.Client.Database("recipe").Collection("user_info")
	w.Header().Set("Content-Type", "application/json")
	regerr := json.NewDecoder(r.Body).Decode(&ret)
	if regerr != nil {
		fmt.Println(r.Body)
		fmt.Println(regerr)
		http.Error(w, regerr.Error(), http.StatusBadRequest)
		return
	}
	if len(ret.Name) == 0 || len(ret.Location) == 0 || len(ret.About) == 0 {
		registerStatus := RegisterUser{
			Name:   ret.Name,
			Token:  "",
			Result: "empty",
		}
		json.NewEncoder(w).Encode(registerStatus)
	} else if ParseToken(ret.Token, MySigningKey) {
		update := UpdateUser{Name: ret.Name, Location: ret.Location, About: ret.About}
		insertResult, insertErr := collection.InsertOne(context.TODO(), update)
		if insertErr != nil {
			log.Fatal(insertErr)
			registerStatus := RegisterUser{
				Name:   ret.Name,
				Token:  "",
				Result: "failure",
			}
			json.NewEncoder(w).Encode(registerStatus)
		} else {
			fmt.Println("New User Created! ", insertResult.InsertedID)
			registerStatus := RegisterUser{
				Name:   ret.Name,
				Token:  ret.Token,
				Result: "success",
			}
			json.NewEncoder(w).Encode(registerStatus)
		}
	}

}
