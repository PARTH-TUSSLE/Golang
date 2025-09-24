package main

import (
	"fmt"
	// "time"
)

func main() {


	//Single SWITCH STATEMENT
	// i := 3

	// switch i {
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// }


	//Multiple SWITCH STATEMENT

	// switch time.Now().Weekday() {
	// 	case time.Saturday, time.Sunday : 
	// 	fmt.Println("It's a Weekend")

	// 	default : 
	// 	fmt.Println("It's a Week day")



	//TYPE SWITCH

	whoAmI := func (i interface{})  {
		switch i.(type) {
		case int:
			fmt.Println("It is an Integer!")
		case bool:
			fmt.Println("It is a Bollean!")
		case string:
			fmt.Println("It is a String!")
		default :
			fmt.Println("Others")
	}

	}

	whoAmI(34.34)
}
