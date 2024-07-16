package game

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
	for _, player := range s.Players {
		for range 4 {
			card := s.Deck[len(s.Deck)-1]
			s.Deck = s.Deck[:len(s.Deck)-1]
			player.Cards = append(player.Cards, card)
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

func (s *State) GetPlayer(player_id string) *Player {
	for _, player := range s.Players {
		if player.ID == player_id {
			return &player
		}
	}
	return nil
}

// add tests here
