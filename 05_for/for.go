package main

import "fmt"

//for -> only construct for looping in go
func main() {

  // while loop 
//   i:=1
//   for i<=5 {
// 	fmt.Println(i)
// 	i++
//   }

  // infinte loop
//   for {
// 	fmt.Println(1)
//   }

   // classic for loop
   for i:=0;i<3;i++{
	// break

	// if i == 2 {
	// 	continue
	// }
	fmt.Println(i)
   }

   // 1.22 - range

   for i:= range 3 {
	fmt.Println(i)
   }
}