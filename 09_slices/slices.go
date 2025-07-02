package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic
// most used construct in go
// useful methods

func main() {
    // unintialized slice is nil

    //var nums[]int
    // fmt.Println(nums == nil)  // -> true
    // fmt.Println(len(nums))

    // var nums = make([]int, 0, 5)
	// capacity -> maximum no. of elements can fit
	// fmt.Println(cap(nums))

	//fmt.Println(nums == nil)   // -> false
    
	// nums = append(nums, 1)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// nums = append(nums, 5)
	// fmt.Println(cap(nums))
	// fmt.Println(nums)

	// slice := []int{}
	// fmt.Println(slice)

	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))

	// copy function
	//copy(nums2,nums)
	//fmt.Println(nums, nums2)

	// slice operator

	var temp = []int{1,2,4}
	fmt.Println(temp[1:2])  // from index 1 to index 2 but exclude index 2
	fmt.Println(temp[:2])   // from starting(index 0) to index 2 but exclude index 2
	fmt.Println(temp[1:])   // from index 1 to last
    
	// slices package
	var num1 = []int{1,2,3}
	var num2 = []int{1,2,4}

	fmt.Println(slices.Equal(num1, num2))

	// 2D
	var nums = [][]int {{1,2,3},{4,6,3}}
    fmt.Println(nums)
}