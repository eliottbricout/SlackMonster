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

func (repo *ChoiceRepository) GetChoice(id string) (models.Choice, error) {
	filter := bson.M{"playerid": id}
	var choice models.Choice
	err := repo.choiceCollection.FindOne(context.TODO(), filter).Decode(&choice)
	return choice, err
}

func (repo *ChoiceRepository) AddChoice(choice models.Choice) {
	repo.choiceCollection.InsertOne(context.TODO(), choice)
}

func (repo *ChoiceRepository) UpdateChoice(choice models.Choice) {
	filter := bson.M{"playerid": bson.M{"$eq": choice.PlayerId}}
	update := bson.M{"$set": bson.M{"roomid": choice.RoomId}}
	repo.choiceCollection.UpdateOne(context.TODO(), filter, update)
}
