package models

import (
	"fmt"
	"tictactoe/contracts"
	"tictactoe/enums"
)

type BoardDTO struct {
	size         int
	boardLayout  [][]contracts.ICell
	blockedCells []contracts.ICell
}

func (board *BoardDTO) GetSize() int {
	return board.size
}
func (board *BoardDTO) GetBoardLayout() [][]contracts.ICell {
	return board.boardLayout
}

func (board *BoardDTO) Print() {
	for _, row := range board.boardLayout {
		fmt.Printf("|")
		for _, cell := range row {
			cell.Print()
		}
		fmt.Printf("\n")
	}
}

func (board *BoardDTO) GetBlockedCells() []contracts.ICell {
	return board.blockedCells
}

func InitBoard(size int, blockedCells []contracts.ICell) contracts.IBoard {
	layout := make([][]contracts.ICell, size)

	for i := 0; i < size; i++ {
		layout[i] = make([]contracts.ICell, size)
		for j := 0; j < size; j++ {
			layout[i][j] = InitCell(i, j)
		}
	}

	for _, blockedCell := range blockedCells {
		layout[blockedCell.GetRow()][blockedCell.GetCol()].SetStatus(enums.CELL_STATUS_BLOCK)
	}
	return &BoardDTO{
		size:         size,
		boardLayout:  layout,
		blockedCells: blockedCells,
	}
}
