package game

type Player struct {
	ID   string
	Name string
}

func NewPlayer(name string) Player {
	return Player{ID: GenerateID(10), Name: name}
}

type State struct {
	Players []Player
	RoomID  string
}

func NewState() State {
	return State{Players: []Player{}, RoomID: GenerateID(6)}
}

func (s *State) AddPlayer(p Player) {
	s.Players = append(s.Players, p)
}
