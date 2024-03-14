package main

import "choice_number_game/game"

func main() {
	gameManager := game.NewGameManager()
	gameManager.InitGameManager()

	gameInterface := game.NewGameInterface()
	gameInterface.InitGameInterface(gameManager)
	gameInterface.StartGame()

}
