package main

import (
	"fmt"

	"github.com/personal-projects/LLD/cache_library/cache"
	"github.com/personal-projects/LLD/tictactoe"
)

func main() {
	//TicTacToeImplementation()
	CacheImplementation()
}

func TicTacToeImplementation() {
	tictactoe.PlayTicTacToeGame()
}

func CacheImplementation() {
	cacheStorage, err := cache.GetCache("cache vendor data", "redis cache")
	if err != nil {
		panic(err)
	}
	fmt.Println(cacheStorage.CurrentCacheSize())

	// todo: add strong testing data here
}
