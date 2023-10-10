package scoring

type Score struct {
	valueA            int
	valueB            int
	typeMode          ScoringType
	ctrl              ScoringCountControl
	reachedScoreEvent []OnScoringGame
	changedScoreEvent []OnScoringChange
}

func NewCustomGameScore(typeMode ScoringType, scc ScoringCountControl) *Score {
	return &Score{
		valueA:            0,
		valueB:            0,
		typeMode:          typeMode,
		ctrl:              scc,
		reachedScoreEvent: make([]OnScoringGame, 0),
		changedScoreEvent: make([]OnScoringChange, 0),
	}
}

func (s *Score) Reset() {
	s.valueA = 0
	s.valueB = 0
	s.executeChangedScoreEvent()
}

func (s Score) GetStatus() (int, int) {
	return s.valueA, s.valueB
}

func (s Score) GetScoringType() ScoringType {
	return s.typeMode
}

func (s *Score) UpdateScore() {
	s.ctrl.UpdateScore(s.ctrl, &s.valueA, &s.valueB)

	s.executeChangedScoreEvent()

	if s.ctrl.IsDone(s.valueA, s.valueB) {
		s.executeScoreGameEvent()
	}
}

func (s Score) GetScoreCountControl() ScoringCountControl {
	return s.ctrl
}

func (s *Score) AddReachedScoreEvent(scoreGameEvent OnScoringGame) {
	s.reachedScoreEvent = append(s.reachedScoreEvent, scoreGameEvent)
}

func (s *Score) AddChangedScoreEvent(changedScoreEvent OnScoringChange) {
	s.changedScoreEvent = append(s.changedScoreEvent, changedScoreEvent)
}

func (s Score) executeChangedScoreEvent() {
	for i := 0; i < len(s.changedScoreEvent); i++ {
		evt := s.changedScoreEvent[i]
		helper := NewScoreDataWrapper(s.typeMode, s.valueA, s.valueB)
		evt(helper.GetValueA(), helper.GetValueB())
	}
}

func (s Score) executeScoreGameEvent() {
	for i := 0; i < len(s.reachedScoreEvent); i++ {
		evt := s.reachedScoreEvent[i]
		helper := NewScoreDataWrapper(s.typeMode, s.valueA, s.valueB)
		evt(helper.GetValueA(), helper.GetValueB())
	}
}
