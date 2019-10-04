package repository

import (
	"../models"
	"context"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type PlayerRepository struct {
	playerCollection *mongo.Collection
}

func CreatePlayerRepository(database mongo.Database) PlayerRepository {
	playerCollection := database.Collection("Player")
	return PlayerRepository{playerCollection}
}

func (repo *PlayerRepository) GetPlayer(name string) (models.Player, error) {
	filter := bson.M{"name":  bson.M{"$regex": name }}
	var player models.PlayerPivot
	err := repo.playerCollection.FindOne(context.TODO(), filter).Decode(&player)
	fmt.Println(player)
	if err != nil {
		return models.Player{}, err
	}
	return player.TransformPlayer(), nil
}

func (repo *PlayerRepository) GetPlayerById(id string) (models.Player, error) {
	filter := bson.M{"id": id}
	var player models.PlayerPivot
	fmt.Print("search")
	cursor, err := repo.playerCollection.Find(context.TODO(), filter)
	if err != nil {
		return models.Player{}, err
	}
	cursor.Next(context.TODO())
	if err := cursor.Decode(&player); err != nil {
		return models.Player{}, err
	}
	return player.TransformPlayer(), nil
}

func (repo *PlayerRepository) AddPlayer(player models.PlayerPivot) {
	repo.playerCollection.InsertOne(context.TODO(), player)
}
func (repo *PlayerRepository) GetAllPlayer() []models.Player{
	cursor, err := repo.playerCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return []models.Player{}
	}
	var players []models.Player
	for cursor.Next(context.Background()) {
		var player models.PlayerPivot
		if err := cursor.Decode(&player); err != nil {
			log.Fatal(err)
		}
		players = append(players, player.TransformPlayer())
	}

	return players
}

func (repo *PlayerRepository) UpdatePlayer(choice models.Choice) {
	filter := bson.M{"playerId": choice.PlayerId}
	repo.playerCollection.UpdateOne(context.TODO(), filter, choice)
}
