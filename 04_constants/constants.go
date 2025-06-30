package main

import "fmt"

const country = "India"

//region:= "urban"     -------> do not allow here

func main() {
	const name = "Shubham"
	const age = 20
	//fmt.Println(country)

	const (
		port = 3000
		host = "localhost"
	)

	fmt.Println(port, host)
}