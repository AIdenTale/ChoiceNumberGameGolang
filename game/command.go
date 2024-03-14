package game

func NewCommand[T any]() *T {
	return new(T)
}

const (
	gameStatusCmd        = "gameStatusCmd"
	processNumberCmd     = "processNumberCmd"
	isPlayerWinsCmd      = "isPlayerWinsCmd"
	getCurrentRoundCmd   = "getCurrentRoundCmd"
	getCurrentAttemptCmd = "getCurrentAttemptCmd"
)

type Command struct {
	manager *GameManager
}

func (c *Command) Init(gameManager *GameManager) {
	c.manager = gameManager
}

type ICommand interface {
	Init(gameManager *GameManager)
	Execute(params any) any
}

type GameStatusCommand struct {
	Command
}

func (c *GameStatusCommand) Execute(params any) any {
	return c.manager.IsGameEnded()
}

type ProcessNumberCommand struct {
	Command
}

func (c *ProcessNumberCommand) Execute(params any) any {
	num, ok := params.(int)
	if !ok {
		panic("Вы передали не число!")
	}
	return c.manager.ChoiceNumber(num)
}

type CheckIsPlayerWinsCommand struct {
	Command
}

func (c *CheckIsPlayerWinsCommand) Execute(params any) any {
	return c.manager.IsPlayerWins()
}

type GetGameCounterCommand struct {
	Command
}

func (c *GetGameCounterCommand) Execute(params any) any {
	counter := make([]int, 0, 1)
	counter[0], counter[1] = c.manager.GetCounter()
	return counter
}

type GetCurrentRoundCommand struct {
	Command
}

func (c *GetCurrentRoundCommand) Execute(params any) any {
	return c.manager.GetCurrentRound()
}

type GetCurrentAttemptCommand struct {
	Command
}

func (c *GetCurrentAttemptCommand) Execute(params any) any {
	return c.manager.GetCurrentChoiceAttempt()
}
