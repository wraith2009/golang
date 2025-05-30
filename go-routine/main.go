package main

import (
	"fmt"
	"sync"
	"time"
)

// worker simulates a task by printing start and end messages with a delay.
//
// Parameters:
//   - id: An integer identifier for the worker.
//	 - wg: A pointer to a sync.WaitGroup used to wait for all goroutines to finish.

// This function is intended to run as a goroutine to simulate concurrent execution.
func worker(id int, wg *sync.WaitGroup) {

	defer wg.Done()
	
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // Simulates work with a 1-second delay
	fmt.Printf("Worker %d done\n", id)
}


// main is the entry point of the program.
//
// It launches two goroutines running the `worker` function with different IDs.
// Then it sleeps for 2 seconds to ensure both workers finish before the program exits.
//
// NOTE: In production code, it's better to use sync.WaitGroup to synchronize goroutines
// rather than using time.Sleep.
func main() {
	
	// go func() {
    // 	fmt.Println("Hello from a goroutine")
	// }()

	// for i := 0; i < 3; i++ {
	// 	go func() {
	// 		fmt.Println(i) // bug: i may be 3 for all goroutines
	// 	}()
	// }

	// for i := 0; i < 3; i++ {
	// 	go func(n int) {
	// 		fmt.Println(n) // correctly prints 0,1,2 (in some order)
	// 	}(i)
	// }
	
	var wg sync.WaitGroup

	wg.Add(2) // we're launching 2 workers
	
	go worker(1, &wg) // launch worker 1 in a new goroutine
	go worker(2, &wg) // launch worker 2 concurrently

	wg.Wait()  // wait for all worker go routing to complete

	fmt.Println("All workers launched")
}


// output (order may vary)

/*
*Worker 2 starting
*Worker 1 starting
*Worker 1 done
*Worker 2 done
*All workers launched
*/

/*
*The Go scheduler multiplexes goroutines onto OS threads.
*As one resource notes: “There might be only one thread in a program with thousands of goroutines.
*Instead, goroutines are multiplexed dynamically onto threads as needed”. 
*This means the Go runtime manages how goroutines map to CPU threads, allowing cheap concurrency. 
*You can control parallelism via runtime.GOMAXPROCS, but by default it uses all CPU cores.
*/

/*
*A goroutine runs until its function returns. 
*You generally cannot kill a goroutine from the outside (there’s no Stop() function). 
*Use context cancellation or other signaling to request termination.
* Always ensure goroutines eventually exit, or you’ll leak resources (the goroutine keeps executing or blocking forever).
*/