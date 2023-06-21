package winningstatergies

import "tictactoe/contracts"

type ColWinningStatergy struct {
	colData []map[contracts.ISymbol]int
}

func InitColWinnigStatergy(players []contracts.IPlayer, size int) contracts.IWinningStatergies {
	colData := make([]map[contracts.ISymbol]int, size)
	for i := 0; i < size; i++ {
		colData[i] = make(map[contracts.ISymbol]int)
		for _, player := range players {
			colData[i][player.GetSymbol()] = 0
		}
	}
	return &ColWinningStatergy{
		colData: colData,
	}
}
func (winningStatergy *ColWinningStatergy) CheckWinner(board contracts.IBoard, move contracts.IMove) bool {
	winningStatergy.colData[move.GetCell().GetCol()][move.GetPlayer().GetSymbol()]++
	return winningStatergy.colData[move.GetCell().GetCol()][move.GetPlayer().GetSymbol()] == board.GetSize()
}
func (winningStatergy *ColWinningStatergy) Undo(board contracts.IBoard, move contracts.IMove) {
	winningStatergy.colData[move.GetCell().GetCol()][move.GetPlayer().GetSymbol()]--
}
