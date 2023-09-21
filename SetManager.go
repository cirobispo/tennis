package tennisstatus

type SetManager interface {
	NewGame() GameManager
	GetCurrentGame() GameManager
}

type StandardSet struct {
	match       *TennisMatch
	gamesToWin  int
	confirmSet  bool
	confirmGame bool
	games       []GameManager
	currentGame GameManager
}

func NewCustomSet(match *TennisMatch, gamesToWin int, confirmSet, confirmGame bool) StandardSet {
	return StandardSet{
		match:       match,
		gamesToWin:  gamesToWin,
		confirmSet:  confirmSet,
		confirmGame: confirmGame,
		games:       make([]GameManager, 0),
	}
}

func NewStandardSet(match *TennisMatch) StandardSet {
	return NewCustomSet(match, 6, true, true)
}

func (s *StandardSet) NewGame() GameManager {
	return s.startNewGame()
}

func (s *StandardSet) startNewGame() GameManager {
	if s.currentGame != nil {
		s.games = append(s.games, s.currentGame)
	}

	new_game := NewSingleStandardGame(s.match)
	new_game.AddGameFinishEvent(func() {
		s.startNewGame()
	})
	s.currentGame = &new_game

	return s.currentGame
}

func (s *StandardSet) GetCurrentGame() GameManager {
	return s.currentGame
}
