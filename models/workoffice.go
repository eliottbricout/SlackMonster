package models

type WorkOffice struct {}

func (r *WorkOffice) Id() int {
	return 4
}

func (r *WorkOffice) Name() string {
	return "Bureau"
}

func (r *WorkOffice) Description() string {
	return "Vous pouvez récupérer les clés d'une nouvelle pièce de la maison"
}

func (r *WorkOffice) PowerRoom(p Player) Player{
	p.RemoveRoomDeck(r)
	return p
}

func (r *WorkOffice) MalusRoom(p Player) Player {
	p.life -= 1
	p.RemoveRoomDeck(r)
	return p
}