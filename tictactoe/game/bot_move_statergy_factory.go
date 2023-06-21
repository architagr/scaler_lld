package game

func BotMoveStatergyFactory(difficultyLevel DifficultyLevel) IBotMoveStatergy {
	return &EasyMoveStatergy{}
}
