package models

type Room interface {
	Id() int
	Name() string
	Description() string
	PowerRoom(p Player) Player
	MalusRoom(p Player) Player
}

func GetAllRooms() []Room {
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

func GetInitRooms() []Room {
	return []Room{
		Room(new(BedRoom)),
		Room(new(Kitchen)),
		Room(new(LivingRoom)),
		Room(new(WorkOffice)),
	}
}

func GetRoom(id int) Room {
	for _, room := range GetAllRooms() {
		if room.Id() == id {
			return room
		}
	}
	return nil
}

func createRooms(idRooms []int) []Room {
	var rooms []Room
	for _, idRoom := range idRooms {
		rooms = addRoom(rooms, GetRoom(idRoom))
	}
	return rooms
}

func getIdRooms(rooms []Room) []int {
	var idRooms []int
	for _, room := range rooms {
		idRooms = append(idRooms, room.Id())
	}
	return idRooms
}
