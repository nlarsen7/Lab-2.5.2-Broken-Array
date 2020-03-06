//Nicholas Larsen
//March 5, 2020
//an array that starts at index 1 instead of 0 and fills it with random numbers
package main

import (
  "fmt"
  "math/rand"
  "time"
)
func main() {
  rand.Seed(time.Now().UnixNano())
  var broken [7]int
  //creates array

  for i:=1; i<7; i++{
    broken[i]= rand.Intn((20)+1)
    //loops starting at index 1
  }
  fmt.Println(broken)
  //prints out the array
}