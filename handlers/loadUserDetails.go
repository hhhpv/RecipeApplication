package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	Dbase "../database"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

type LoadRequest struct {
	Name  string `json:"username"`
	Token string `json:"token"`
}

type UserDetailResponse struct {
	Result   string `json:"response"`
	Name     string `json:"username"`
	Location string `json:"location"`
	About    string `json:"about"`
}

type FetchUserDetailFromDatabase struct {
	Name     string `json:"name"`
	Token    string `json:"token"`
	Location string `json:"location"`
	About    string `json:"about"`
}

func LoadUserDetailsHandler(r *mux.Router) {
	r.HandleFunc("/load_details", LoadUserDetailsHandlerFunc)
}

func LoadUserDetailsHandlerFunc(w http.ResponseWriter, r *http.Request) {

	var req LoadRequest
	var query FetchUserDetailFromDatabase

	collection := Dbase.Client.Database("recipe").Collection("user_info")

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&req)
	fmt.Println("debug")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if ParseToken(req.Token, MySigningKey) {
		filter := bson.D{{"name", req.Name}}
		err = collection.FindOne(context.TODO(), filter).Decode(&query)

		if err != nil {
			userDetail := UserDetailResponse{
				Result:   "failure",
				Name:     query.Name,
				Location: "",
				About:    "",
			}
			json.NewEncoder(w).Encode(userDetail)

		} else {
			userDetail := UserDetailResponse{Result: "success",
				Name:     query.Name,
				Location: query.Location,
				About:    query.About,
			}
			json.NewEncoder(w).Encode(userDetail)
		}
	} else {
		userDetail := UserDetailResponse{
			Result:   "unauthorized",
			Name:     "",
			Location: "",
			About:    "",
		}
		json.NewEncoder(w).Encode(userDetail)
	}

}
