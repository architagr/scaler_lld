package models

import "fmt"

type PlayerDTO struct {
	symbol     *symbolDTO
	name       string
	playerType PlayerType
}

func (player *PlayerDTO) GetSymbol() *symbolDTO {
	return player.symbol
}

func (player *PlayerDTO) GetName() string {
	return player.name
}

func (player *PlayerDTO) GetPlayerType() PlayerType {
	return player.playerType
}

func (player *PlayerDTO) MakeMove() *CellDTO {
	cell := InitCell(0, 0)
	var row, col int
	fmt.Printf("it is player %q turn enter row and col", player.GetName())
	fmt.Scanf("%d", &row)
	fmt.Scanf("%d", &col)
	return cell
}

func InitHumanPlayer(symbol *symbolDTO, name string) *PlayerDTO {
	return &PlayerDTO{
		symbol:     symbol,
		name:       name,
		playerType: PLAYER_TYPE_HUMAN,
	}
}

type PlayerType string

var (
	PLAYER_TYPE_HUMAN PlayerType = "HUMAN"
	PLAYER_TYPE_BOT   PlayerType = "BOT"
	PLAYER_TYPE_GUEST PlayerType = "GUEST"
)

type DifficultyLevel string

var (
	DIFFICULTY_LEVEL_EASY   DifficultyLevel = "EASY"
	DIFFICULTY_LEVEL_MEDIUM DifficultyLevel = "MEDIUM"
	DIFFICULTY_LEVEL_HARD   DifficultyLevel = "HARD"
)
