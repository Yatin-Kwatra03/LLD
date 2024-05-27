package thread_pool

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const maxThreadLimit = 5

func ImplementThreadPoolToEfficientlyUseResourcesUnoptimalApproach() {

	// Let's say we are getting a number of requests to print a number
	// we'll try implement it using a thread pool basically limit the
	// no of threads being spawned / go routines being fired.

	wg := &sync.WaitGroup{}
	for idx := 1; idx <= 100; idx++ {

		for runtime.NumGoroutine() == maxThreadLimit {
		}

		wg.Add(1)
		go func(eventData int) {
			defer wg.Done()
			processEvent(eventData)
		}(idx)
	}
	wg.Wait()
	fmt.Println("\nsuccessfully processed all the events!")
}

func processEvent(data int) {
	// let's say processing can take some time
	// so adding a sleep of 1 second
	time.Sleep(1 * time.Second)
	fmt.Println(data)
}
