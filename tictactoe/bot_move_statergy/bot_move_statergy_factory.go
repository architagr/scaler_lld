package botmovestatergy

import (
	"tictactoe/contracts"
	"tictactoe/enums"
)

func BotMoveStatergyFactory(difficultyLevel enums.DifficultyLevel) contracts.IBotMoveStatergy {
	return InitEasyMoveStatergy()
}
