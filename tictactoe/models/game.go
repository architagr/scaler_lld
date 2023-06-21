package models

import (
	"fmt"
	"tictactoe/contracts"
	"tictactoe/enums"
)

type GameDTO struct {
	moves             []contracts.IMove
	board             contracts.IBoard
	players           []contracts.IPlayer
	currentPlayer     int
	winningStatergies []contracts.IWinningStatergies
	gameState         enums.GameState
	winner            contracts.IPlayer
}

func (game *GameDTO) GetMoves() []contracts.IMove {
	return game.moves
}

func (game *GameDTO) GetBoard() contracts.IBoard {
	return game.board
}

func (game *GameDTO) GetPlayers() []contracts.IPlayer {
	return game.players
}

func (game *GameDTO) GetGameState() enums.GameState {
	return game.gameState
}

func (game *GameDTO) PrintResult() {
	if game.gameState == enums.GAME_STATE_ENDED {
		fmt.Printf("Game has ended\n")
		fmt.Printf("Winner is: %s\n", game.winner.GetName())
	} else {
		fmt.Printf("Game is draw\n")
	}
}
func (game *GameDTO) PrintBoard() {
	game.GetBoard().Print()
}

func (game *GameDTO) validateMove(cell contracts.ICell) bool {
	row, col := cell.GetRow(), cell.GetCol()
	if row < 0 || col < 0 || row >= game.board.GetSize() || col >= game.board.GetSize() {
		return false
	}
	return game.board.GetBoardLayout()[row][col].GetCellStatus() == enums.CELL_STATUS_EMPTY
}
func (game *GameDTO) checkDraw() bool {
	if len(game.moves) == (game.board.GetSize()*game.board.GetSize())-len(game.board.GetBlockedCells()) {
		game.gameState = enums.GAME_STATE_DRAW
		return true
	}
	return false
}
func (game *GameDTO) MakeMove() {
	currentPlayer := game.players[game.currentPlayer]

	row, col := currentPlayer.MakeMove(game.board)
	proposedCell := InitCell(row, col)
	fmt.Printf("Move made at row: %d col: %d.\n", proposedCell.GetRow(), proposedCell.GetCol())

	if !game.validateMove(proposedCell) {
		fmt.Printf("Invalid move. Retry.\n")
		return
	}

	cellInBoard := game.board.GetBoardLayout()[proposedCell.GetRow()][proposedCell.GetCol()]
	cellInBoard.SetStatus(enums.CELL_STATUS_FILLED)
	cellInBoard.SetPlayer(currentPlayer)

	move := InitMove(cellInBoard, currentPlayer)
	game.moves = append(game.moves, move)

	if game.checkGameWon(currentPlayer, move) {
		return
	}

	if game.checkDraw() {
		return
	}

	game.currentPlayer = (game.currentPlayer + 1) % len(game.players)
}
func (game *GameDTO) Undo() {
	if len(game.moves) == 0 {
		fmt.Println("no moves to undo")
		return
	}
	lastIndex := len(game.moves) - 1
	lastMove := game.moves[lastIndex]
	for _, winningStatergy := range game.winningStatergies {
		winningStatergy.Undo(game.board, lastMove)
	}

	game.moves = game.moves[:lastIndex]
	lastCell := lastMove.GetCell()
	lastCell.SetStatus(enums.CELL_STATUS_EMPTY)
}
func (game *GameDTO) checkGameWon(currentPlayer contracts.IPlayer, move contracts.IMove) bool {
	for _, winningStatergy := range game.winningStatergies {
		if winningStatergy.CheckWinner(game.board, move) {
			game.gameState = enums.GAME_STATE_ENDED
			game.winner = currentPlayer
			return true
		}
	}
	return false
}

func InitGame(size int, players []contracts.IPlayer,
	winningStatergies []contracts.IWinningStatergies) GameDTO {
	board := InitBoard(size, []contracts.ICell{})
	return GameDTO{
		moves:             []contracts.IMove{},
		board:             board,
		players:           players,
		currentPlayer:     0,
		winningStatergies: winningStatergies,
		gameState:         enums.GAME_STATE_IN_PROGRESS,
		winner:            nil,
	}
}
