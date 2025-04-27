package main

//  find sum from 0-X :

// func main() {
// 	var sum = 0
// 	for i := 1; i < 10; i++ {

// 		sum += i
// 		println(`sum at :`, sum, i)
// 	}

// 	println(sum)
// }

//  ranged syntax :

func main(){
	var sum int = 0

	for i:= range 10 {
		sum += i
	}

	println(sum)
}



