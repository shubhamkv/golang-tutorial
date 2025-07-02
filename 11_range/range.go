package main

import "fmt"

// iterating over data structures
func main() {
   nums:= []int{6,7,8}
   m:= map[string]string{"fname":"Shubham", "lname": "Kumar"}

   for i:=0; i<len(nums); i++ {
	fmt.Println(nums[i])
   }

   for i, num:= range nums{
	fmt.Println(i, num)
   }

   for k, v := range m{
	fmt.Println(k,v)
   }
   
   // unicode code point rune
   // i -> starting byte of rune
   // 300 -> 1 byte, 2 byte
   for i,c:= range "Shubham"{
	fmt.Println(i,c)
	fmt.Println(i,string(c))
   }
}