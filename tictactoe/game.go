package main

import (
	"fmt"
	"tictactoe/models"
	winningstatergies "tictactoe/winning_statergies"
)

type GameState string

var (
	GAME_STATE_IN_PROGRESS GameState = "In progress"
	GAME_STATE_ENDED       GameState = "Ended"
	GAME_STATE_DRAW        GameState = "Draw"
)

type gameDTO struct {
	moves             []*models.MoveDTO
	board             *models.BoardDTO
	players           []*models.PlayerDTO
	currentPlayer     int
	winningStatergies []winningstatergies.IWinningStatergies
	gameState         GameState
	winner            *models.PlayerDTO
}

func (game *gameDTO) GetMoves() []*models.MoveDTO {
	return game.moves
}

func (game *gameDTO) GetBoard() *models.BoardDTO {
	return game.board
}

func (game *gameDTO) GetPlayers() []*models.PlayerDTO {
	return game.players
}

func (game *gameDTO) GetGameState() GameState {
	return game.gameState
}

func (game *gameDTO) PrintResult() {
	if game.gameState == GAME_STATE_ENDED {
		fmt.Printf("Game has ended\n")
		fmt.Printf("Winner is: %s\n", game.winner.GetName())
	} else {
		fmt.Printf("Game is draw\n")
	}
}
func (game *gameDTO) PrintBoard() {
	game.GetBoard().Print()
}

func (game *gameDTO) validateMove(cell *models.CellDTO) bool {
	row, col := cell.GetRow(), cell.GetCol()
	if row < 0 || col < 0 || row >= game.board.GetSize() || col >= game.board.GetSize() {
		return false
	}
	return game.board.GetBoardLayout()[row][col].GetCellStatus() == models.CELL_STATUS_EMPTY
}
func (game *gameDTO) checkDraw() bool {
	if len(game.moves) == (game.board.GetSize()*game.board.GetSize())-len(game.board.GetBlockedCells()) {
		game.gameState = GAME_STATE_DRAW
		return true
	}
	return false
}
func (game *gameDTO) makeMove() {
	currentPlayer := game.players[game.currentPlayer]

	proposedCell := currentPlayer.MakeMove()

	fmt.Printf("Move made at row: %d col: %d.\n", proposedCell.GetRow(), proposedCell.GetCol())

	if !game.validateMove(proposedCell) {
		fmt.Printf("Invalid move. Retry.\n")
		return
	}

	cellInBoard := game.board.GetBoardLayout()[proposedCell.GetRow()][proposedCell.GetCol()]
	cellInBoard.SetStatus(models.CELL_STATUS_FILLED)
	cellInBoard.SetPlayer(currentPlayer)

	move := models.InitMove(cellInBoard, currentPlayer)
	game.moves = append(game.moves, move)

	if game.checkGameWon(currentPlayer, move) {
		return
	}

	if game.checkDraw() {
		return
	}

	game.currentPlayer = (game.currentPlayer + 1) % len(game.players)
}

func (game *gameDTO) checkGameWon(currentPlayer *models.PlayerDTO, move *models.MoveDTO) bool {
	for _, winningStatergy := range game.winningStatergies {
		if winningStatergy.CheckWinner(game.board, move) {
			game.gameState = GAME_STATE_ENDED
			game.winner = currentPlayer
			return true
		}
	}
	return false
}
