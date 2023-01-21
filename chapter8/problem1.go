package main

import "fmt"

func ex (x *int){
	*x = 1
}

func ey (y *int){
	*y = 2
}

func main(){
	x := 5 
	y := 5

	ex(&y)
	fmt.Println("y",y)

	ey(&x)
	fmt.Println("x",x)
}