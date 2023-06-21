package models

import (
	"fmt"
	"tictactoe/contracts"
	"tictactoe/enums"
)

type CellDTO struct {
	col, row   int
	cellStatus enums.CellStatus
	player     contracts.IPlayer
}

func (cell *CellDTO) GetCol() int {
	return cell.col
}

func (cell *CellDTO) GetRow() int {
	return cell.row
}

func (cell *CellDTO) GetCellStatus() enums.CellStatus {
	return cell.cellStatus
}
func (cell *CellDTO) GetPlayer() contracts.IPlayer {
	return cell.player
}

func (cell *CellDTO) SetPlayer(player contracts.IPlayer) {
	cell.player = player
}

func (cell *CellDTO) SetStatus(status enums.CellStatus) {
	cell.cellStatus = status
}

func (cell *CellDTO) SetCoodinates(row, col int) {
	cell.row, cell.col = row, col
}

func (cell *CellDTO) Print() {
	switch cell.cellStatus {
	case enums.CELL_STATUS_BLOCK:
		fmt.Printf(" - |")
	case enums.CELL_STATUS_EMPTY:
		fmt.Printf(" _ |")
	case enums.CELL_STATUS_FILLED:
		fmt.Printf(" %c |", cell.player.GetSymbol().GetChar())
	}
}

func InitCell(row, col int) contracts.ICell {
	return &CellDTO{
		row:        row,
		col:        col,
		cellStatus: enums.CELL_STATUS_EMPTY,
	}
}
