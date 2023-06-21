package enums

type GameState string

var (
	GAME_STATE_IN_PROGRESS GameState = "In progress"
	GAME_STATE_ENDED       GameState = "Ended"
	GAME_STATE_DRAW        GameState = "Draw"
)
