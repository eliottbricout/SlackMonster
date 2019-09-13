package models

import "fmt"

type Player struct {
	life int
	name string
	deck []Room
	graveyard []Room
}

type PlayerPivot struct {
	life uint32
	name string
	deck []uint32
	graveyard []uint32
}

func CreatePlayer(name string) Player{
	return Player{life: 2, name: name, deck: getAllRooms(), graveyard: []Room{}}
}

func (p *Player) Graveyard() []Room {
	return p.graveyard
}

func (p *Player) Deck() []Room {
	return p.deck
}

func (p *Player) Infos() string {
	return fmt.Sprintf("Vie %d\nJeu: %s\nCimetière: %s", p.life, displayDeck(p.deck), displayDeck(p.graveyard))
}


func (p *Player) RecoverRoomGraveyard(room Room) {
	p.deck = addRoom(p.deck, room)
	p.graveyard = removeRoom(p.graveyard, room)
}

func (p *Player) RemoveRoomDeck(room Room) {
	p.deck = removeRoom(p.deck, room)
	p.graveyard = addRoom(p.graveyard, room)
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
