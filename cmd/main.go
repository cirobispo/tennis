package main

import (
	"fmt"
	"math/rand"
	"tennisstatus"
)

func main() {
	//gameScore()
	//setScore()
	//setMatch()
	match := GetMatch()
	game := tennisstatus.NewSingleStandardGame(&match)
	gameSimulation(&game)
}

func GetMatch() tennisstatus.TennisMatch {
	teamA := tennisstatus.NewTeam([]string{"Ciro", "Gabriel"})
	teamB := tennisstatus.NewTeam([]string{"Leo", "Lisandra"})

	match := tennisstatus.NewTennisMatch(teamA, teamB, 1, 4, false, false)
	//	matchManager := tennisstatus.NewMatchManager(&match)
	return match
}

func gameSimulation(game tennisstatus.GameManager) {
	service := func() tennisstatus.GamePointing {
		switch rand.Intn(5) {
		case 4:
			return tennisstatus.NewGamePointAce()
		case 3:
			return tennisstatus.NewGamePointServeOut()
		case 2:
			return tennisstatus.NewGamePointServeIn()
		case 1:
			return tennisstatus.NewGamePointServeNet()
		default:
			return tennisstatus.NewGamePointServeLet()
		}
	}

	returnService := func() tennisstatus.GamePointing {
		switch rand.Intn(3) {
		case 2:
			return tennisstatus.NewGamePointReturnIn()
		case 1:
			return tennisstatus.NewGamePointReturnNet()
		default:
			return tennisstatus.NewGamePointReturnOut()
		}
	}

	eRally := func() tennisstatus.GamePointing {
		switch rand.Intn(3) {
		case 2:
			return tennisstatus.NewGamePointIn()
		case 1:
			return tennisstatus.NewGamePointNet()
		default:
			return tennisstatus.NewGamePointOut()
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

	game.AddUpdatePointEvent(func(increasedPoint bool, increasedBy tennisstatus.TurnPosition) {
		if increasedPoint {
			g := game.(*tennisstatus.StandardGame)
			v1, v2 := g.GetScoreData().GetScoreValues(increasedBy)
			data := tennisstatus.NewScoreDataWrapper(tennisstatus.TPEven, g.GetScoreData().ScoreType(), v1, v2)
			fmt.Printf("%s x %s\n", data.GetValueA(), data.GetValueB())
		}
	})

	for {
		b := service()
		game.AddPointing(b)
		if b.GetType() == tennisstatus.GPTServeIn {
			returns := returnService()
			game.AddPointing(returns)

			if returns.GetType() == tennisstatus.GPTReturnIn {
				for {
					rally := eRally()
					game.AddPointing(rally)
					if rally.GetType() != tennisstatus.GPTIn {
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
	teamA := tennisstatus.NewTeam([]string{"Ciro", "Gabriel"})
	teamB := tennisstatus.NewTeam([]string{"Leo", "Lisandra"})

	match := tennisstatus.NewTennisMatch(teamA, teamB, 1, 4, false, false)
	matchManager := tennisstatus.NewMatchManager(&match)

	rallySimulation(matchManager)
}

func rallySimulation(match tennisstatus.MatchManager) {
	currentSet := match.NewSet()

	game := currentSet.NewGame()
	game.AddFinishedGameEvent(func() {
		game = currentSet.NewGame()
	})

	service := func() tennisstatus.GamePointing {
		switch rand.Intn(5) {
		case 4:
			return tennisstatus.NewGamePointAce()
		case 3:
			return tennisstatus.NewGamePointServeOut()
		case 2:
			return tennisstatus.NewGamePointServeIn()
		case 1:
			return tennisstatus.NewGamePointServeNet()
		default:
			return tennisstatus.NewGamePointServeLet()
		}
	}

	returnService := func() tennisstatus.GamePointing {
		switch rand.Intn(3) {
		case 2:
			return tennisstatus.NewGamePointReturnIn()
		case 1:
			return tennisstatus.NewGamePointReturnNet()
		default:
			return tennisstatus.NewGamePointReturnOut()
		}
	}

	eRally := func() tennisstatus.GamePointing {
		switch rand.Intn(3) {
		case 2:
			return tennisstatus.NewGamePointReturnIn()
		case 1:
			return tennisstatus.NewGamePointReturnNet()
		default:
			return tennisstatus.NewGamePointReturnOut()
		}
	}

	for i := 0; i < 10; i++ {
		b := service()
		game.AddPointing(b)
		if b.GetType() == tennisstatus.GPTServeIn {
			returns := returnService()
			game.AddPointing(returns)

			if returns.GetType() == tennisstatus.GPTReturnIn {
				for {
					rally := eRally()
					game.AddPointing(rally)
					if rally.GetType() != tennisstatus.GPTIn {
						break
					}
				}
			}
		}

		//		scoreData := game.GetScore().(tennisstatus.ScoreData)
	}
}

func setScore() {
	sair := false
	largestSet := 0
	score := tennisstatus.NewSetScore()
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
		fmt.Print(AB)

		score.UpdateScore(tennisstatus.TurnPosition(AB))
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
	ballSide := tennisstatus.NewTurnManager(tennisstatus.TPEven)
	score := tennisstatus.NewGameScore()

	score.AddReachedScoreEvent(func(valueA, valueB int) {
		sair = true
	})

	score.AddChangedScoreEvent(func() {
		v1, v2 := score.GetScoreValues(tennisstatus.TPEven)
		wrapper := tennisstatus.NewScoreDataWrapper(tennisstatus.TPEven, score.ScoreType(), v1, v2)

		fmt.Printf("\t%s x %s\n", wrapper.GetValueA(), wrapper.GetValueB())

		if vA, vB := score.GetScoreValues(0); vA < vB {
			largestGame = vB
		} else {
			largestGame = vA
		}
		ballSide.Turn()
	})

	ballSide.AddTurnChangeEvent(func(turn tennisstatus.TurnPosition) {
		result := "sacador"
		if turn != tennisstatus.TPEven {
			result = "recebedor"
		}

		fmt.Printf("A bola está com o %s\n", result)
	})

	for {
		AB := rand.Intn(2)
		fmt.Print(AB)

		score.UpdateScore(tennisstatus.TurnPosition(AB))
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
	teamA := tennisstatus.NewTeam([]string{"Ciro", "Gabriel"})
	teamB := tennisstatus.NewTeam([]string{"Leo", "Lisandra"})

	match := tennisstatus.NewTennisMatch(teamA, teamB, 1, 4, false, false)
	matchManager := tennisstatus.NewMatchManager(&match)

	currentSet := matchManager.GetCurrentSet()

	currentGame := currentSet.GetCurrentGame()
	currentGame.AddPointing(tennisstatus.NewGamePointAce())

	currentGame.AddPointing(tennisstatus.NewGamePointServeIn())
	currentGame.AddPointing(tennisstatus.NewGamePointReturnIn())
	currentGame.AddPointing(tennisstatus.NewGamePointIn())
	currentGame.AddPointing(tennisstatus.NewGamePointOut())

	currentGame.AddPointing(tennisstatus.NewGamePointServeIn())
	currentGame.AddPointing(tennisstatus.NewGamePointReturnIn())
	currentGame.AddPointing(tennisstatus.NewGamePointOut())

	currentGame.AddPointing(tennisstatus.NewGamePointAce())

	currentGame.AddPointing(tennisstatus.NewGamePointServeIn())
	currentGame.AddPointing(tennisstatus.NewGamePointReturnIn())
	currentGame.AddPointing(tennisstatus.NewGamePointIn())
	currentGame.AddPointing(tennisstatus.NewGamePointOut())

	showScore := func() {
		a, b := 0, 0
		fmt.Printf("%v x %v\n", a, b)
	}

	showScore()
}
