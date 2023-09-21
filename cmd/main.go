package main

import (
	"fmt"
	"math/rand"
	"tennisstatus"
)

func main() {
	// gameScore()
	// setScore()
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

		score.UpdateScore(AB)
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
	score := tennisstatus.NewGameScore()
	score.AddReachedScoreEvent(func(valueA, valueB int) {
		sair = true
	})

	score.AddChangedScoreEvent(func() {
		wrapper := tennisstatus.NewScoreDataWrapper(0, score)

		fmt.Printf("\t%s x %s\n", wrapper.GetValueA(), wrapper.GetValueB())

		if vA, vB := score.GetScoreValues(0); vA < vB {
			largestGame = vB
		} else {
			largestGame = vA
		}
	})

	for {
		AB := rand.Intn(2)
		fmt.Print(AB)

		score.UpdateScore(AB)
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

	matchManager.StartMatch()
	currentSet := matchManager.GetCurrentSet()

	currentGame := currentSet.GetCurrentGame()
	currentGame.AddPointing(tennisstatus.NewGamePointAce())

	currentGame.AddPointing(tennisstatus.NewGamePointServeIn())
	currentGame.AddPointing(tennisstatus.NewGamePointReturn())
	currentGame.AddPointing(tennisstatus.NewGamePointIn())
	currentGame.AddPointing(tennisstatus.NewGamePointOut())

	currentGame.AddPointing(tennisstatus.NewGamePointServeIn())
	currentGame.AddPointing(tennisstatus.NewGamePointReturn())
	currentGame.AddPointing(tennisstatus.NewGamePointOut())

	currentGame.AddPointing(tennisstatus.NewGamePointAce())

	currentGame.AddPointing(tennisstatus.NewGamePointServeIn())
	currentGame.AddPointing(tennisstatus.NewGamePointReturn())
	currentGame.AddPointing(tennisstatus.NewGamePointIn())
	currentGame.AddPointing(tennisstatus.NewGamePointOut())

	showScore := func() {
		a, b := 0, 0
		fmt.Printf("%v x %v\n", a, b)
	}

	showScore()
}
