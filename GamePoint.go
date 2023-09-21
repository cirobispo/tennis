package tennisstatus

type GamePointType int
type GamePointUpdate int

type GamePointing interface {
	GetType() GamePointType
	UpdateScore() GamePointUpdate
}

const (
	GPTFoot      GamePointType = 1
	GPTAce       GamePointType = 2
	GPTServeLet  GamePointType = 3
	GPTServeOut  GamePointType = 4
	GPTServeNet  GamePointType = 5
	GPTServeIn   GamePointType = 6
	GPTReturnOut GamePointType = 7
	GPTReturnNet GamePointType = 8
	GPTReturnIn  GamePointType = 9
	GPTNet       GamePointType = 10
	GPTIn        GamePointType = 11
	GPTOut       GamePointType = 12
	GPTToast     GamePointType = 13
	GPTOther     GamePointType = 14
)

const (
	GPUNo          GamePointUpdate = 0
	GPUYes         GamePointUpdate = 1
	GPUCondicional GamePointUpdate = 2
)

type GamePoint struct {
	pointType   GamePointType
	updateScore GamePointUpdate
}

func NewGamePointAce() GamePoint {
	return GamePoint{pointType: GPTAce, updateScore: GPUYes}
}

func NewGamePointServeLet() GamePoint {
	return GamePoint{pointType: GPTServeLet, updateScore: GPUNo}
}

func NewGamePointServeIn() GamePoint {
	return GamePoint{pointType: GPTServeIn, updateScore: GPUNo}
}

func NewGamePointServeOut() GamePoint {
	return GamePoint{pointType: GPTServeOut, updateScore: GPUCondicional}
}

func NewGamePointServeNet() GamePoint {
	return GamePoint{pointType: GPTServeNet, updateScore: GPUCondicional}
}

func NewGamePointReturnOut() GamePoint {
	return GamePoint{pointType: GPTReturnOut, updateScore: GPUYes}
}

func NewGamePointReturnNet() GamePoint {
	return GamePoint{pointType: GPTReturnNet, updateScore: GPUYes}
}

func NewGamePointReturnIn() GamePoint {
	return GamePoint{pointType: GPTReturnIn, updateScore: GPUNo}
}

func NewGamePointIn() GamePoint {
	return GamePoint{pointType: GPTIn, updateScore: GPUNo}
}

func NewGamePointOut() GamePoint {
	return GamePoint{pointType: GPTOut, updateScore: GPUYes}
}

func NewGamePointNet() GamePoint {
	return GamePoint{pointType: GPTNet, updateScore: GPUYes}
}

func (g GamePoint) GetType() GamePointType {
	return g.pointType
}

func (g GamePoint) UpdateScore() GamePointUpdate {
	return g.updateScore
}
