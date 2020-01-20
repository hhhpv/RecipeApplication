package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	Dbase "../database"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

type UniqueUser struct {
	Name  string `json:"name"`
	Valid string `json:"valid"`
}

type Check struct {
	Username string `json:"username"`
}

type UserValidate struct {
	Username string
	Password string
}

func UniqueUserHandler(r *mux.Router) {
	r.HandleFunc("/uniname", UniqueUserHandlerFunc)
}

func UniqueUserHandlerFunc(w http.ResponseWriter, r *http.Request) {

	var p Check
	var result UserValidate

	collection := Dbase.Client.Database("recipe").Collection("users")

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filter := bson.D{{"username", p.Username}}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		userValidate := UniqueUser{
			Name:  p.Username,
			Valid: "a valid",
		}
		json.NewEncoder(w).Encode(userValidate)
	} else {
		userValidate := UniqueUser{
			Name:  p.Username,
			Valid: "an invalid",
		}
		json.NewEncoder(w).Encode(userValidate)
	}
}
