package models

type BathRoom struct {}

func (r *BathRoom) Id() int {
	return 6
}

func (r *BathRoom) Name() string {
	return "Salle de bain"
}

func (r *BathRoom) Description() string {
	return "Utiliser la trousse de soin dans la salle de bain pour vous soignez de la dernière agression du dark mickey"
}

func (r *BathRoom) PowerRoom(p Player) Player{
	if p.life < 2 {
		p.life += 1
	}
	p.RemoveRoomDeck(r)
	return p
}

func (r *BathRoom) MalusRoom(p Player) Player {
	p.life -= 1
	p.RemoveRoomDeck(r)
	return p
}