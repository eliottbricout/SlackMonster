package service

import (
	"github.com/eliottbricout/SlackMonster/models"
	"fmt"
	"net/http"
)

func ListRoom(w http.ResponseWriter, _ *http.Request) {
	desc := ""
	for _ , room := range models.GetAllRooms() {
		desc += fmt.Sprintf("[%d] %s : %s\n", room.Id(), room.Name(), room.Description())
	}
	w.Write([]byte(desc))
}
