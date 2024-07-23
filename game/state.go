package game

import (
	"fmt"
)

type Command string

const (
	SHOW_DECK                        Command = "show_deck"
	PLAY_STRAIGHT_TO_PACK            Command = "play_straight_to_pack"
	SWAP_DECK_CARD_WITH_CARD_IN_HAND Command = "swap_deck_card_with_card_in_hand"
	END_TURN                         Command = "end_turn"
)

type Player struct {
	ID    string
	Name  string
	Cards []Card
}

func NewPlayer(name string) Player {
	return Player{ID: GenerateID(10), Name: name, Cards: []Card{}}
}

type State struct {
	Players       []Player
	RoomID        string
	Deck          []Card
	Pack          []Card
	Turn          int
	Started       bool
	NumberPlayers int
}

func NewState(number_of_players int) State {
	return State{Players: []Player{}, RoomID: GenerateID(4), Deck: NewDeck(), Pack: []Card{}, Turn: 0, Started: false, NumberPlayers: number_of_players}
}

func (s *State) AddPlayer(p Player) {
	s.Players = append(s.Players, p)
}

func (s *State) DealCards() {
	for i := range s.Players {
		for range [4]int{} {
			card := s.Deck[len(s.Deck)-1]
			s.Deck = s.Deck[:len(s.Deck)-1]
			s.Players[i].Cards = append(s.Players[i].Cards, card)
		}
	}
}

func (s *State) MoveCardFromDeckToPack() {
	card := s.Deck[len(s.Deck)-1]
	s.Deck = s.Deck[:len(s.Deck)-1]
	s.Pack = append(s.Pack, card)
}

func (s *State) StartGame() {
	s.Started = true
	s.DealCards()
	s.MoveCardFromDeckToPack()
}

func (s *State) ShouldStart() bool {
	return s.NumberPlayers == len(s.Players)
}

func (s *State) IncrementTurn() {
	s.Turn += 1
}

func (s *State) GetPlayersTurn() *Player {
	index := s.Turn % s.NumberPlayers
	return &s.Players[index]
}

func (s *State) GetPlayersTurnString() string {
	return fmt.Sprintf("%s's turn!", s.GetPlayersTurn().Name)
}

func (s *State) IsTurn(player_id string) bool {
	players_turn := s.GetPlayersTurn()
	return players_turn.ID == player_id
}

func (s *State) ShowDeck() {
	card := &s.Deck[len(s.Deck)-1]
	card.FaceUp = true
}

func (s *State) PlayStraightToPack() error {
	card := s.Deck[len(s.Deck)-1]
	if !card.FaceUp {
		return CommandError{message: "expected card at top of deck to be face up"}
	}
	s.Pack = append(s.Pack, card)
	s.Deck = s.Deck[:len(s.Deck)-1]
	return nil
}

func (s *State) ProcessCommand(command Command) error {
	switch command {
	case END_TURN:
		s.IncrementTurn()
		return nil
	case SHOW_DECK:
		s.ShowDeck()
		return nil
	case PLAY_STRAIGHT_TO_PACK:
		err := s.PlayStraightToPack()
		return err
	case SWAP_DECK_CARD_WITH_CARD_IN_HAND:
		// s.SwapDeckCardWithCardInHand()
	}
	return nil
}

type CommandError struct {
	message string
}

func (c CommandError) Error() string {
	return c.message
}
