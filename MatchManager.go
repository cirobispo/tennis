package cbtennis

type MatchManager struct {
	match *TennisMatch
	turn  *TurnManager
	score *GameScore

	sets       []SetManager
	currentSet SetManager
}

func NewMatchManager(match *TennisMatch) MatchManager {
	score := NewGameScore()
	turn := NewTurnManager(TPEven)

	return MatchManager{
		match: match,
		score: &score,
		turn:  &turn,
		sets:  make([]SetManager, 0),
	}
}

func (m *MatchManager) NewSet() SetManager {
	if m.currentSet != nil {
		m.sets = append(m.sets, m.currentSet)
	}

	new_set := NewCustomSet(m.match, m.match.gamesToWin, m.match.confirmSet, m.match.confirmGame)
	m.currentSet = &new_set

	return m.currentSet
}

func (m *MatchManager) GetCurrentSet() SetManager {
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
