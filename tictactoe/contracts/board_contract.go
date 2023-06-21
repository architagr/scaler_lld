package contracts

type IBoard interface {
	GetSize() int
	GetBoardLayout() [][]ICell
	Print()
	GetBlockedCells() []ICell
}
