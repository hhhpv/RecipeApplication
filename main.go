package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	Handlers "./handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	Dbase "./database"
)

func main() {

	r := mux.NewRouter()
	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("images/"))))
	Handlers.ViewHandler(r)
	Handlers.LoginHandler(r)
	Handlers.SignUpHandler(r)
	Handlers.UserInfoHandler(r)
	Handlers.SignMeInHandler(r)
	Handlers.LoadUserDetailsHandler(r)
	Handlers.UniqueUserHandler(r)
	Handlers.UserRegisterHandler(r)
	Handlers.UserInfoUpdateHandler(r)
	Handlers.AddRecipeHandler(r)
	Handlers.AboutViewHandler(r)
	Handlers.NewRecipeHandler(r)
	Handlers.RetrieveRecipeHandler(r)
	Handlers.DisplayRecipeHandler(r)
	Handlers.DeleteRecipeHandler(r)
	Handlers.AboutAppHandler(r)

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	Dbase.Client, _ = mongo.Connect(context.TODO(), clientOptions)

	err := Dbase.Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to MongoDB")
	}

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}
	srv.ListenAndServe()

}
