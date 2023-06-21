package game

type RowWinningStatergy struct {
	rowData []map[*SymbolDTO]int
}

func InitRowWinnigStatergy(players []IPlayer, size int) IWinningStatergies {
	rowData := make([]map[*SymbolDTO]int, size)
	for i := 0; i < size; i++ {
		rowData[i] = make(map[*SymbolDTO]int)
		for _, player := range players {
			rowData[i][player.GetSymbol()] = 0
		}
	}
	return &RowWinningStatergy{
		rowData: rowData,
	}
}
func (winningStatergy *RowWinningStatergy) CheckWinner(board *BoardDTO, move *MoveDTO) bool {
	winningStatergy.rowData[move.GetCell().GetRow()][move.GetPlayer().GetSymbol()]++
	return winningStatergy.rowData[move.GetCell().GetRow()][move.GetPlayer().GetSymbol()] == board.GetSize()
}
func (winningStatergy *RowWinningStatergy) Undo(board *BoardDTO, move *MoveDTO) {
	winningStatergy.rowData[move.GetCell().GetRow()][move.GetPlayer().GetSymbol()]--
}
