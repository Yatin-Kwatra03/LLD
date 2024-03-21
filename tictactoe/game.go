package tictactoe

const (
	draw = "draw"
)

// game : will be player b/w players on a board
type game struct {
	players            []*player
	board              *board
	noOfMovesPerformed int
}

func newGame(board *board) *game {
	return &game{
		// we will add support to add players separately
		players:            make([]*player, 0),
		board:              board,
		noOfMovesPerformed: 0,
	}
}

// addPlayerToGame: adds player to the game
// NOTE - will be helpful if we need to support
// to add players at any point of time in the game.
func (s *game) addPlayerToGame(player *player) {
	s.players = append(s.players, player)
}

// returns a boolean that represents if the game is over after the current move and notifies the winner / draw
func (s *game) playMove(playerIdx, x, y int) (bool, string) {
	// update the move details in our storage entities
	s.storeMoveDetails(playerIdx, x, y)

	// check if player won after this move
	if s.didPlayerWinAfterThisMove(playerIdx, x, y) {
		return true, s.players[playerIdx].name
	}

	// check if game ended in a draw
	if s.hasGameEndedInADraw() {
		return true, draw
	}

	// game is still on!
	return false, ""
}

func (s *game) hasGameEndedInADraw() bool {
	return s.noOfMovesPerformed == (s.board.boardDimension * s.board.boardDimension)
}

func (s *game) didPlayerWinAfterThisMove(playerIdx, x, y int) bool {
	playerId := s.players[playerIdx].id
	boardDimension := s.board.boardDimension

	// checks if whole xth row is covered by playerIdx
	if count, ok := s.board.rowMemo[x][playerId]; ok && count == boardDimension {
		return true
	}

	// checks if whole yth column is covered by playerIdx
	if count, ok := s.board.colMemo[x][playerId]; ok && count == boardDimension {
		return true
	}

	if isCoordinateOnForwardDiagonal(x, y) {
		// checks if whole forward diagonal is covered by playerIdx
		if count, ok := s.board.forwardDiagonal[playerId]; ok && count == boardDimension {
			return true
		}
	}

	if isCoordinateOnReverseDiagonal(x, y, boardDimension) {
		// checks if whole reverse diagonal is covered by playerIdx
		if count, ok := s.board.reverseDiagonal[playerId]; ok && count == boardDimension {
			return true
		}
	}

	return false
}

func (s *game) storeMoveDetails(playerIdx, x, y int) {
	// update total moves
	s.noOfMovesPerformed = s.noOfMovesPerformed + 1

	playerId := s.players[playerIdx].id

	// update the matrix
	s.board.matrix[x][y] = playerId

	// update row cache
	s.board.rowMemo[x][playerId] = s.board.rowMemo[x][playerId] + 1
	// update column cache
	s.board.colMemo[y][playerId] = s.board.colMemo[y][playerId] + 1

	// update forward diagonal cache
	if isCoordinateOnForwardDiagonal(x, y) {
		s.board.forwardDiagonal[playerId] = s.board.forwardDiagonal[playerId] + 1
	}

	// update reverse diagonal cache
	if isCoordinateOnReverseDiagonal(x, y, s.board.boardDimension) {
		s.board.reverseDiagonal[playerId] = s.board.reverseDiagonal[playerId] + 1
	}
}

func isCoordinateOnForwardDiagonal(x, y int) bool {
	return x == y
}

func isCoordinateOnReverseDiagonal(x, y, dimension int) bool {
	return x+y == dimension-1
}
