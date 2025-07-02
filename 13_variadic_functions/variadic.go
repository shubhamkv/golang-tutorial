package main

import "fmt"

func sum(nums ...int) int {           // function sum can accept any no. of parameters : variadic function
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	nums:= []int{23, 56, 8, 12}
	res:= sum(nums...)
	//res := sum(23, 57, 1, 19)
	fmt.Println(res)
}