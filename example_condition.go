package main

import "fmt"


func main(){
  var x interface{}
  switch i:=x.(type) {
  case nil:
    fmt.Printf("type of x : %T", i)
  case int:
    fmt.Printf("x is int")
  case float64:
    fmt.Printf("x is float64")
  }

}
