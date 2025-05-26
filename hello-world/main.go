// package declaration
package main

// import the "fmt" package for printing
import (
	"fmt"
)

// a function is defined using func keyword following the function name
// a function will have parameters -> a parameter will be defined using parameter name follwing the type
//  then we will have a return type
// func add (x int, y int) int {
// 	return x + y
// }

//  go has a feature that a function can return multiple values
// func swap(x, y string) (string, string) {
// 	return y,x
// }

//  in this format x and y are initialized to 0 that's why we are assigning values to both x and y
// func split (sum int) (x, y int) {
// 	x = sum * 4/9
// 	y = sum - x
// 	return
// }
func main(){
	// using a function from a package
	fmt.Println("hello world")
	// fmt.Println(add(42, 34))
	// a, b := swap("hello", "world")
	// fmt.Println(a,b)
	// fmt.Println(split(17))

	// var i int
	// var f float32
	// var b bool
	// var s string

	// fmt.Printf("%v %v %v %v\n", i, f, b, s)

	//  for loop -------------->
	// sum := 0
	// for i := 0; i< 10; i++ {
	// 	sum += i
	// }


	// while loop ------------>
	// sum := 1
	// for sum < 100 {
	// 	sum += sum
	// 	fmt.Println("sum at : " , sum)
	// }

	// fmt.Println(sum)

	//  switch case ------------->
	// fmt.Print("GO runs on :")

	// switch os := runtime.GOOS; os {
	// case "darwin": 
	// 	fmt.Println("OS X.")
	// case "linux":
	// 	fmt.Println("Linux")
	// default:
	// 	fmt.Printf("%s. \n", os)
	// }


	//  defer --------->
		//  when returning go pops out all of the defer elements -> so flow : counting -> 
		//  all defer values are pushed in stack 
		//  print -> done
		//  now when returning we pop defer values from stack i.e, LIFO (stack)

	// fmt.Println("counting")
	// for i:= 0; i<10; i++ {
	// 	defer fmt.Println(i)
	// }

	// fmt.Println("done")

	//  pointers ---------->
		//  pointers basically contains the address of an variable or other things 
	// i, j := 24, 32

	// p:= &i
	// fmt.Println(p)  //prints the address of i
	// fmt.Println(*p) // prints the value at the address p contains (reads i through p)

	// *p = 21
	// fmt.Println(i)

	// p = &j  // points to j 
	// *p = *p / 37 // divide j through the pointer
	// fmt.Println(j) 

}