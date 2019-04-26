package main


import "fmt"

func main(){
  var x int = 45
  var y int = 5
  var z float64

  z = float64(x) / float64(y)
  fmt.Printf("%f float value\n",z)
}
