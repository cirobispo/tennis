package player

type Challenging interface {
}

type Challenge struct {
	defiantA Playing
	defiantB Playing
}

func NewChallenge(defiantA, defiantB Playing) Challenging {
	return &Challenge{
		defiantA: defiantA,
		defiantB: defiantB,
	}
}
