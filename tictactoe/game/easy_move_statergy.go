package game

type EasyMoveStatergy struct {
}

func (*EasyMoveStatergy) SuggestMove(board *BoardDTO) (row, col int) {
	for rowIdx, rowCells := range board.GetBoardLayout() {
		for colIdx, cell := range rowCells {
			if cell.GetCellStatus() == CELL_STATUS_EMPTY {
				row, col = rowIdx, colIdx
				return
			}
		}
	}
	return
}
