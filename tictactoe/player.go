package tictactoe

import "math/rand"

// player :
// 1. a player will participate in the game
// 2. each player will be represented using these standard fields

type player struct {
	id    int
	name  string
	skill string
}

func newPlayer(name, skill string) *player {
	return &player{
		id:    rand.Int(),
		name:  name,
		skill: skill,
	}
}
