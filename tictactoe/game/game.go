package game

import (
	"fmt"
)

type GameState string

var (
	GAME_STATE_IN_PROGRESS GameState = "In progress"
	GAME_STATE_ENDED       GameState = "Ended"
	GAME_STATE_DRAW        GameState = "Draw"
)

type GameDTO struct {
	moves             []*MoveDTO
	board             *BoardDTO
	players           []IPlayer
	currentPlayer     int
	winningStatergies []IWinningStatergies
	gameState         GameState
	winner            IPlayer
}

func (game *GameDTO) GetMoves() []*MoveDTO {
	return game.moves
}

func (game *GameDTO) GetBoard() *BoardDTO {
	return game.board
}

func (game *GameDTO) GetPlayers() []IPlayer {
	return game.players
}

func (game *GameDTO) GetGameState() GameState {
	return game.gameState
}

func (game *GameDTO) PrintResult() {
	if game.gameState == GAME_STATE_ENDED {
		fmt.Printf("Game has ended\n")
		fmt.Printf("Winner is: %s\n", game.winner.GetName())
	} else {
		fmt.Printf("Game is draw\n")
	}
}
func (game *GameDTO) PrintBoard() {
	game.GetBoard().Print()
}

func (game *GameDTO) validateMove(cell *CellDTO) bool {
	row, col := cell.GetRow(), cell.GetCol()
	if row < 0 || col < 0 || row >= game.board.GetSize() || col >= game.board.GetSize() {
		return false
	}
	return game.board.GetBoardLayout()[row][col].GetCellStatus() == CELL_STATUS_EMPTY
}
func (game *GameDTO) checkDraw() bool {
	if len(game.moves) == (game.board.GetSize()*game.board.GetSize())-len(game.board.GetBlockedCells()) {
		game.gameState = GAME_STATE_DRAW
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
	cellInBoard.SetStatus(CELL_STATUS_FILLED)
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

func (game *GameDTO) checkGameWon(currentPlayer IPlayer, move *MoveDTO) bool {
	for _, winningStatergy := range game.winningStatergies {
		if winningStatergy.CheckWinner(game.board, move) {
			game.gameState = GAME_STATE_ENDED
			game.winner = currentPlayer
			return true
		}
	}
	return false
}

func InitGame(size int, players []IPlayer,
	winningStatergies []IWinningStatergies) GameDTO {
	board := InitBoard(size, []*CellDTO{})
	return GameDTO{
		moves:             []*MoveDTO{},
		board:             board,
		players:           players,
		currentPlayer:     0,
		winningStatergies: winningStatergies,
		gameState:         GAME_STATE_IN_PROGRESS,
		winner:            nil,
	}
}
