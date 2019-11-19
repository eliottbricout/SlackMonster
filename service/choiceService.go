package service

import (
	"github.com/eliottbricout/SlackMonster/models"
	"github.com/eliottbricout/SlackMonster/repository"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strconv"
)

type ChoiceService struct {
	repositoryChoice repository.ChoiceRepository
	playerService    PlayerService
}

func CreateChoiceService(database mongo.Database, service PlayerService) ChoiceService {
	return ChoiceService{repository.CreateChoiceRepository(database), service}
}

func (service *ChoiceService) ChoiceRest(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.Write([]byte("Erreur"))
		return
	}
	userId := r.PostForm.Get("user_id")
	idRoom := r.PostForm.Get("text")
	player, err := service.playerService.getPlayer(userId)
	if err != nil {
		w.Write([]byte("Vous n'êtes pas dans la partie"))
	} else if id, err := strconv.Atoi(idRoom); err == nil {
		service.addChoice(w, player, id)
	} else {
		w.Write([]byte("c'est pas un nombre ça"))
	}

}

func (service *ChoiceService) addChoice(w http.ResponseWriter, player models.Player, id int) {
	choice, noRoom := player.Choice(id)
	if noRoom {
		w.Write([]byte("la pièce n'est pas dans votre deck"))
	} else {
		service.addUpdateChoice(w, player, choice)
	}
}

func (service *ChoiceService) addUpdateChoice(w http.ResponseWriter, player models.Player, choice models.Choice) {
	if _, err := service.repositoryChoice.GetChoice(player.Id()); err == nil {
		service.repositoryChoice.UpdateChoice(choice)
		w.Write([]byte(fmt.Sprintf("votre choix a été sauvegardé")))
	} else {
		service.repositoryChoice.AddChoice(choice)
		w.Write([]byte(fmt.Sprintf("votre choix a été sauvegardé")))
	}
}

func (service *ChoiceService) searchChoice(player models.Player) (models.Choice, error) {
	return service.repositoryChoice.GetChoice(player.Id())
}
