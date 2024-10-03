package main

import (
	"log"
	"net/http"

	"go-docer/database"
	"go-docer/handlers"
	"go-docer/socket"
)

func main() {
	database.InitRedisClient()

	_, err := database.DB.Ping(database.CTX).Result()
	if err != nil {
		log.Fatal(err)
	}
	hub := socket.NewHub()
	go hub.Run()
	router := http.NewServeMux()
	// landing page
	router.HandleFunc("GET /", handlers.Index)
	router.HandleFunc("POST /new_room", handlers.NewRoom)
	router.HandleFunc("POST /join_room", handlers.JoinRoom)

	// game page
	router.HandleFunc("GET /room/{room_id}/player/{player_id}/", handlers.GetGameState)
	router.HandleFunc("GET /ws", func(w http.ResponseWriter, r *http.Request) {
		socket.ServeWebSocket(hub, w, r)
	})

	http.ListenAndServe(":9000", router)
}
