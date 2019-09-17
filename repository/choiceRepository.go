package repository

import (
	"../models"
	"context"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChoiceRepository struct {
	choiceCollection *mongo.Collection
}

func CreateChoiceRepository(database mongo.Database) ChoiceRepository {
	choiceCollection := database.Collection("Choice")
	return ChoiceRepository{choiceCollection}
}

func (repo *ChoiceRepository) GetChoice(name string) models.Choice {
	filter := bson.M{"playerName": name}
	var choice models.Choice
	repo.choiceCollection.FindOne(context.TODO(), filter).Decode(&choice)
	return choice
}

func (repo *ChoiceRepository) AddChoice(Choice models.Choice) {
	repo.choiceCollection.InsertOne(context.TODO(), Choice)
}