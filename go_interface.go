package main

import "fmt"

type Rectangle struct{
  height float64
  width float64
}

type Circle struct{
  radius float64
}


type Shape interface{
   area() float64
}

func (r1 Rectangle) area() float64 {
  return r1.height*r1.width
}

func (c1 Circle) area() float64 {
  return 10*c1.radius
}

func getArea(s1 Shape) float64 {
  return s1.area()
}

func main(){
  fmt.Println(getArea(Circle{40}))
  fmt.Println(getArea(Rectangle{10,20}))
}
