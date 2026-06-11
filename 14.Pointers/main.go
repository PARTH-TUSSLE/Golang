package main

import (
	"fmt"
)

func main() {
	// stores the mem address of any value
	// &x -> address of x (makes a pointer)
	// *p -> deref operator (go to that address and perform CRUD)
	// used fo changing a value, inside a func, without returning 
	score := 10
	fmt.Println("Before", score)
	addScore(&score)
	fmt.Println("After", score)

}

func addScore ( score *int )  {
	*score = *score+5
}