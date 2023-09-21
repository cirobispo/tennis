package tennisstatus

type GamePointType int

type GamePointing interface {
	GetType() GamePointType
	UpdateScore() bool
}

const (
	GPTFoot     GamePointType = 1
	GPTAce      GamePointType = 2
	GPTServeOut GamePointType = 3
	GPTServeNet GamePointType = 4
	GPTServeIn  GamePointType = 5
	GPTReturn   GamePointType = 6
	GPTLet      GamePointType = 7
	GPTNet      GamePointType = 8
	GPTIn       GamePointType = 9
	GPTOut      GamePointType = 10
	GPTToast    GamePointType = 11
	GPTOther    GamePointType = 12
)

type GamePoint struct {
	pointType   GamePointType
	updateScore bool
}

func NewGamePointAce() GamePoint {
	return GamePoint{pointType: GPTAce, updateScore: true}
}

func NewGamePointServeIn() GamePoint {
	return GamePoint{pointType: GPTServeIn, updateScore: false}
}

func NewGamePointReturn() GamePoint {
	return GamePoint{pointType: GPTReturn, updateScore: false}
}

func NewGamePointIn() GamePoint {
	return GamePoint{pointType: GPTIn, updateScore: false}
}

func NewGamePointOut() GamePoint {
	return GamePoint{pointType: GPTOut, updateScore: true}
}

func (g GamePoint) GetType() GamePointType {
	return g.pointType
}

func (g GamePoint) UpdateScore() bool {
	return g.updateScore
}
