package main

import (
	"context"
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/globalsign/mgo/bson"
	"time"

	"./models"
)

func main() {
	url := "mongodb+srv://slackog:slackog@cluster0-qwwtk.mongodb.net/test?	retryWrites=true&w=majority"
	ctx, Cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer Cancel()

	client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI(url))

	db := client.Database("SlackMonster")
	playerCollection := db.Collection("Player")
	var player models.PlayerPivot
	filter := bson.M{"name": "eliott"}
	cur := playerCollection.FindOne(ctx, filter).Decode(&player)
	p, err := playerCollection.Find(ctx, bson.M{})
	p.Next(context.Background())
	fmt.Println(p.Current)
	fmt.Println(err)

	fmt.Println(client)
	fmt.Println(cur.Error())
	fmt.Println(player)


	r := chi.NewRouter()
	r.Get("/info", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(""))
	})
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
