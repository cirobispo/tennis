package main

import (
	"cbtennis"
	"fmt"
	"math/rand"
)

func main() {
	//gameScore()
	//setScore()
	//setMatch()
	match := GetMatch()
	game := cbtennis.NewSingleStandardGame(&match)
	gameSimulation(&game)
}

func GetMatch() cbtennis.TennisMatch {
	teamA := cbtennis.NewTeam([]string{"Ciro", "Gabriel"})
	teamB := cbtennis.NewTeam([]string{"Leo", "Lisandra"})

	match := cbtennis.NewTennisMatch(teamA, teamB, 1, 4, false, false)
	//	matchManager := cbtennis.NewMatchManager(&match)
	return match
}

func gameSimulation(game cbtennis.GameManager) {
	service := func() cbtennis.GamePointing {
		switch rand.Intn(5) {
		case 4:
			return cbtennis.NewGamePointAce()
		case 3:
			return cbtennis.NewGamePointServeOut()
		case 2:
			return cbtennis.NewGamePointServeIn()
		case 1:
			return cbtennis.NewGamePointServeNet()
		default:
			return cbtennis.NewGamePointServeLet()
		}
	}

	returnService := func() cbtennis.GamePointing {
		switch rand.Intn(3) {
		case 2:
			return cbtennis.NewGamePointReturnIn()
		case 1:
			return cbtennis.NewGamePointReturnNet()
		default:
			return cbtennis.NewGamePointReturnOut()
		}
	}

	eRally := func() cbtennis.GamePointing {
		switch rand.Intn(3) {
		case 2:
			return cbtennis.NewGamePointIn()
		case 1:
			return cbtennis.NewGamePointNet()
		default:
			return cbtennis.NewGamePointOut()
		}
	}

	game.AddGameStartingEvent(func() {
		fmt.Println("iniciei o jogo")
	})

	sair := false
	game.AddFinishedGameEvent(func() {
		fmt.Println("terminei o jogo")
		sair = true
	})

	updatePointEvent := func(increasedPoint bool, increasedBy cbtennis.TurnPosition) {
		if increasedPoint {
			fmt.Println(increasedBy)
		}
	}

	game.AddUpdatePointEvent(updatePointEvent)

	for {
		b := service()
		game.AddPointing(b)
		if b.GetType() == cbtennis.GPTServeIn {
			returns := returnService()
			game.AddPointing(returns)

			if returns.GetType() == cbtennis.GPTReturnIn {
				for {
					rally := eRally()
					game.AddPointing(rally)
					if rally.GetType() != cbtennis.GPTIn {
						break
					}
				}
			}
		}

		if sair {
			break
		}
	}
}

func setMatch() {
	teamA := cbtennis.NewTeam([]string{"Ciro", "Gabriel"})
	teamB := cbtennis.NewTeam([]string{"Leo", "Lisandra"})

	match := cbtennis.NewTennisMatch(teamA, teamB, 1, 4, false, false)
	matchManager := cbtennis.NewMatchManager(&match)

	rallySimulation(matchManager)
}

func rallySimulation(match cbtennis.MatchManager) {
	currentSet := match.NewSet()

	game := currentSet.NewGame()
	game.AddFinishedGameEvent(func() {
		game = currentSet.NewGame()
	})

	service := func() cbtennis.GamePointing {
		switch rand.Intn(5) {
		case 4:
			return cbtennis.NewGamePointAce()
		case 3:
			return cbtennis.NewGamePointServeOut()
		case 2:
			return cbtennis.NewGamePointServeIn()
		case 1:
			return cbtennis.NewGamePointServeNet()
		default:
			return cbtennis.NewGamePointServeLet()
		}
	}

	returnService := func() cbtennis.GamePointing {
		switch rand.Intn(3) {
		case 2:
			return cbtennis.NewGamePointReturnIn()
		case 1:
			return cbtennis.NewGamePointReturnNet()
		default:
			return cbtennis.NewGamePointReturnOut()
		}
	}

	eRally := func() cbtennis.GamePointing {
		switch rand.Intn(3) {
		case 2:
			return cbtennis.NewGamePointReturnIn()
		case 1:
			return cbtennis.NewGamePointReturnNet()
		default:
			return cbtennis.NewGamePointReturnOut()
		}
	}

	for i := 0; i < 10; i++ {
		b := service()
		game.AddPointing(b)
		if b.GetType() == cbtennis.GPTServeIn {
			returns := returnService()
			game.AddPointing(returns)

			if returns.GetType() == cbtennis.GPTReturnIn {
				for {
					rally := eRally()
					game.AddPointing(rally)
					if rally.GetType() != cbtennis.GPTIn {
						break
					}
				}
			}
		}

		//		scoreData := game.GetScore().(cbtennis.ScoreData)
	}
}

func setScore() {
	sair := false
	largestSet := 0
	score := cbtennis.NewSetScore()
	score.AddReachedScoreEvent(func(valueA, valueB int) {
		sair = true
	})

	score.AddChangedScoreEvent(func() {
		valueA, valueB := score.GetScoreValues(0)
		fmt.Printf("\t%v x %v\n", valueA, valueB)
		if valueA < valueB {
			largestSet = valueB
		}
		largestSet = valueA
	})

	for {
		AB := rand.Intn(2)
		fmt.Printf("AB:%v\t", AB)

		GPD := rand.Intn(2)
		fmt.Printf("\tGPD: %v\t", GPD)

		score.UpdateScore(cbtennis.TurnPosition(AB), cbtennis.GamePointDestination(GPD))
		if largestSet >= 7 {
			fmt.Println()
			break
		}

		if sair {
			score.ResetScore()
			sair = false
		}

	}
}

func gameScore() {
	sair := false
	largestGame := 0
	ballSide := cbtennis.NewTurnManager(cbtennis.TPEven)
	score := cbtennis.NewGameScore()

	score.AddReachedScoreEvent(func(valueA, valueB int) {
		sair = true
	})

	score.AddChangedScoreEvent(func() {
		v1, v2 := score.GetScoreValues(cbtennis.TPEven)
		wrapper := cbtennis.NewScoreDataWrapper(cbtennis.TPEven, score.ScoreType(), v1, v2)

		fmt.Printf("\t%s x %s\n", wrapper.GetValueA(), wrapper.GetValueB())

		if vA, vB := score.GetScoreValues(0); vA < vB {
			largestGame = vB
		} else {
			largestGame = vA
		}
		ballSide.Turn()
	})

	ballSide.AddTurnChangeEvent(func(turn cbtennis.TurnPosition) {
		result := "sacador"
		if turn != cbtennis.TPEven {
			result = "recebedor"
		}

		fmt.Printf("A bola está com o %s\n", result)
	})

	for {
		AB := rand.Intn(2)
		fmt.Print(AB)

		//		score.UpdateScore(cbtennis.TurnPosition(AB))
		if largestGame >= 5 {
			fmt.Println()
			break
		}

		if sair {
			score.ResetScore()
			sair = false
		}

	}
}

func buildAMatch() {
	teamA := cbtennis.NewTeam([]string{"Ciro", "Gabriel"})
	teamB := cbtennis.NewTeam([]string{"Leo", "Lisandra"})

	match := cbtennis.NewTennisMatch(teamA, teamB, 1, 4, false, false)
	matchManager := cbtennis.NewMatchManager(&match)

	currentSet := matchManager.GetCurrentSet()

	currentGame := currentSet.GetCurrentGame()
	currentGame.AddPointing(cbtennis.NewGamePointAce())

	currentGame.AddPointing(cbtennis.NewGamePointServeIn())
	currentGame.AddPointing(cbtennis.NewGamePointReturnIn())
	currentGame.AddPointing(cbtennis.NewGamePointIn())
	currentGame.AddPointing(cbtennis.NewGamePointOut())

	currentGame.AddPointing(cbtennis.NewGamePointServeIn())
	currentGame.AddPointing(cbtennis.NewGamePointReturnIn())
	currentGame.AddPointing(cbtennis.NewGamePointOut())

	currentGame.AddPointing(cbtennis.NewGamePointAce())

	currentGame.AddPointing(cbtennis.NewGamePointServeIn())
	currentGame.AddPointing(cbtennis.NewGamePointReturnIn())
	currentGame.AddPointing(cbtennis.NewGamePointIn())
	currentGame.AddPointing(cbtennis.NewGamePointOut())

	showScore := func() {
		a, b := 0, 0
		fmt.Printf("%v x %v\n", a, b)
	}

	showScore()
}
