package game

type BotDTO struct {
	BasePlayerDTO
	difficulty DifficultyLevel
}

func (bot *BotDTO) GetDifficulty() DifficultyLevel {
	return bot.difficulty
}

func (player *BotDTO) MakeMove(board *BoardDTO) (row int, col int) {
	return BotMoveStatergyFactory(player.difficulty).SuggestMove(board)
}

func InitBotPlayer(symbol *SymbolDTO, name string, diffcultyLevel DifficultyLevel) IPlayer {
	return &BotDTO{
		BasePlayerDTO: BasePlayerDTO{
			symbol:     symbol,
			name:       name,
			playerType: PLAYER_TYPE_BOT,
		},
		difficulty: diffcultyLevel,
	}
}
