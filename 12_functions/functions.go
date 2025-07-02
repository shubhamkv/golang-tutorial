package main

import "fmt"

func add(a int, b int) int {    // parameter can be written as (a, b int) if both a and b are of same(int here) type
	return a + b
}

func getLanguages() (string, string, bool){    // multiple return values
	return "golang", "typescript", true
}

func processIt(fn func(a int) int) {         // takes function as input
	fn(1)
}

func process() func (a int) int {          // return function
	return func(a int) int {
		return 8
	}
}

func main() {
	res := add(2, 6)
	lang1, lang2, lang3 := getLanguages()

	fn:= func(a int) int{
		return 2
	}

	processIt(fn)

	fn2:= process()
    fn2(2)
	
	fmt.Println(res)
	fmt.Println(lang1, lang2, lang3)
}