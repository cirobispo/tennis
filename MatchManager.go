package tennisstatus

type MatchManager struct {
	match *TennisMatch
	turn  *TurnManager
	score *GameScore

	sets       []*StandardSet
	currentSet *StandardSet
}

func NewMatchManager(match *TennisMatch) MatchManager {
	score := NewGameScore()
	turn := NewTurnManager()

	return MatchManager{
		match: match,
		score: &score,
		turn:  &turn,
		sets:  make([]*StandardSet, 0),
	}
}

func (m *MatchManager) StartMatch() {
	m.turn.StartTurn()
	m.startNewSet()
}

func (m *MatchManager) startNewSet() {
	if m.currentSet != nil {
		m.sets = append(m.sets, m.currentSet)
	}

	new_set := NewStandardSet(m)
	m.currentSet = &new_set
	m.currentSet.StartSet()
}

func (m *MatchManager) GetCurrentSet() *StandardSet {
	return m.currentSet
}

func (m *MatchManager) GetMatch() *TennisMatch {
	return m.match
}

func (m *MatchManager) GetScoreManager() *GameScore {
	return m.score
}

func (m *MatchManager) GetTurnManager() *TurnManager {
	return m.turn
}
