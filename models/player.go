package models

import (
	"fmt"
	"math/rand"
)

type Player struct {
	id        string
	isMonster bool
	name      string
	life      int
	deck      []Room
	graveyard []Room
}

type PlayerPivot struct {
	Id        string
	IsMonster bool
	Name      string
	Life      int
	Deck      []int
	Graveyard []int
}

func CreatePlayer(name string, id string) Player {
	return Player{id: id, isMonster: false, life: 0, name: name, deck: []Room{}, graveyard: []Room{}}
}

func (p *Player) InitParty() {
	p.life = 2
	p.isMonster = false
	p.deck = GetInitRooms()
	p.graveyard = []Room{}
}

func (p *Player) Id() string {
	return p.id
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Life() int {
	return p.life
}

func (p *Player) Graveyard() []Room {
	return p.graveyard
}

func (p *Player) Deck() []Room {
	return p.deck
}

func (p *Player) IsMonster() bool {
	return p.isMonster
}

func (p *Player) SetMonster(isMonster bool) {
	p.isMonster = isMonster
}

func (p *Player) Infos() string {
	if p.isMonster {
		return ":darkmickey: *" + p.name + "* est le dark Mickey\n\n"
	}
	if p.life == 0 {
		return "*" + p.name +"* : ☠\n"
	}
	return fmt.Sprintf("Nom *%s*\n %d ❤ ️ \nJeu: %s\nDefausse: %s\n\n",
		p.name, p.life, displayDeck(p.deck), displayDeck(p.graveyard))
}

func (p *Player) RecoverRoomGraveyard(room Room) {
	p.deck = addRoom(p.deck, room)
	p.graveyard = removeRoom(p.graveyard, room)
}

func (p *Player) RemoveRoomDeck(room Room) {
	p.deck = removeRoom(p.deck, room)
	p.graveyard = addRoom(p.graveyard, room)
}

func (p *Player) Choice(idRoom int) (Choice, bool) {
	room := GetRoom(idRoom)
	return Choice{p.id, idRoom}, room == nil || !isPresent(p.deck, room)
}

func (p *PlayerPivot) TransformPlayer() Player {
	var player Player
	player.id = p.Id
	player.isMonster = p.IsMonster
	player.name = p.Name
	player.life = p.Life
	player.deck = createRooms(p.Deck)
	player.graveyard = createRooms(p.Graveyard)
	return player
}

func (p *Player) TransformPlayerPivot() PlayerPivot {
	var player PlayerPivot
	player.Id = p.id
	player.IsMonster = p.isMonster
	player.Name = p.name
	player.Life = p.life
	player.Deck = getIdRooms(p.deck)
	player.Graveyard = getIdRooms(p.graveyard)
	return player
}

func addRoom(rooms []Room, room Room) []Room {
	return append(rooms, room)
}

func removeRoom(rooms []Room, room Room) []Room {
	var slice []Room
	for _, r := range rooms {
		if r.Id() != room.Id() {
			slice = append(slice, r)
		}
	}
	return slice
}

func isPresent(rooms []Room, room Room) bool {
	for _, r := range rooms {
		if r.Id() == room.Id() {
			return true
		}
	}
	return false
}

func displayDeck(rooms []Room) string {
	display := ""
	for _, room := range rooms {
		display += fmt.Sprintf(", %d:%s", room.Id(), room.Name())
	}
	if len(display) > 1 {
		display = display[2:]
	}
	return "[" + display + "]"
}

func (p *Player) RandomChoicePlayer() Room {
	return randomRoom(p.deck)
}

func (p *Player) RandomGraveyardPlayer() Room {
	return randomRoom(p.graveyard)
}

func randomRoom(rooms []Room) Room {
	if len(rooms) == 0 {
		return nil
	}
	index := rand.Intn(len(rooms))
	return rooms[index]
}
