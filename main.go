package main

import (
		"fmt"
        "net/http"

        "./models"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"message": "hello world"}`))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8001", nil)
}

func test(){
	player := models.CreatePlayer("eliott")
    fmt.Println(player.Infos())
    player = player.Deck()[0].PowerRoom(player)
    fmt.Println(player.Infos())
    player = player.Deck()[0].PowerRoom(player)
    fmt.Println(player.Infos())
    player = player.Deck()[0].MalusRoom(player)
    fmt.Println(player.Infos())
    player = player.Deck()[0].PowerRoom(player)
    fmt.Println(player.Infos())
}