package contracts

import "tictactoe/enums"

type ICell interface {
	GetCol() int
	GetRow() int
	GetCellStatus() enums.CellStatus
	GetPlayer() IPlayer
	SetPlayer(player IPlayer)
	SetStatus(status enums.CellStatus)
	SetCoodinates(row, col int)
	Print()
}
