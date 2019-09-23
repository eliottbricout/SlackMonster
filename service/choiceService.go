package service

import (
	"../repository"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type ChoiceService struct {
	repositoryChoice repository.ChoiceRepository
	playerService PlayerService
}

func CreateChoiceService(database mongo.Database, service PlayerService) ChoiceService{
	return ChoiceService {repository.CreateChoiceRepository(database), service}
}

func(service *ChoiceService) ChoiceRest(w http.ResponseWriter, r *http.Request) {
	player := service.playerService.GetPlayer("eliott")

	choice, err := player.Choice(0)
	service.repositoryChoice.AddChoice(choice)
	if err {
		w.Write([]byte("la pièce n'est pas dans votre deck"))
	} else {
		w.Write([]byte(fmt.Sprintf("votre choix a été sauvegardé")))
	}
}
