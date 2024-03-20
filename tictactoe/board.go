package tictactoe

const (
	// default value that represents that the current board entry is empty / unvisited / unmarked
	unmarkedSpot = -1
)

// board: will be used to play the game
type board struct {
	matrix [][]int
	rows   int
	cols   int
}

func newBoard(rows, cols int) *board {
	matrix := make([][]int, rows)
	for rowNo := 0; rowNo < rows; rowNo++ {
		matrix[rowNo] = make([]int, cols)

		for colNo := 0; colNo < cols; colNo++ {
			matrix[rowNo][colNo] = unmarkedSpot
		}
	}

	return &board{
		matrix: matrix,
		rows:   rows,
		cols:   cols,
	}

}
