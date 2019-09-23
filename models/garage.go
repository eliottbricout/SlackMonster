package models

type Garage struct {}

func (r *Garage) Id() int {
	return 5
}

func (r *Garage) Name() string {
	return "Garage"
}

func (r *Garage) Description() string {
	return "Récupération du garage + 2 pièce dans le defausse (utiliser 2 fois la commande : /recoverRoom {idRoom})"
}

func (r *Garage) PowerRoom(p Player) Player{
	p.RemoveRoomDeck(r)
	return p
}

func (r *Garage) MalusRoom(p Player) Player {
	p.life -= 1
	p.RemoveRoomDeck(r)
	return p
}