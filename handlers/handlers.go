package handlers

import (
	"encoding/json"
	"fmt"
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
	player_id := r.PathValue("player_id")
	fmt.Println(player_id)
	val, err := database.DB.Get(database.CTX, room_id).Result()
	if err != nil {
		w.Write([]byte("error reading from redis store"))
	}
	var state game.State
	if err := json.Unmarshal([]byte(val), &state); err != nil {
		w.Write([]byte("error converting json to state struct"))
	}
	b, err := json.Marshal(state)
	if err != nil {
		w.Write([]byte("error converting state to json"))
	}
	w.Write(b)
}

func NewRoom(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}
	name := r.Form.Get("name")
	state := game.NewState()
	state.AddPlayer(game.NewPlayer(name))
	b, err := json.Marshal(state)
	if err != nil {
		http.Error(w, "Failed to marshal state", http.StatusInternalServerError)
		return
	}
	err = database.DB.Set(database.CTX, state.RoomID, b, 0).Err()
	if err != nil {
		http.Error(w, "Failed to write to redis db", http.StatusBadRequest)
		return
	}
	redirect_path := fmt.Sprintf("/room/%s/player/%s", state.RoomID, state.Players[0].ID)
	w.Header().Set("HX-Redirect", redirect_path)
}
