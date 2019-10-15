package models

type Kitchen struct {}

func (r *Kitchen) Id() int {
	return 2
}

func (r *Kitchen) Name() string {
	return "Cuisine"
}

func (r *Kitchen) Description() string {
	return "Récupération de la cuisine + 1 pièce aléatoire dans le defausse"
}

func (r *Kitchen) PowerRoom(p Player) Player{
	if room := p.RandomGraveyardPlayer(); room != nil {
		p.RecoverRoomGraveyard(room)
	}
	return p
}

func (r *Kitchen) MalusRoom(p Player) Player {
	p.life -= 1
	p.RemoveRoomDeck(r)
	return p
}