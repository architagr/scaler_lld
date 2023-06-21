package contracts

type IBotMoveStatergy interface {
	SuggestMove(board IBoard) (row int, col int)
}
