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

func (ga *GameAction) ExecuteAction(point GamePointing, turn TurnPosition) {
	if point.UpdateScore() == GPUYes {
		ga.sm.UpdateScore(turn)
	}
}
