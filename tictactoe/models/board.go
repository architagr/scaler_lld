package models

import "fmt"

type BoardDTO struct {
	size         int
	boardLayout  [][]*CellDTO
	blockedCells []*CellDTO
}

func (board *BoardDTO) GetSize() int {
	return board.size
}
func (board *BoardDTO) GetBoardLayout() [][]*CellDTO {
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

func (board *BoardDTO) GetBlockedCells() []*CellDTO {
	return board.blockedCells
}

func InitBoard(size int, blockedCells []*CellDTO) *BoardDTO {
	layout := make([][]*CellDTO, size)

	for i := 0; i < size; i++ {
		layout[i] = make([]*CellDTO, size)
		for j := 0; j < size; j++ {
			layout[i][j] = InitCell(i, j)
		}
	}

	for _, blockedCell := range blockedCells {
		layout[blockedCell.GetRow()][blockedCell.GetCol()].SetStatus(CELL_STATUS_BLOCK)
	}
	return &BoardDTO{
		size:         size,
		boardLayout:  layout,
		blockedCells: blockedCells,
	}
}
