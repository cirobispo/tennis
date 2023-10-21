package scoring

type ScoringCountControl interface {
	MaxValue() int
	HasToConfirm() bool
	PlugToScoring(s *Score)
	IsDone(valueA, valueB int) bool
	IsTie() bool
	UpdateScore(self ScoringCountControl, valueA, valueB *int)
}

type UpdatingScoreHandler func(scc ScoringCountControl, valueA, valueB *int)
type ScoreIsDoneHandler func(maxValue int, hasToConfirm bool, valueA, valueB int) bool

type ScoreCountControl struct {
	maxValue     int
	hasToConfirm bool

	score *Score

	UpdateHandler UpdatingScoreHandler
	IsDoneHandler ScoreIsDoneHandler
}

func NewScoreCountControl(maxValue int, confirm bool, update UpdatingScoreHandler, done ScoreIsDoneHandler) *ScoreCountControl {
	return &ScoreCountControl{
		maxValue:      maxValue,
		hasToConfirm:  confirm,
		UpdateHandler: update,
		IsDoneHandler: done,
	}
}

func (c *ScoreCountControl) PlugToScoring(s *Score) {
	c.score = s
}

func (c ScoreCountControl) MaxValue() int {
	return c.maxValue
}

func (c ScoreCountControl) HasToConfirm() bool {
	return c.hasToConfirm
}

func (c ScoreCountControl) IsTie() bool {
	return (c.score.valueA == c.maxValue) && (c.score.valueB == c.maxValue)
}

func (c ScoreCountControl) IsDone(valueA, valueB int) bool {
	if c.UpdateHandler != nil {
		return c.IsDoneHandler(c.MaxValue(), c.HasToConfirm(), valueA, valueB)
	}
	return false
}

func (c ScoreCountControl) UpdateScore(scc ScoringCountControl, valueA, valueB *int) {
	if c.UpdateHandler != nil {
		c.UpdateHandler(scc, valueA, valueB)
	}
}
