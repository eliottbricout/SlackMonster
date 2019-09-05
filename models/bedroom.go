package models

type BedRoom struct {}

func (r *BedRoom) Id() int {
	return 1
}

func (r *BedRoom) Name() string {
	return "Chambre du dark mickey"
}

func (r *BedRoom) Description() string {
	return "Récupération de toutes pièces mais attention si le dark mickey vous surprend dans ça chambre il vous tue direct"
}

func (r *BedRoom) PowerRoom(p Player) Player{
	for _, room := range p.Graveyard() {
		p.RecoverRoomGraveyard(room)
	}
	p.RemoveRoomDeck(r)
	return p
}

func (r *BedRoom) MalusRoom(p Player) Player {
	p.life -= 2
	p.RemoveRoomDeck(r)
	return p
}