package models

import "fmt"

type Player struct {
	life int
	name string
	deck []Room
	graveyard []Room
}

type PlayerPivot struct {
	Life int
	Name string
	Deck []int
	Graveyard []int
}

func CreatePlayer(name string) Player{
	return Player{life: 2, name: name, deck: GetAllRooms(), graveyard: []Room{}}
}

func (p *Player) Graveyard() []Room {
	return p.graveyard
}

func (p *Player) Deck() []Room {
	return p.deck
}

func (p *Player) Infos() string {
	return fmt.Sprintf("Nom %s\nVie %d\nJeu: %s\nDefausse: %s",
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
	room := getRoom(idRoom)
	return Choice{p.name, idRoom} , room == nil || !isPresent(p.deck, room)
}

func (p *PlayerPivot) TransformPlayer() Player{
	var player Player
	player.name = p.Name
	player.life = p.Life
	player.deck = createRooms(p.Deck)
	player.graveyard = createRooms(p.Graveyard)
	return player
}

func (p *Player) TransformPlayerPivot() PlayerPivot{
	var player PlayerPivot
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

func displayDeck(rooms []Room) string{
	display := ""
	for _, room := range rooms {
		display += fmt.Sprintf(", %d:%s", room.Id(), room.Name())
	}
	if len(display) > 1 {
		display = display[2:]
	}
	return "[" + display + "]"
}
