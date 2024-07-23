package game

import (
	"go-docer/game"
	"testing"
)

func createTestState(number_of_players int) game.State {
	state := game.NewState(number_of_players)
	player_1 := game.NewPlayer("Bob")
	player_2 := game.NewPlayer("James")
	player_3 := game.NewPlayer("Jimmy")
	player_4 := game.NewPlayer("Jeff")
	state.AddPlayer(player_1)
	state.AddPlayer(player_2)
	state.AddPlayer(player_3)
	state.AddPlayer(player_4)
	return state
}

func TestStartGame(t *testing.T) {
	number_of_players := 4
	state := createTestState(number_of_players)

	state.StartGame()
	if !state.Started {
		t.Errorf("state.Started is false, should be true.")
	}

	cards_in_deck := len(state.Deck)
	cards_in_pack := len(state.Pack)

	if cards_in_pack != 1 {
		t.Errorf("expected 1 card in pack, found %d", cards_in_pack)
	}

	expected_cards_in_deck := 52 - (number_of_players * 4) - 1
	if cards_in_deck != expected_cards_in_deck {
		t.Errorf("expected %d cards in deck, found %d", expected_cards_in_deck, cards_in_deck)
	}

	for i := range state.Players {
		player := &state.Players[i]
		if len(player.Cards) != 4 {
			t.Errorf("expected player to have 4 cards, got %d", len(player.Cards))
		}
	}
}

func TestProcessCommand(t *testing.T) {
	state := createTestState(4)
	if state.Turn != 0 {
		t.Errorf("expected turn to be 0, got %d", state.Turn)
	}

	state.ProcessCommand(game.END_TURN)
	if state.Turn != 1 {
		t.Errorf("expected turn to be 1, got %d", state.Turn)
	}

	state.ProcessCommand(game.SHOW_DECK)
	face_up_card := state.Deck[len(state.Deck)-1]
	if !face_up_card.FaceUp {
		t.Error("expected last card in deck to be face up")
	}

	state.ProcessCommand(game.PLAY_STRAIGHT_TO_PACK)
	pack_card := state.Pack[len(state.Pack)-1]
	if pack_card.Value != face_up_card.Value || pack_card.Suit != face_up_card.Suit {
		t.Errorf("expected card at top of pack to be %s, got %s", face_up_card.ToString(), pack_card.ToString())
	}

}
