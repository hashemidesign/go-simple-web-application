package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK         = 0
	PAPER        = 1
	SCISSORS     = 2
	PLAYERWINS   = 1
	COMPUTERWINS = 2
	DRAW         = 3
)

type Round struct {
	Winner        int    `json:"winner"`
	ComputerCoice string `json:"computer_choice"`
	RoundResult   string `json:"round_result"`
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""

	winner := 0

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
		break
	case PAPER:
		computerChoice = "Computer chose PAPER"
		break
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
		break
	default:
	}

	if playerValue == computerValue {
		roundResult = "It's a draw"
		winner = DRAW
	} else if playerValue == (computerValue+3)%3 {
		roundResult = "Player wins"
		winner = PLAYERWINS
	} else {
		roundResult = "Computer wins"
		winner = COMPUTERWINS
	}

	var result Round
	result.Winner = winner
	result.ComputerCoice = computerChoice
	result.RoundResult = roundResult

	return result
}
