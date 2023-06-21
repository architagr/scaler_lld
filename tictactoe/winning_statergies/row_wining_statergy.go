package winningstatergies

import "tictactoe/contracts"

type RowWinningStatergy struct {
	rowData []map[contracts.ISymbol]int
}

func InitRowWinnigStatergy(players []contracts.IPlayer, size int) contracts.IWinningStatergies {
	rowData := make([]map[contracts.ISymbol]int, size)
	for i := 0; i < size; i++ {
		rowData[i] = make(map[contracts.ISymbol]int)
		for _, player := range players {
			rowData[i][player.GetSymbol()] = 0
		}
	}
	return &RowWinningStatergy{
		rowData: rowData,
	}
}
func (winningStatergy *RowWinningStatergy) CheckWinner(board contracts.IBoard, move contracts.IMove) bool {
	winningStatergy.rowData[move.GetCell().GetRow()][move.GetPlayer().GetSymbol()]++
	return winningStatergy.rowData[move.GetCell().GetRow()][move.GetPlayer().GetSymbol()] == board.GetSize()
}
func (winningStatergy *RowWinningStatergy) Undo(board contracts.IBoard, move contracts.IMove) {
	winningStatergy.rowData[move.GetCell().GetRow()][move.GetPlayer().GetSymbol()]--
}
