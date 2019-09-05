package models

type Toilet struct {}

func (r *Toilet) Id() int {
	return 7
}

func (r *Toilet) Name() string {
	return "Toilettes"
}

func (r *Toilet) Description() string {
	return "Vous avez un petit moment de répit le dark mickey ne risque pas de vous déranger ici"
}

func (r *Toilet) PowerRoom(p Player) Player{
	p.RemoveRoomDeck(r)
	return p
}

func (r *Toilet) MalusRoom(p Player) Player {
	p.RemoveRoomDeck(r)
	return p
}