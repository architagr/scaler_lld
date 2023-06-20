package models

type BotDTO struct {
	PlayerDTO
	difficulty DifficultyLevel
}

func (bot *BotDTO) GetDifficulty() DifficultyLevel {
	return bot.difficulty
}
func InitBotPlayer(symbol *symbolDTO, name string, diffcultyLevel DifficultyLevel) *BotDTO {
	return &BotDTO{
		PlayerDTO: PlayerDTO{
			symbol:     symbol,
			name:       name,
			playerType: PLAYER_TYPE_BOT,
		},
		difficulty: diffcultyLevel,
	}
}
