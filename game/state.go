package game

import "fmt"

type Command string

const (
	SHOW_DECK             Command = "show_deck"
	PLAY_STRAIGHT_TO_PACK Command = "play_straight_to_pack"
	TAKE_FROM_PACK        Command = "take_from_pack"
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

func (s *State) StartGame() {
	s.Started = true
	s.DealCards()
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

// func (s *State) TakeTurn(command Command) {
// 	player := s.GetPlayersTurn()

// }

func (s *State) GetPlayersTurnString() string {
	return fmt.Sprintf("%s's turn!", s.GetPlayersTurn().Name)
}
