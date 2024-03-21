package tictactoe

import (
	"errors"
	"fmt"
)

func PlayTicTacToeGame() {
	defer func() {
		fmt.Println("Game over guys! see you later")
	}()

	var (
		players, boardDimension int
	)
	fmt.Print("Enter no of players to play the game : ")
	_, err := fmt.Scan(&players)
	if err != nil {
		fmt.Println("error reading input", err)
		return
	}

	fmt.Print("Enter length of the board required : ")
	_, err = fmt.Scan(&boardDimension)
	if err != nil {
		fmt.Println("error reading input", err)
		return
	}

	game := initialiseGameEntities(players, boardDimension)

	gameOver := false

	// represents name of the winner player
	// or notifies draw in case of a draw
	var result string
	for turnIdx := 0; !gameOver; turnIdx = (turnIdx + 1) % players {

		fmt.Println("turn of the player with name", game.players[turnIdx].name)

		var x, y int
		fmt.Print("Write x and y coordinate for the move : ")
		_, err := fmt.Scan(&x, &y)
		if err != nil {
			fmt.Println("error reading input", err)
			return
		}

		if err = isValidMove(x, y, game); err != nil {
			fmt.Println(err)
			turnIdx = turnIdx - 1
			continue
		}

		gameOver, result = game.playMove(turnIdx, x-1, y-1)
		if gameOver {
			if result == draw {
				fmt.Println("game ended in a draw, so everyone is a winner!")
			} else {
				fmt.Println("winner of the game is ", result)
			}
		}
	}
}

// isValidMove : validates if the move that player is trying to make is valid
func isValidMove(x, y int, game *game) error {
	boardDimension := game.board.boardDimension

	// check if coordinates are inside the board
	if !(x >= 1 && x <= boardDimension && y >= 1 && y <= boardDimension) {
		return errors.New("please enter a coordinate within the board")
	}

	// check if coordinates have not been visited yet
	if game.board.matrix[x-1][y-1] != unmarkedSpot {
		return errors.New("Already marked, please enter a different coordinate ")
	}

	return nil
}

func initialiseGameEntities(playersCount, dimensions int) *game {
	board := newBoard(dimensions)
	game := newGame(board)

	for idx := 0; idx < playersCount; idx++ {
		game.addPlayerToGame(newPlayer(fmt.Sprintf("player %v", idx+1), "ok ok"))
	}
	return game
}
