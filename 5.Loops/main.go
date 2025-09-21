package main

import "fmt"

// for -> only construct in go for looping

func main() {


	// While loop
	// i := 1

	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// infinte loop
	// for {
	// 	fmt.Println(1) 
	// }

	//For loop 

	// for i := 0; i < 3; i++ {
	// 	fmt.Println(i)
	// }

	//Range -> new feature 

	for i := range 3 { // 0->3 (3 excluded)
		fmt.Println(i)
	}
	
}