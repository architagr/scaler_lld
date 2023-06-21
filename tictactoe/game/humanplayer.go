package game

import (
	"fmt"
)

type HumanPlayerDTO struct {
	BasePlayerDTO
}

func (player *HumanPlayerDTO) MakeMove(board *BoardDTO) (row int, col int) {
	fmt.Printf("it is player %q turn enter row and col", player.GetName())
	fmt.Scanf("%d", &row)
	fmt.Scanf("%d", &col)
	return
}

func InitHumanPlayer(symbol *SymbolDTO, name string) IPlayer {
	return &HumanPlayerDTO{
		BasePlayerDTO: BasePlayerDTO{
			symbol:     symbol,
			name:       name,
			playerType: PLAYER_TYPE_HUMAN,
		},
	}
}
