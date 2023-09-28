package cbtennis

type TennisMatch struct {
	setsToVictory int
	gamesToWin    int
	confirmGame   bool
	confirmSet    bool
	teamA, teamB  *Team
	turn          *TurnManager
	score         *GameScore
}

func NewTennisMatch(teamA, teamB *Team, setsToVictory, gamesToWin int, confirmGame, confirmSet bool) TennisMatch {
	score := NewGameScore()
	turn := NewTurnManager(TPEven)

	return TennisMatch{
		teamA:         teamA,
		teamB:         teamB,
		score:         &score,
		turn:          &turn,
		setsToVictory: setsToVictory,
		gamesToWin:    gamesToWin,
		confirmGame:   confirmGame,
		confirmSet:    confirmSet,
	}
}

func (m TennisMatch) SetsToVictory() int {
	return m.setsToVictory
}

func (m TennisMatch) GamesToWinEachSet() int {
	return m.gamesToWin
}

func (m TennisMatch) GameConfirmation() bool {
	return m.confirmGame
}

func (m TennisMatch) SetConfirmation() bool {
	return m.confirmSet
}

func (m *TennisMatch) GetScoreManager() *GameScore {
	return m.score
}

func (m *TennisMatch) GetTurnManager() *TurnManager {
	return m.turn
}

func (m *TennisMatch) GetTeamA() *Team {
	return m.teamA
}

func (m *TennisMatch) GetTeamB() *Team {
	return m.teamB
}
