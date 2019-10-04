package service

import (
	"../models"
	"../repository"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type PlayerService struct {
	repositoryPlayer repository.PlayerRepository
}

func CreatePlayerService(database mongo.Database) PlayerService {
	return PlayerService{repository.CreatePlayerRepository(database)}
}

func (service *PlayerService) InfoPlayer(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.Write([]byte("Erreur"))
		return
	}
	userId := r.PostForm.Get("user_id")
	fmt.Print(userId)
	player, err := service.getPlayer(userId)
	if err != nil {
		w.Write([]byte("Joueur introuvable"))
	} else {
		w.Write([]byte(player.Infos()))
	}
}

func (service *PlayerService) JoinPlayer(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.Write([]byte("Erreur"))
		return
	}
	userId := r.PostForm.Get("user_id")
	username := r.PostForm.Get("user_name")

	if _ , err := service.getPlayer(userId); err == nil {
		w.Write([]byte("Vous êtes déjà dans la partie"))
	} else {
		service.addPlayer(username, userId)
		w.Write([]byte("vous avez join la partie"))
	}
}

func (service *PlayerService) CheckPlayer(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.Write([]byte("Erreur"))
		return
	}
	username := r.PostForm.Get("text")
	if username == "all" {
		service.checkAllPlayer(w)
	} else {
		service.checkUsernamePlayer(w, username)
	}
}

func (service *PlayerService) addPlayer(name string, id string) {
	player := models.CreatePlayer(name, id)
	service.repositoryPlayer.AddPlayer(player.TransformPlayerPivot())
}

func (service *PlayerService) getPlayer(id string) (models.Player, error){
	return service.repositoryPlayer.GetPlayerById(id)
}

func (service *PlayerService) checkAllPlayer(w http.ResponseWriter){
	players := service.repositoryPlayer.GetAllPlayer()
	for _, player := range players {
		w.Write([]byte(player.Infos()))
	}
}

func (service *PlayerService) checkUsernamePlayer(w http.ResponseWriter, username string) {
	player , err := service.repositoryPlayer.GetPlayer(username)
	if err != nil {
		w.Write([]byte("Le joueur " + username + " n'existe pas"))
	} else {
		w.Write([]byte(player.Infos()))
	}
}