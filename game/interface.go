package game

import (
	"fmt"
)

type GameInterface struct {
	gameManager *GameManager

	cmds map[string]ICommand
}

func NewGameInterface() *GameInterface {
	return new(GameInterface)
}

func (ic *GameInterface) InitGameInterface(gameManager *GameManager) {
	ic.gameManager = gameManager
	ic.InitCommands()
	fmt.Println("Приветствую тебя в игре \" Угадай число! \" \n Правила игры:\n 1.Я загадываю число от 1 до 10, ты должен его угадать\n2.Если ты угадываешь число, ты получаешь 1 очко и мы переходим к следующему раунду\n3.Всего 3 раунда.\n4.Чтобы выиграть, нужно угадать 2 числа")
}

func (ic *GameInterface) InitCommands() {
	ic.cmds = make(map[string]ICommand, 10)

	ic.cmds[gameStatusCmd] = NewCommand[GameStatusCommand]()
	ic.cmds[processNumberCmd] = NewCommand[ProcessNumberCommand]()
	ic.cmds[isPlayerWinsCmd] = NewCommand[CheckIsPlayerWinsCommand]()
	ic.cmds[getCurrentRoundCmd] = NewCommand[GetCurrentRoundCommand]()
	ic.cmds[getCurrentAttemptCmd] = NewCommand[GetCurrentAttemptCommand]()

	for _, v := range ic.cmds {
		v.Init(ic.gameManager)
	}
}

func (ic *GameInterface) StartGame() {
	fmt.Println("Начнём игру! Я уже загадал число, попробуй угадать!")

	for !ic.cmds[gameStatusCmd].Execute(struct{}{}).(bool) {
		fmt.Println(fmt.Sprintf("Номер попытки:%d\n", ic.cmds[getCurrentAttemptCmd].Execute(struct{}{}).(int)))

		val := ic.GetInputValue()

		result := ic.cmds[processNumberCmd].Execute(val).(bool)

		if result == true {
			fmt.Println(fmt.Sprintf("Поздравляем, ты выиграл в этом раунде! %d", val))
			fmt.Println(fmt.Sprintf("Следующий раунд: %d", ic.cmds[getCurrentRoundCmd].Execute(struct{}{}).(int)))
			continue
		}

	}

	isPlayerWin := ic.cmds[isPlayerWinsCmd].Execute(struct{}{}).(bool)
	if isPlayerWin {
		fmt.Println("Поздравляю, ты выиграл!")
		return
	}
	fmt.Println("Ты проиграл :( Попробуй в следующий раз")

}

func (ic *GameInterface) GetInputValue() int {
	for {
		var val int

		fmt.Println("Введи число: ")
		_, err := fmt.Scanf("%d\n", &val)
		if err != nil {
			fmt.Println("Вы ввели неверное число!")
			continue
		}

		return val
	}
}
