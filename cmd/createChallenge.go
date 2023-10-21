package cmd

import "cbtennis/internal/player"

func CreateChallenge() player.Challenging {
	// teamA := player.NewTeam([]string{"Ciro", "Leo"})
	// teamB := player.NewTeam([]string{"Jailson", "Ibrahim"})
	teamA := player.NewPlayer("Jailson")
	teamB := player.NewPlayer("Ibrahin")

	return player.NewChallenge(teamA, teamB)
}
