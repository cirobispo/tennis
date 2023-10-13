package set

import (
	"cbtennis/internal/game"
	"cbtennis/internal/scoring"
	"cbtennis/internal/scoring/gamescore"
	"cbtennis/internal/scoring/set"
	"cbtennis/internal/turning"
)

type Set struct {
	gamesToWin            int
	confirmGame           bool
	confirmSet            bool
	confirmNthGameWithTie bool
	sideServing           turning.Turning
	score                 scoring.Scoring
	started               bool
	//	challege    player.Challenging
	games []game.GameManager

	startingNewGameEvent []OnStartingNewGame
	updatePointEvent     []OnUpdatePoint
	startedSetEvent      []OnStartedSet
	finnishedSetEvent    []OnFinnishedSet
}

func New(beginningTurn turning.TurnPosition, gamesToWin int, confirmGame bool, confirmSet bool, confirmNthGameWithTie bool) *Set {
	turn := turning.New(beginningTurn)
	sscc := set.NewSetScoreCountControl(gamesToWin, confirmSet)
	score := set.New(sscc)

	result := &Set{
		gamesToWin:            gamesToWin,
		confirmGame:           confirmGame,
		confirmSet:            confirmSet,
		confirmNthGameWithTie: confirmNthGameWithTie,
		sideServing:           turn,
		score:                 score,
		started:               false,
		games:                 make([]game.GameManager, 0),
		startingNewGameEvent:  make([]OnStartingNewGame, 0),
	}

	score.AddChangedScoreEvent(func(valueA, valueB string) {
		result.executeUpdatePointEvent()
	})

	score.AddReachedScoreEvent(func(valueA, valueB string) {
		result.executeFinnishedSetEvent()
	})

	return result
}

func (s *Set) StartSet() {
	s.started = true
	s.executeStartedSetEvent()
}

func (s *Set) NewGame() game.GameManager {
	if s.started {
		s.sideServing.DoTurn()
		sscc := s.score.GetScoreCountControl()
		if !sscc.IsTie() || !sscc.HasToConfirm() {
			gscc := gamescore.NewGameScoreCountControl(4, s.confirmGame)
			standardgame := game.NewSingleStandardGame(gscc, nil, s.sideServing.CurrentTurn())

			s.executeStartingNewGameEvent()
			return standardgame
		} else {
			return nil
		}
	}

	return nil
}

func (s *Set) GetScore() scoring.Scoring {
	return s.score
}

func (s *Set) AddGame(game game.GameManager) {
	if s.started {
		s.sideServing.DoTurn()
		s.games = append(s.games, game)
		scc := s.score.GetScoreCountControl().(*set.SetScoreCountControl)
		scc.SetCurrentGame(game)
		s.score.UpdateScore()
	}
}

func (s *Set) AddStartingNewGameEvent(startingNewGameEvent OnStartingNewGame) {
	s.startingNewGameEvent = append(s.startingNewGameEvent, startingNewGameEvent)
}

func (s *Set) AddUpdatePointEvent(updatePointEvent OnUpdatePoint) {
	s.updatePointEvent = append(s.updatePointEvent, updatePointEvent)
}

func (s *Set) AddStartedSetEvent(startedSetEvent OnStartedSet) {
	s.startedSetEvent = append(s.startedSetEvent, startedSetEvent)
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

	s.started = false
}

func (s *Set) executeStartedSetEvent() {
	for _, evt := range s.startedSetEvent {
		evt()
	}
}
