package main

import (
	"fmt"
	"tictactoe/game"
)

func main() {
	fmt.Println("game start")
	players := []game.IPlayer{
		game.InitHumanPlayer(game.InitSymbol('X'), "Archit"),
		game.InitBotPlayer(game.InitSymbol('O'), "Bot", game.DIFFICULTY_LEVEL_EASY),
	}
	size := 3
	tictoctoeGame := game.InitGame(size, players,
		[]game.IWinningStatergies{
			game.InitRowWinnigStatergy(players, size),
			game.InitColWinnigStatergy(players, size),
		})
	tictoctoeGame.PrintBoard()

	for tictoctoeGame.GetGameState() == game.GAME_STATE_IN_PROGRESS {
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
