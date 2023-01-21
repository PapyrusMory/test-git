package main

import "fmt"

func main(){
	x:= make([]float64, 5, 10)
	arr := []float64{1,2,3,4,5}
	y := arr[0:5]
	z := arr[0:]
	l := arr[:5]
	p := arr[:]

	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)

	fmt.Println("x",x)
	fmt.Println("arr",arr)
	fmt.Println("y",y)
	fmt.Println("z",z)
	fmt.Println("l", l)
	fmt.Println("p", p)

	fmt.Println(slice1, slice2)
}