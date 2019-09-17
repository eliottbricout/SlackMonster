package repository

import (
	"../models"
	"context"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlayerRepository struct {
	playerCollection *mongo.Collection
}

func CreatePlayerRepository(database mongo.Database) PlayerRepository {
	playerCollection := database.Collection("Player")
	return PlayerRepository{playerCollection}
}

func (repo *PlayerRepository) GetPlayer(name string) models.Player {
	filter := bson.M{"name": name}
	var player models.PlayerPivot
	repo.playerCollection.FindOne(context.TODO(), filter).Decode(&player)
	return player.TransformPlayer()
}

func (repo *PlayerRepository) AddPlayer(player models.PlayerPivot) {
	repo.playerCollection.InsertOne(context.TODO(), player)
}