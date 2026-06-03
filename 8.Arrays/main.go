package main

import "fmt"

func main() {

	//zeroed values

	var nums [5]int
	//All the element are initialized as 0(in case of INT Arr)
	nums[0] = 1
	fmt.Println(nums)

	
	var vals [3]bool
	fmt.Println(vals)

	var name[3] string 
	name[1] = "Parth"
	fmt.Println(name)
	

	// To declare in a single line
	// 2D also possible -> [2][2]
	nums2 := [3] int {1,2,4}
	nums2D := [2][3] int {{1,2,3}, {1,2,3}}

	fmt.Println(nums2)
	fmt.Println(nums2D)
}