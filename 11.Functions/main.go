package main

import "fmt"

func sumAndProduc(a int, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}

// Named return values

func divide (a int, b int) (ans int, remainder int) {
	ans = a/b
	remainder = a%b
	return // naked return returns the latest/current values of the named variables 
}

// Variadic functions

func sumAll (nums ...int) int {
	total := 0
	for _ ,currVal := range nums {
		total = total + currVal
	}
	return total
}



func main() {
	s, p := sumAndProduc(2, 4)
	fmt.Println(s, p);
// if we dont want to return a value 
	onlySum, _ := sumAndProduc(2, 4)
	fmt.Println(onlySum)

	fmt.Println(divide(10, 3))

	fmt.Println(sumAll(1,2,3,4,5,6,7,8,9,10))
	// for slices
	slice1 := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(sumAll(slice1...))


// Anonymous func // IIFE (if invoked instantly)

	res := func (n int ) int{
 		return n*2
	}
	fmt.Println(res(21))
}

