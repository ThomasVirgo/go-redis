package database

import (
	"encoding/json"
	"go-docer/game"
)

func GetState(room_id string) (game.State, error) {
	val, err := DB.Get(CTX, room_id).Result()
	if err != nil {
		return game.State{}, err
	}
	var state game.State
	if err := json.Unmarshal([]byte(val), &state); err != nil {
		return game.State{}, err
	}
	return state, nil
}

func SetState(state game.State) error {
	b, err := json.Marshal(state)
	if err != nil {
		return nil
	}
	err = DB.Set(CTX, state.RoomID, b, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
