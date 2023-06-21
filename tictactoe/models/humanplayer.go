package models

import (
	"fmt"
	"tictactoe/contracts"
	"tictactoe/enums"
)

type HumanPlayerDTO struct {
	BasePlayerDTO
}

func (player *HumanPlayerDTO) MakeMove(board contracts.IBoard) (row int, col int) {
	fmt.Printf("it is player %q turn enter row and col", player.GetName())
	fmt.Scanf("%d", &row)
	fmt.Scanf("%d", &col)
	return
}

func InitHumanPlayer(symbol contracts.ISymbol, name string) contracts.IPlayer {
	return &HumanPlayerDTO{
		BasePlayerDTO: BasePlayerDTO{
			symbol:     symbol,
			name:       name,
			playerType: enums.PLAYER_TYPE_HUMAN,
		},
	}
}
