package tennisstatus

type TennisAction interface {
	ExecuteAction()
}

type GameAction struct {
	sm ScoreManager
}

func NewGameAction(sm ScoreManager) GameAction {
	return GameAction{sm: sm}
}

func (ga *GameAction) ExecuteAction(point GamePointing, turnIndex int) {
	if point.UpdateScore() {
		ga.sm.UpdateScore(turnIndex)
	}
}
