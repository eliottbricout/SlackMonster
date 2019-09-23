package service

import (
	"../models"
	"../repository"
	"bytes"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"net/http"
)

type PlayerService struct {
	repositoryPlayer repository.PlayerRepository
}

func CreatePlayerService(database mongo.Database) PlayerService {
	return PlayerService{repository.CreatePlayerRepository(database)}
}

type SlackBody struct {
	UserId string `json:"user_id"`
	Text string `json:"text"`
}

func (service *PlayerService) PlayerRest(w http.ResponseWriter, r *http.Request) {
	var slackBody SlackBody
	json.NewDecoder(r.Body).Decode(&slackBody)
	body, _ := ioutil.ReadAll(r.Body)
	rdr1 := ioutil.NopCloser(bytes.NewBuffer(body))
	fmt.Println(r.PostFormValue("text"))
	fmt.Println(rdr1)
	fmt.Println(r.Header.Get("Content-type"))
	fmt.Println(r.Header.Get("user_id"))

	player := service.GetPlayer("eliott")
	w.Write([]byte(player.Infos()))
}

func (service *PlayerService) GetPlayer(name string) models.Player {
	return service.repositoryPlayer.GetPlayer("eliott")
}

func (service *PlayerService) addPlayer(name string) {
	player := models.CreatePlayer(name)
	service.repositoryPlayer.AddPlayer(player.TransformPlayerPivot())
}