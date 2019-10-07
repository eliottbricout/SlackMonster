package service

import (
	"../models"
	"math/rand"
	"net/http"
)

type DayService struct {
	playerService PlayerService
	choiceService ChoiceService
}

func CreateDayService(choicesService ChoiceService, playerService PlayerService) DayService {
	return DayService{playerService, choicesService}
}

func (service *DayService) NextDay(w http.ResponseWriter, r *http.Request) {
	players := service.playerService.GetAllPlayer()
	choiceRoomMonster, monster := service.choiceMonster(players)
	service.resolveDay(players, choiceRoomMonster)
}

func (service *DayService) choiceMonster(players []models.Player) (int, models.Player) {
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

func (service *DayService) resolveDay(players []models.Player, choiceRoomMonster int) {
	for _, player := range players {
		if !player.IsMonster() {
			choice, err := service.choiceService.searchChoice(player)
			if err != nil {
				activeRoom(player,  randomChoicePlayer(player), choiceRoomMonster)
			} else {
				activeRoom(player,  choice.RoomId, choiceRoomMonster)
			}
		}
	}
}

func activeRoom(player models.Player, roomId int, choiceRoomMonster int) {
	if roomId == choiceRoomMonster {
		models.GetRoom(roomId).MalusRoom(player)
	} else {
		models.GetRoom(roomId).PowerRoom(player)
	}
}

func randomChoice() int {
	return rand.Intn(7) + 1
}

func randomChoicePlayer(player models.Player) int {
	randRoom := rand.Intn(len(player.Deck()))
	for i, room := range player.Deck() {
		if i == randRoom {
			return room.Id()
		}
	}
}
