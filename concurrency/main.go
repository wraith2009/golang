// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func say(wg *sync.WaitGroup, label string) {
// 	defer wg.Done()
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Printf("%s [%d] at %s\n", label, i, time.Now().Format("15:04:05.000"))
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup

// 	// Spawn a goroutine
// 	wg.Add(1)
// 	go say(&wg, "world")

// 	// Run concurrently in main goroutine
// 	wg.Add(1)
// 	say(&wg, "hello")

// 	// Wait for all goroutines to finish
// 	wg.Wait()
// }

package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v:=range s{
		sum += v
	}

	c <- sum
}

func main() {
	s :=  []int{7,2,8,-9,4,0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c , <-c

	fmt.Println(x,y, x+y)
}