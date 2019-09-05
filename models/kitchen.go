package models

type Kitchen struct {}

func (r *Kitchen) Id() int {
	return 2
}

func (r *Kitchen) Name() string {
	return "Cuisine"
}

func (r *Kitchen) Description() string {
	return "Récupération de la cuisine + 1 pièce dans le cimetière (commande : /recoverRoom {idRoom})"
}

func (r *Kitchen) PowerRoom(p Player) Player{
	return p
}

func (r *Kitchen) MalusRoom(p Player) Player {
	p.life -= 1
	p.RemoveRoomDeck(r)
	return p
}