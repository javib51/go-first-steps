/*
Concurrent attempt 1
*/

package main

import (
 "fmt"
)

func worker () <-chan string {
   message := make(chan string)
   go func() {
      message<-"worker\n"
   }()

   return message
}

func main() {
 if true {
  fmt.Println("hello!")
 }

   var n = 3 
   

   for i:= 0; i < n; i++ {
       message:= worker()
      fmt.Println(<-message)
   }
}
