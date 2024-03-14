package game

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MaxRounds         = 3
	MaxChoiceAttempts = 3
)

type GameManager struct {
	hiddenNumber       int
	round              int
	choiceAttemptCount int

	playerGoodAttempts  int
	managerGoodAttempts int
}

func NewGameManager() *GameManager {
	return new(GameManager)
}

func (g *GameManager) InitGameManager() {
	g.round = 1
	g.InitNewNumber()
}

func (g *GameManager) InitNewNumber() {
	rand.Seed(time.Now().UnixNano())
	g.hiddenNumber = rand.Intn(10)
	fmt.Println(fmt.Sprintf("INITED: %d", g.hiddenNumber))
}

func (g *GameManager) ChoiceNumber(n int) bool {
	if n == g.hiddenNumber {
		g.nextRound(true)
		return true
	}
	g.choiceAttemptCount++
	if g.choiceAttemptCount >= MaxChoiceAttempts {
		g.nextRound(false)
	}
	return false
}

func (g *GameManager) nextRound(isPlayer bool) {
	g.choiceAttemptCount = 0
	g.round++
	g.InitNewNumber()
	if isPlayer {
		g.playerGoodAttempts++
		return
	}
	g.managerGoodAttempts++

}

func (g *GameManager) GetCurrentRound() int {
	return g.round
}

func (g *GameManager) GetCurrentChoiceAttempt() int {
	return g.choiceAttemptCount
}

func (g *GameManager) IsGameEnded() bool {
	if g.round >= MaxRounds || g.managerGoodAttempts > g.playerGoodAttempts {
		return true
	}
	return false
}

func (g *GameManager) IsPlayerWins() bool {
	return g.playerGoodAttempts > g.managerGoodAttempts
}

func (g *GameManager) GetCounter() (int, int) {
	return g.playerGoodAttempts, g.managerGoodAttempts
}
