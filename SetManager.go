package tennisstatus

type SetManager interface {
	StartSet()
	GetCurrentGame() GameManager
}

type StandardSet struct {
	match       *MatchManager
	games       []GameManager
	currentGame GameManager
}

func NewStandardSet(match *MatchManager) StandardSet {
	return StandardSet{
		match: match,
		games: make([]GameManager, 0),
	}
}

func (s *StandardSet) StartSet() {
	s.startNewGame()
}

func (s *StandardSet) startNewGame() {
	if s.currentGame != nil {
		s.games = append(s.games, s.currentGame)
	}

	new_game := NewStandardGame(s.match)
	new_game.AddGameFinishEvent(func() {
		s.startNewGame()
	})
	s.currentGame = &new_game
	s.currentGame.StartGame()
}

func (s *StandardSet) GetCurrentGame() GameManager {
	return s.currentGame
}
