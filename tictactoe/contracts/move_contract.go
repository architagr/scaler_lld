package contracts

type IMove interface {
	GetPlayer() IPlayer
	GetCell() ICell
}
