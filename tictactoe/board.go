package tictactoe

const (
	// default value that represents that the current board entry is empty / unvisited / unmarked
	unmarkedSpot = -1
)

// board: will be used to play the game
type board struct {
	matrix         [][]int
	boardDimension int
	// stores the no of entries with given id for the given row
	rowMemo []map[int]int
	// stores the no of entries with given id for the given column
	colMemo []map[int]int
	// stores the no of entries with given id for the diagonal with starting from top left
	forwardDiagonal map[int]int
	// stores the no of entries with given id for the diagonal with starting from top right
	reverseDiagonal map[int]int
}

func newBoard(boardDimension int) *board {
	matrix := make([][]int, boardDimension)
	for rowNo := 0; rowNo < boardDimension; rowNo++ {
		matrix[rowNo] = make([]int, boardDimension)

		for colNo := 0; colNo < boardDimension; colNo++ {
			matrix[rowNo][colNo] = unmarkedSpot
		}
	}

	return &board{
		matrix:          matrix,
		boardDimension:  boardDimension,
		rowMemo:         initializeListWithDimension(boardDimension),
		colMemo:         initializeListWithDimension(boardDimension),
		forwardDiagonal: make(map[int]int),
		reverseDiagonal: make(map[int]int),
	}
}

func initializeListWithDimension(entries int) []map[int]int {
	list := make([]map[int]int, entries)
	// now we need to initialize each entry of the list
	for idx := 0; idx < entries; idx++ {
		list[idx] = make(map[int]int)
	}
	return list
}
