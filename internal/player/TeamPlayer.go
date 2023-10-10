package player

import "cbtennis/internal/turning"

type Team struct {
	players []Playing
	turn    *turning.Turn
}

func NewTeam(playersName []string) *Team {
	players := make([]Playing, 0)

	for _, p := range playersName {
		player := NewPlayer(p)
		players = append(players, player)
	}

	turnM := turning.New(turning.TPBegin)

	t := &Team{
		players: players,
		turn:    turnM,
	}

	return t
}

func (t Team) GetTeamName() (result string) {
	size := len(t.players)
	for i, p := range t.players {
		result += p.GetName()
		if i < size-1 {
			result += "/"
		}
	}

	return result
}

func (t Team) GetName() (result string) {
	currentPlayer := t.players[0].GetName()
	if t.turn.CurrentTurn() != t.turn.BeginningTurn() {
		currentPlayer = t.players[1].GetName()
	}

	result = t.GetTeamName() + ": " + currentPlayer
	return result
}

func (t *Team) DoTurn() {
	t.turn.DoTurn()
}
