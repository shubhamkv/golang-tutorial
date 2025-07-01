package main

import "fmt"

func main() {
	// age := 23

	// if age >= 18 {
	// 	fmt.Println("person is adult")
	// } else if age>=13 {
	// 	fmt.Println("person is teenager")
	// } else {
	// 	fmt.Println("person is kid")
	// }

	//  var role = "admin"
	//  var hasPermissions = true

	//  if role == "admin" && hasPermissions {
	// 	fmt.Println("Access Granted!")
	//  } else {
	// 	fmt.Println("Access Denied!")
	//  }

    // we can declare a variable inside if construct
	if age:= 23; age >=18 {
		fmt.Println("Adult")
	} else if age>=12 {
		fmt.Println("Teenager")
	}
    
	// go does not have ternary operator, we will have to use normal if else
}