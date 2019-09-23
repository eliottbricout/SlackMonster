package service

import (
	"fmt"
	"net/http"
	"../models"
)

func ListRoomRest(w http.ResponseWriter, r *http.Request) {
	desc := ""
	for _ , room := range models.GetAllRooms() {
		desc += fmt.Sprintf("[%d] %s : %s\n", room.Id(), room.Name(), room.Description())
	}
	w.Write([]byte(desc))
}
