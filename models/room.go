package models

type Room interface {
    Id() int
    Name() string
    Description() string
    PowerRoom(p Player) Player
    MalusRoom(p Player) Player
}



func getAllRooms() []Room {
	return []Room{
		Room(new(BedRoom)), 
		Room(new(Kitchen)),
		Room(new(LivingRoom)), 
		Room(new(WorkOffice)),
		Room(new(Garage)),
		Room(new(BathRoom)), 
		Room(new(Toilet)),
	}
}