package botmovestatergy

import (
	"tictactoe/contracts"
	"tictactoe/enums"
)

type EasyMoveStatergy struct {
}

func InitEasyMoveStatergy() contracts.IBotMoveStatergy {
	return &EasyMoveStatergy{}
}
func (moveStatergy *EasyMoveStatergy) SuggestMove(board contracts.IBoard) (row, col int) {
	for rowIdx, rowCells := range board.GetBoardLayout() {
		for colIdx, cell := range rowCells {
			if cell.GetCellStatus() == enums.CELL_STATUS_EMPTY {
				row, col = rowIdx, colIdx
				return
			}
		}
	}
	return
}
