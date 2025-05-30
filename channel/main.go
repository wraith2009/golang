package main

import (
	"fmt"
	"sync"
)

func unbufferedDemo(wg *sync.WaitGroup) {
	ch := make(chan string)

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("[Unbuffered] Sending candy ...")
		ch <- "candy" // blocks until received
		fmt.Println("[Unbuffered] Candy sent!")
	}()

	fmt.Println("[Unbuffered] Receiving candy ...")
	msg := <-ch
	fmt.Println("[Unbuffered] Received:", msg)
}

func bufferedDemo(wg *sync.WaitGroup) {
	ch := make(chan string, 2)

	// Send first two candies (buffered, won't block)
	fmt.Println("[Buffered] Sending 1st candy ...")
	ch <- "candy1"
	fmt.Println("[Buffered] Sent 1st candy!")

	fmt.Println("[Buffered] Sending 2nd candy ...")
	ch <- "candy2"
	fmt.Println("[Buffered] Sent 2nd candy!")

	// Now receive to make room for third
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("[Buffered] Receiving candy ...")
		fmt.Println("[Buffered] Got:", <-ch)
		fmt.Println("[Buffered] Got:", <-ch)
	}()

	wg.Wait() // wait for both messages to be received before next send

	// Now it's safe to send 3rd candy
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("[Buffered] Sending 3rd candy ...")
		ch <- "candy3" // will succeed because buffer has space now
		fmt.Println("[Buffered] Sent 3rd candy!")
	}()

	// Receive 3rd candy
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("[Buffered] Receiving 3rd candy ...")
		fmt.Println("[Buffered] Got:", <-ch)
	}()
}

func main() {
	var wg sync.WaitGroup

	// fmt.Println("channel ---------->")
	// fmt.Println("=== Unbuffered Channel Demo ===")
	// unbufferedDemo(&wg)
	// wg.Wait() // Wait for unbuffered goroutine to finish

	fmt.Println("\n=== Buffered Channel Demo ===")
	bufferedDemo(&wg)
	wg.Wait() // Wait for all buffered goroutines
}

// Unbuffered channel (default):
	// c := make(chan int)      
	// go func() {
	// 	c <- 10 // blocks until main receives
	// }()
	// fmt.Println(<-c) // prints 10
	
	// // Buffered channel:
	// c2 := make(chan int, 2) 
	// c2 <- 1 // doesn't block
	// c2 <- 2 // fills the buffer, still doesn't block
	// // c2 <- 3 // would block, since buffer size is 2
	// fmt.Println(<-c2) // reads 1
	// fmt.Println(<-c2) // reads 2
	

	// func worker(jobs <-chan int, results chan<- int) {
	// 	for j := range jobs {
	// 		results <- j*2 // send results
	// 	}
	// // }
	// Here jobs is <-chan int (receive-only) and results is chan<- int (send-only). This helps document intent and prevents misuse (the compiler will forbid sends on a <-chan or receives on a chan<-)