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
	cacheStorage, err := cache.GetCache("cache vendor data", "redis cache", 3)
	if err != nil {
		panic(err)
	}

	fmt.Println("\nBeware, don't say anything bad, everything is being cached and will be remembered!")

	for {
		fmt.Println("1. Add a key value pair")
		fmt.Println("2. Get value for a key")
		fmt.Println("3. Update value of a key")
		fmt.Println("4. Delete a key")
		fmt.Println("5. Current Cache size")
		fmt.Println("6. Remaining Cache capacity")
		fmt.Println("7. Exit")

		fmt.Print("Choose operation you wish to perform : ")

		var (
			operationType int32
			key, value    string
		)

		if _, err = fmt.Scan(&operationType); err != nil {
			fmt.Println(err)
			continue
		}

		// for better readability
		fmt.Println()

		switch operationType {
		case 1:
			fmt.Print("Enter the key value pair you want to add: ")
			if _, err = fmt.Scan(&key, &value); err != nil {
				fmt.Println(err)
				break
			}
			if err = cacheStorage.Set(key, value); err != nil {
				fmt.Println(err)
			}
		case 2:
			fmt.Print("Enter the key you want to fetch the value for: ")

			if _, err = fmt.Scan(&key); err != nil {
				fmt.Println(err)
				break
			}
			if value, err = cacheStorage.Get(key); err != nil {
				fmt.Println(err)
			}
			fmt.Println(value)

		case 3:
			fmt.Print("Enter the key and the value you want to update: ")
			if _, err = fmt.Scan(&key, &value); err != nil {
				fmt.Println(err)
				break
			}
			if err = cacheStorage.Update(key, value); err != nil {
				fmt.Println(err)
			}
		case 4:
			fmt.Print("Enter the key you want to delete: ")

			if _, err = fmt.Scan(&key); err != nil {
				fmt.Println(err)
				break
			}
			if err = cacheStorage.Delete(key); err != nil {
				fmt.Println(err)
			}
		case 5:
			cacheSize, err := cacheStorage.CurrentCacheSize()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Current cache size is : ", cacheSize)
		case 6:
			remainingCacheCapacity, err := cacheStorage.RemainingCacheCapacity()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Remaining cache capacity is : ", remainingCacheCapacity)
		case 7:
			fmt.Println("It was nice having cache session with you!")
			return
		default:
			fmt.Println("Please choose a valid option")
		}
		fmt.Println()
	}

}
