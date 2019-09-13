package main

import (
	"context"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"

	"./models"
)

type PlayerPivot struct {
	Life int
	Name string
	Deck []int
	Graveyard []int
}

func main() {
	url := "mongodb+srv://slackog:slackog@cluster0-qwwtk.mongodb.net/test?	retryWrites=true&w=majority"
	ctx, Cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer Cancel()

	client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI(url))

	db := client.Database("SlackMonster")
	playerCollection := db.Collection("Player")
	play := PlayerPivot{1, "fab", []int{1,2,3,4,5},[]int{}}

	fmt.Println(play)
	eee, tt := playerCollection.InsertOne(context.TODO(), play)
	if tt != nil {
		log.Fatal(tt)
	}
	fmt.Println("eee")
	fmt.Println(eee)
	fmt.Println(eee.InsertedID)


	var player PlayerPivot
	filter := bson.M{"name": "fab"}
	cur := playerCollection.FindOne(ctx, filter)
	p, err := playerCollection.Find(ctx, bson.M{})
	p.Next(context.Background())
	fmt.Println(p.Current)

	fmt.Println(client)
	x, _ := cur.DecodeBytes()
	fmt.Println(x.String())

	if err != playerCollection.FindOne(ctx, filter).Decode(&player) {
		log.Fatal(err)
	} else {
		fmt.Println(player)
	}



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
