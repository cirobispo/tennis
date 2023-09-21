package tennisstatus

type Player struct {
	name string
}

func NewPlayer(name string) Player {
	return Player{name: name}
}

func (p Player) TurnDescription() string {
	return p.name
}

type Team struct {
	players []*Player
	turn    TurnManager
}

func NewTeam(playersName []string) *Team {
	players := make([]*Player, 0)

	for _, p := range playersName {
		player := NewPlayer(p)
		players = append(players, &player)
	}

	turnM := NewTurnManager(TPEven)

	t := &Team{
		players: players,
		turn:    turnM,
	}

	return t
}
