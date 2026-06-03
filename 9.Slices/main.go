// Arrays -> Static
//Slices -> Dynamic -> most used construct in GO

package main

import "fmt"

func main() {
	//uninitialised slice is nil
	// var nums[]int	-> nil

	var nums = make([]int, 2 )
	// type, initial size

	fmt.Println(cap(nums))
	
}