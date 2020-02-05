package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	Dbase "../database"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

type RegisterUser struct {
	Name   string `json:"name"`
	Token  string `json:"token"`
	Result string `json:"result"`
}

type RegFields struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserExists struct {
	Result string `json:"result"`
}

func UserRegisterHandler(r *mux.Router) {
	r.HandleFunc("/user_register", UserRegisterHandlerFunc)
}

func UserRegisterHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var regFields RegFields
	var result UserValidate

	w.Header().Set("Content-Type", "application/json")

	regerr := json.NewDecoder(r.Body).Decode(&regFields)
	if regerr != nil {
		http.Error(w, regerr.Error(), http.StatusBadRequest)
		return
	}

	createdToken, err := GenerateToken([]byte(MySigningKey), regFields.Username)
	if err != nil {
		fmt.Println("Creating token failed")
	}

	collection := Dbase.Client.Database("recipe").Collection("users")

	filter := bson.D{{"username", regFields.Username}}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {

		err := r.ParseForm()
		if err != nil {
			fmt.Println("cannot parse")
		}

		name := regFields.Username
		password := regFields.Password

		newUserHashedPassword := hashAndSalt([]byte(password), len(password))
		newUser := Signin{name, newUserHashedPassword}
		insertResult, insertErr := collection.InsertOne(context.TODO(), newUser)
		if insertErr != nil {
			log.Fatal(insertErr)
			registerStatus := RegisterUser{
				Name:   name,
				Token:  "",
				Result: "failure",
			}
			json.NewEncoder(w).Encode(registerStatus)
		} else {
			registerStatus := RegisterUser{
				Name:   name,
				Token:  createdToken,
				Result: "success",
			}
			json.NewEncoder(w).Encode(registerStatus)
		}
	} else {
		res := UserExists{
			Result: "User Exists",
		}
		json.NewEncoder(w).Encode(res)
	}

}
