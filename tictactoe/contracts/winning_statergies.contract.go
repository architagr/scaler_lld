package contracts

type IWinningStatergies interface {
	CheckWinner(board IBoard, move IMove) bool
	Undo(board IBoard, move IMove)
}
