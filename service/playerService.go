package service

import (
	"../models"
	"../repository"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type PlayerService struct {
	repositoryPlayer repository.PlayerRepository
}

func CreatePlayerService(database mongo.Database) PlayerService {
	return PlayerService{repository.CreatePlayerRepository(database)}
}

func (service *PlayerService) GetPlayerRest(w http.ResponseWriter, r *http.Request) {
	player := service.GetPlayer("eliott")
	w.Write([]byte(player.Infos()))
}

func (service *PlayerService) GetPlayer(name string) models.Player {
	return service.repositoryPlayer.GetPlayer("eliott")
}
