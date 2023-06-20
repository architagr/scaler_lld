package winningstatergies

import (
	models "tictactoe/models"
)

type IWinningStatergies interface {
	CheckWinner(board *models.BoardDTO, move *models.MoveDTO) bool
}
