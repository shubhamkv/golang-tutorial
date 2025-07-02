package main

import "fmt"

func main() {
	// zeroed values
	// int -> 0, string -> "", bool -> false

	var nums [4]int
	nums[0] = 21

	fmt.Println(nums[0])
	fmt.Println(nums)
	fmt.Println(len(nums))
    
	// to declare array elements in single line
	arrayList := [3]int {1,2,3}
	fmt.Println(arrayList)

	// 2D array
	twoDArray := [2][2]int{{2,3},{4,6}}
	fmt.Println(twoDArray)

	// - fixed aize, that is predictable
	// - Memory optimization
	// - Constant time access
}