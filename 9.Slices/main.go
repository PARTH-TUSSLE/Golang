// Arrays -> Static
//Slices -> Dynamic -> most used construct in GO

package main

import (
	"fmt"
	"slices"
)

func main() {
	//uninitialised slice is nil
	// var nums[]int	-> nil

	var nums = make([]int, 0, 5 )
	// type, initial size, capacity

	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5) // capacity automatically doubles


	// nums2 := []int{} // copy func doesn't work
	var nums2 = make([]int, len(nums))

	fmt.Println(cap(nums))

	fmt.Println(nums,nums2)

	// COPY function

	copy(nums2, nums)
	fmt.Println(nums,nums2)

	//Slice operation

	fmt.Println(nums[0:3])
	fmt.Println(slices.Equal(nums2, nums)) // in-built slices package 
	
}