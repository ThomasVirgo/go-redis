package handlers

import (
	"fmt"
	"go-docer/database"
	"go-docer/game"
	"html/template"
	"net/http"
	"strconv"
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

func GetGameState(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/game.html")
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}
	room_id := r.PathValue("room_id")
	player_id := r.PathValue("player_id")
	fmt.Println(player_id)
	state, err := database.GetState(room_id)
	if err != nil {
		http.Error(w, "Failed to read state from DB", http.StatusInternalServerError)
		return
	}
	fmt.Println(state.RoomID)
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
		return
	}
}

func NewRoom(w http.ResponseWriter, r *http.Request) {
	// Get Form Data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}
	name := r.Form.Get("name")
	number_of_players := r.Form.Get("number_of_players")
	n, err := strconv.Atoi(number_of_players)
	if err != nil {
		http.Error(w, "number of players must be an integer", http.StatusBadRequest)
		return
	}

	// Create State
	state := game.NewState(n)
	state.AddPlayer(game.NewPlayer(name))

	// Write to DB
	err = database.SetState(state)
	if err != nil {
		http.Error(w, "Failed to write to database", http.StatusBadRequest)
		return
	}

	// Redirect
	redirect_path := fmt.Sprintf("/room/%s/player/%s", state.RoomID, state.Players[0].ID)
	w.Header().Set("HX-Redirect", redirect_path)
}

func JoinRoom(w http.ResponseWriter, r *http.Request) {
	// Get form Data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}
	name := r.Form.Get("name")
	room_id := r.Form.Get("room_id")

	// Read from DB
	state, err := database.GetState(room_id)
	if err != nil {
		http.Error(w, "Failed to read from database", http.StatusBadRequest)
		return
	}

	// Add player
	new_player := game.NewPlayer(name)
	state.AddPlayer(new_player)
	if state.ShouldStart() {
		state.StartGame()
	}

	// Update State in DB
	err = database.SetState(state)
	if err != nil {
		http.Error(w, "Failed to write to database", http.StatusBadRequest)
		return
	}

	redirect_path := fmt.Sprintf("/room/%s/player/%s", state.RoomID, new_player.ID)
	w.Header().Set("HX-Redirect", redirect_path)
}
