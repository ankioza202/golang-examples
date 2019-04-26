package main

import "fmt"
import "errors"


func compare_num(num1 int)(int,error){
  if(num1<0){
    return 0,errors.New("Num1 is less than 0")
  }else{
    return num1,nil
  }
}

func main(){
  x := 5
  fmt.Println(compare_num(x))
}
