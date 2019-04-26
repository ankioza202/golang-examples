// function return with mutliple values.
package main

import "fmt"

func name(firstname,lastname string) (string, string) {
  return firstname,lastname
}

func main(){
  a, b := name("Ankita","Oza")
  fmt.Println(a,b)
}
