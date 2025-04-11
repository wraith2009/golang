package main

// importing multiple packages
import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Random number", rand.Int63n(10))
}