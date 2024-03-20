package tictactoe

// game : will be player b/w players on a board
type game struct {
	players []*player
	board   *board
}

func newGame(board *board) *game {
	return &game{
		// we will add support to add players separately
		players: make([]*player, 0),
		board:   board,
	}
}

// addPlayerToGame: adds player to the game
// NOTE - will be helpful if we need to support
// to add players at any point of time in the game.
func (s *game) addPlayerToGame(player *player) {
	s.players = append(s.players, player)
}

// playerIdx : turn of the playerIdx player
// dx : net change in coordinate x
// dy : net change in coordinate y
// returns a boolean that represents if the game is over after the current move
func (s *game) playMove(playerIdx, dx, dy int) bool {
	// implement move logic
	return false
}
