package player

type Challenging interface {
	GetDefiantA() Playing
	GetDefiantB() Playing
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

func (c Challenge) GetDefiantA() Playing {
	return c.defiantA
}

func (c Challenge) GetDefiantB() Playing {
	return c.defiantB
}
