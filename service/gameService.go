package service

import (
	"github.com/eliottbricout/SlackMonster/models"
	"github.com/eliottbricout/SlackMonster/repository"
	"fmt"
	"math/rand"
	"net/http"
)

type GameService struct {
	playerService   PlayerService
	choiceService   ChoiceService
	printRepository repository.PrintRepository
}

func CreateGameService(choicesService ChoiceService, playerService PlayerService, url string) GameService {
	return GameService{
		playerService,
		choicesService,
		repository.CreatePrintRepository(url),
	}
}

func (service *GameService) StartParty(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.Write([]byte("Erreur"))
		return
	}
	if r.PostForm.Get("user_id") == "UCPENAHK2" {
		w.Write([]byte("Seul l'élu peut lancer une party"))
		return
	}

	players := service.playerService.GetAllPlayer()
	indexMonster := rand.Intn(len(players))
	fmt.Print()
	for i, player := range players {
		player.InitParty()
		if indexMonster == i {
			player.SetMonster(true)
			service.printRepository.Print("*" + player.Name() + "* est le DarkMickey :darkmickey: !")
		}
		service.playerService.UpdatePlayer(player)
	}
}

func (service *GameService) NextDay(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.Write([]byte("Erreur"))
		return
	}
	if r.PostForm.Get("user_id") == "UCPENAHK2" {
		w.Write([]byte("Seul l'élu peut voyager dans le temps"))
		return
	}
	players := service.playerService.GetAllPlayer()
	players = allAlivePlayer(players)
	choiceRoomMonster, monster := service.choiceMonster(players)
	service.resolveDay(players, choiceRoomMonster, monster)
}

func (service *GameService) choiceMonster(players []models.Player) (int, models.Player) {
	for _, player := range players {
		if player.IsMonster() {
			if choice, err := service.choiceService.searchChoice(player); err == nil {
				return choice.RoomId, player
			} else {
				return randomChoice(), player
			}
		}
	}
	return 0, models.Player{}
}

func (service *GameService) resolveDay(players []models.Player, choiceRoomMonster int, monster models.Player) {
	service.printRepository.Print("\n 🌙 fin de la journée 🌙 \n\n")
	for _, player := range players {
		if !player.IsMonster() {
			choice, err := service.choiceService.searchChoice(player)
			if err != nil {
				service.activeRoom(player, player.RandomChoicePlayer(), choiceRoomMonster, monster)
			} else {
				service.activeRoom(player, models.GetRoom(choice.RoomId), choiceRoomMonster, monster)
			}
		}
	}
}

func (service *GameService) activeRoom(player models.Player, room models.Room, choiceRoomMonster int, monster models.Player) {
	if room.Id() == choiceRoomMonster {
		player := room.MalusRoom(player)
		service.printRepository.Print(":darkmickey: *" + monster.Name() + "* 🔪 *" + player.Name() + "* dans _" + room.Name() + "_")
		service.choiceService.playerService.UpdatePlayer(player)
	} else {
		player := room.PowerRoom(player)
		service.printRepository.Print("*" + player.Name() + "* entre dans _" + room.Name() + "_")
		service.choiceService.playerService.UpdatePlayer(player)
	}
}

func randomChoice() int {
	return rand.Intn(7) + 1
}


func allAlivePlayer(allplayer []models.Player) []models.Player {
	var players []models.Player
	for _, player := range allplayer {
		if player.Life() > 0 {
			players = append(players, player)
		}
	}
	return players
}