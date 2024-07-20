package main

import (
	"log"
	"net/http"

	"go-docer/database"
	"go-docer/handlers"
)

func main() {
	database.InitRedisClient()

	_, err := database.DB.Ping(database.CTX).Result()
	if err != nil {
		log.Fatal(err)
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /", handlers.Index)
	router.HandleFunc("GET /room/{room_id}/player/{player_id}/", handlers.GetGameState)
	router.HandleFunc("GET /room/{room_id}/player/{player_id}/state", handlers.GetPlayerGameState)
	router.HandleFunc("POST /room/{room_id}/player/{player_id}/take_turn", handlers.TakeTurn)
	router.HandleFunc("POST /new_room", handlers.NewRoom)
	router.HandleFunc("POST /join_room", handlers.JoinRoom)

	http.ListenAndServe(":9000", router)
}
