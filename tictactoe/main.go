package main

import (
	"fmt"
	"tictactoe/contracts"
	"tictactoe/enums"
	"tictactoe/models"
	winningstatergies "tictactoe/winning_statergies"
)

func main() {
	fmt.Println("game start")
	players := []contracts.IPlayer{
		models.InitHumanPlayer(models.InitSymbol('X'), "Archit"),
		models.InitBotPlayer(models.InitSymbol('O'), "Bot", enums.DIFFICULTY_LEVEL_EASY),
	}
	size := 3
	tictoctoeGame := models.InitGame(size, players,
		[]contracts.IWinningStatergies{
			winningstatergies.InitRowWinnigStatergy(players, size),
			winningstatergies.InitColWinnigStatergy(players, size),
		})
	tictoctoeGame.PrintBoard()

	for tictoctoeGame.GetGameState() == enums.GAME_STATE_IN_PROGRESS {
		fmt.Printf("do you want to do undo y/n")
		str := ""
		fmt.Scanf("%s", &str)
		if str == "y" {

			continue
		}
		tictoctoeGame.MakeMove()
		tictoctoeGame.PrintBoard()
	}
	tictoctoeGame.PrintResult()
}
