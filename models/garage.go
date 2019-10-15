package models

type Garage struct {}

func (r *Garage) Id() int {
	return 5
}

func (r *Garage) Name() string {
	return "Garage"
}

func (r *Garage) Description() string {
	return "Récupération du garage + 2 pièce aléatoire dans le defausse"
}

func (r *Garage) PowerRoom(p Player) Player{
	if room := p.RandomGraveyardPlayer(); room != nil {
		p.RecoverRoomGraveyard(room)
	}
	if room := p.RandomGraveyardPlayer(); room != nil {
		p.RecoverRoomGraveyard(room)
	}
	return p
}

func (r *Garage) MalusRoom(p Player) Player {
	p.life -= 1
	p.RemoveRoomDeck(r)
	return p
}