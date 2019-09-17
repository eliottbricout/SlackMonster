package main

import (
	"./models"
	"./service"
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)


func main() {
	url := "mongodb+srv://slackog:slackog@cluster0-qwwtk.mongodb.net/test?	retryWrites=true&w=majority"
	client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	database := client.Database("SlackMonster")
	servicePlayer :=  service.CreatePlayerService(*database)
	serviceChoice :=  service.CreateChoiceService(*database, servicePlayer)
	r := chi.NewRouter()
	r.Get("/info", servicePlayer.GetPlayerRest)
	r.Post("/choice", serviceChoice.PostChoiceRest)
	http.ListenAndServe(":3000", r)
}


func test() {
	player := models.CreatePlayer("eliott")
	fmt.Println(player.Infos())
	player = player.Deck()[0].PowerRoom(player)
	fmt.Println(player.Infos())
	player = player.Deck()[0].PowerRoom(player)
	fmt.Println(player.Infos())
	player = player.Deck()[0].MalusRoom(player)
	fmt.Println(player.Infos())
	player = player.Deck()[0].PowerRoom(player)
	fmt.Println(player.Infos())
}
