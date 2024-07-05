package handlers

import (
	"encoding/json"
	"go-docer/database"
	"go-docer/game"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
		return
	}
}

func GetRoomState(w http.ResponseWriter, r *http.Request) {
	room_id := r.PathValue("room_id")
	val, err := database.DB.Get(database.CTX, room_id).Result()
	if err != nil {
		w.Write([]byte("error reading from redis store"))
	}
	var state game.State
	if err := json.Unmarshal([]byte(val), &state); err != nil {
		w.Write([]byte("error converting json to state struct"))
	}
	state.AddPlayer(game.NewPlayer("Bob"))
	b, err := json.Marshal(state)
	if err != nil {
		w.Write([]byte("error converting state to json"))
	}
	w.Write(b)
}

func NewRoom(w http.ResponseWriter, r *http.Request) {
	state := game.NewState()
	state.AddPlayer(game.NewPlayer("Tom"))
	b, err := json.Marshal(state)
	if err != nil {
		w.Write([]byte("error converting state to json"))
	}
	err = database.DB.Set(database.CTX, state.RoomID, b, 0).Err()
	if err != nil {
		w.Write([]byte("error setting data in redis"))
	}
	w.Write(b)
}
