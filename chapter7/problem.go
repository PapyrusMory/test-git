package main

import "fmt"

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func main(){
	x := []int{1,2,3,4,5}
	fmt.Println(sum(x...))
}