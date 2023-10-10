package scoring

type OnScoringChange func(valueA, valueB string)
type OnScoringGame func(valueA, valueB string)

type ScoringType int

const (
	STCustom        ScoringType = 0
	STGame          ScoringType = 1
	STSet           ScoringType = 2
	STMatch         ScoringType = 3
	STTieBreak      ScoringType = 4
	STSuperTieBreak ScoringType = 5
)

type Scoring interface {
	Reset()
	GetStatus() (int, int)
	GetScoringType() ScoringType
	UpdateScore()
	GetScoreCountControl() ScoringCountControl
	AddChangedScoreEvent(scoreChangeEvent OnScoringChange)
	AddReachedScoreEvent(scoreGameEvent OnScoringGame)
}
