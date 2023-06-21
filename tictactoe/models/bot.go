package models

import (
	botmovestatergy "tictactoe/bot_move_statergy"
	"tictactoe/contracts"

	"tictactoe/enums"
)

type BotDTO struct {
	BasePlayerDTO
	difficulty   enums.DifficultyLevel
	moveStatergy contracts.IBotMoveStatergy
}

func (bot *BotDTO) GetDifficulty() enums.DifficultyLevel {
	return bot.difficulty
}

func (player *BotDTO) MakeMove(board contracts.IBoard) (row int, col int) {
	return player.moveStatergy.SuggestMove(board)
}

func InitBotPlayer(symbol contracts.ISymbol, name string, diffcultyLevel enums.DifficultyLevel) contracts.IPlayer {
	return &BotDTO{
		BasePlayerDTO: BasePlayerDTO{
			symbol:     symbol,
			name:       name,
			playerType: enums.PLAYER_TYPE_BOT,
		},
		difficulty:   diffcultyLevel,
		moveStatergy: botmovestatergy.BotMoveStatergyFactory(diffcultyLevel),
	}
}
