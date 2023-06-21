package game

type IBotMoveStatergy interface {
	SuggestMove(board *BoardDTO) (row int, col int)
}
