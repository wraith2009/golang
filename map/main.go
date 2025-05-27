package main

import "fmt"

// type Vertex struct{
// 	Lat, Long float64
// }

// var m map[string]Vertex
func main(){
	fmt.Println("Map ------------>")

	// m = make(map[string]Vertex)
	// m["Bell Labs"] = Vertex{
	// 	40.32424, -53.3232,
	// }

	// fmt.Println(m["Bell Labs"])

	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The Value", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The Value", m["Answer"])

	// delete(m, "Answer")
	// fmt.Println("The Value", m["Answer"])

	// v, ok := m["Answer2"]
	// fmt.Println("The Value", v, "Present?", ok)

	for key, value := range m {
		fmt.Println(key, value)
	}
}