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

func getRoom(id int) Room {
	for _, room := range getAllRooms() {
		if room.Id() == id {
			return room
		}
	}
	return nil
}

func createRooms(idRooms []int) []Room{
	var rooms []Room
	for _, idRoom := range idRooms {
		rooms = addRoom(rooms, getRoom(idRoom))
	}
	return rooms
}