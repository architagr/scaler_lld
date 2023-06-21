package contracts

import "tictactoe/enums"

type IPlayer interface {
	MakeMove(board IBoard) (row int, col int)
	GetPlayerType() enums.PlayerType
	GetName() string
	GetSymbol() ISymbol
}
