package set

import (
	"cbtennis/internal/game"
	"cbtennis/internal/scoring"
	gamescore "cbtennis/internal/scoring/game"
	"cbtennis/internal/scoring/set"
	"cbtennis/internal/turning"
)

type OnStartingNewGame func()
type OnUpdatePoint func(valueA, valueB int)
type OnFinnishedSet func(valueA, valueB int)

type Set struct {
	gamesToWin  int
	confirmGame bool
	confirmSet  bool

	sideServing turning.Turning
	score       scoring.Scoring
	//	challege    player.Challenging
	games []game.GameManager

	startingNewGameEvent []OnStartingNewGame
	updatePointEvent     []OnUpdatePoint
	finnishedSetEvent    []OnFinnishedSet
}

func New(gamesToWin int, confirmGame bool, confirmSet bool) *Set {
	turn := turning.New(turning.TPOpposite) // Aqui eu deixo oposto para alinhar ao chamar NewGame
	sscc := set.NewSetScoreCountControl(gamesToWin, confirmSet)
	score := set.New(sscc)

	result := &Set{
		gamesToWin:           gamesToWin,
		confirmGame:          confirmGame,
		confirmSet:           confirmSet,
		sideServing:          turn,
		score:                score,
		games:                make([]game.GameManager, 0),
		startingNewGameEvent: make([]OnStartingNewGame, 0),
	}

	score.AddChangedScoreEvent(func(valueA, valueB string) {
		result.executeUpdatePointEvent()
	})

	score.AddReachedScoreEvent(func(valueA, valueB string) {
		result.executeFinnishedSetEvent()
	})

	return result
}

func (s *Set) NewGame() game.GameManager {
	s.sideServing.DoTurn()
	gscc := gamescore.NewGameScoreCountControl(4, s.confirmGame)
	standardgame := game.NewSingleStandardGame(gscc, nil)

	s.executeStartingNewGameEvent()
	return standardgame
}

func (s *Set) AddGame(game game.GameManager) {
	s.games = append(s.games, game)
	scc := s.score.GetScoreCountControl().(*set.SetScoreCountControl)
	scc.SetCurrentGame(game)
	s.score.UpdateScore()
}

func (s *Set) AddStartingNewGameEvent(startingNewGameEvent OnStartingNewGame) {
	s.startingNewGameEvent = append(s.startingNewGameEvent, startingNewGameEvent)
}

func (s *Set) AddUpdatePointEvent(updatePointEvent OnUpdatePoint) {
	s.updatePointEvent = append(s.updatePointEvent, updatePointEvent)
}

func (s *Set) AddFinishedSetEvent(finnishedSetEvent OnFinnishedSet) {
	s.finnishedSetEvent = append(s.finnishedSetEvent, finnishedSetEvent)
}

func (s *Set) executeStartingNewGameEvent() {
	for _, evt := range s.startingNewGameEvent {
		evt()
	}
}

func (s *Set) executeUpdatePointEvent() {
	for _, evt := range s.updatePointEvent {
		A, B := s.score.GetStatus()
		evt(A, B)
	}
}

func (s *Set) executeFinnishedSetEvent() {
	for _, evt := range s.finnishedSetEvent {
		A, B := s.score.GetStatus()
		evt(A, B)
	}
}
