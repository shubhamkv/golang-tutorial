package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dict
func main() {
  // creating map
  m := make(map[string]string)
  m1 := make(map[string]int)

  // setting an element
  m["name"] = "Shubham"
  m["domain"] = "Backend"
  m1["age"] = 23

  // get an element
  fmt.Println(m["name"])
  fmt.Println(m1["age"])

  // Imp -> if key does not exist in map, then it returns zero value 
  fmt.Println(m["phone"])    // here it is empty string
  fmt.Println(m1["phone"])   // here it is 0

  fmt.Println(len(m))
  delete(m, "domain")
  clear(m)

  // another way to create a map
  m2:= map[string]int{"price": 40, "phones": 3}
  m3:= map[string]int{"price": 40, "phones": 3}
  fmt.Println(m2)

  v, ok:= m2["phones"]    // v return the value of k , ok returns true or false depending key exist or not
  fmt.Println(v)
  if ok {
	fmt.Println("all ok")
  } else{
	fmt.Println("not ok")
  }

  fmt.Println(maps.Equal(m2,m3))
}