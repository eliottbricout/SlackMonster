package main

import (
	"github.com/eliottbricout/SlackMonster/service"
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

func main() {
	url := "mongodb+srv://slackog:slackog@cluster0-qwwtk.mongodb.net/test?retryWrites=true&w=majority"
	urlHookSlack := "https://hooks.slack.com/services/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"
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("connect")
	}
	database := client.Database("SlackMonster")
	servicePlayer := service.CreatePlayerService(*database)
	serviceChoice := service.CreateChoiceService(*database, servicePlayer)
	serviceDay := service.CreateGameService(serviceChoice, servicePlayer, urlHookSlack)
	r := chi.NewRouter()
	r.Post("/info", servicePlayer.InfoPlayer)
	r.Post("/join", servicePlayer.JoinPlayer)
	r.Post("/check", servicePlayer.CheckPlayer)
	r.Post("/rooms", service.ListRoom)
	r.Post("/choose", serviceChoice.ChoiceRest)
	r.Post("/startParty", serviceDay.StartParty)
	r.Post("/nextDay", serviceDay.NextDay)
	http.ListenAndServe(":3000", r)
}
