package main

import (
	"./service"
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

func main() {
	url := "mongodb+srv://slackog:slackog@cluster0-qwwtk.mongodb.net/test?retryWrites=true&w=majority"
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("connect")
	}
	database := client.Database("SlackMonster")
	servicePlayer := service.CreatePlayerService(*database)
	serviceChoice := service.CreateChoiceService(*database, servicePlayer)
	r := chi.NewRouter()
	r.Post("/info", servicePlayer.InfoPlayer)
	r.Post("/join", servicePlayer.JoinPlayer)
	r.Post("/check", servicePlayer.CheckPlayer)
	r.Post("/rooms", service.ListRoom)
	r.Post("/choose", serviceChoice.ChoiceRest)
	http.ListenAndServe(":3000", r)
}
