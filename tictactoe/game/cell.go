package game

import (
	"fmt"
)

type CellDTO struct {
	col, row   int
	cellStatus CellStatus
	player     IPlayer
}

func (cell *CellDTO) GetCol() int {
	return cell.col
}

func (cell *CellDTO) GetRow() int {
	return cell.row
}

func (cell *CellDTO) GetCellStatus() CellStatus {
	return cell.cellStatus
}
func (cell *CellDTO) GetPlayer() IPlayer {
	return cell.player
}

func (cell *CellDTO) SetPlayer(player IPlayer) {
	cell.player = player
}

func (cell *CellDTO) SetStatus(status CellStatus) {
	cell.cellStatus = status
}

func (cell *CellDTO) SetCoodinates(row, col int) {
	cell.row, cell.col = row, col
}

func (cell *CellDTO) Print() {
	switch cell.cellStatus {
	case CELL_STATUS_BLOCK:
		fmt.Printf(" - |")
	case CELL_STATUS_EMPTY:
		fmt.Printf(" _ |")
	case CELL_STATUS_FILLED:
		fmt.Printf(" %c |", cell.player.GetSymbol().GetChar())
	}
}

func InitCell(row, col int) *CellDTO {
	return &CellDTO{
		row:        row,
		col:        col,
		cellStatus: CELL_STATUS_EMPTY,
	}
}
