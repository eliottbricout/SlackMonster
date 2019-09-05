package models

type LivingRoom struct {}

func (r *LivingRoom) Id() int {
	return 3
}

func (r *LivingRoom) Name() string {
	return "Salon"
}

func (r *LivingRoom) Description() string {
	return ""
}

func (r *LivingRoom) PowerRoom(p Player) Player{
	
	p.RemoveRoomDeck(r)
	return p
}

func (r *LivingRoom) MalusRoom(p Player) Player {
	p.life -= 1
	p.RemoveRoomDeck(r)
	return p
}