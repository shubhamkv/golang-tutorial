package main

import (
	"fmt"
)

func main() {
	// i := 3
	// switch i {
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// default:
	// 	fmt.Println("Other")
	// }
	
	// multiple switch condition

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("Weekend")
    // default:
	// 	fmt.Println("Workday")
	// }

	// type switch

	whoAmI := func (i interface{}) {
		switch i.(type){
		case int:
			fmt.Println("Integer")
		case string:
			fmt.Println("String")
		case bool:
			fmt.Println("Boolean")
	    default:
			fmt.Println("other")
		}
	}
	whoAmI("golang")
}