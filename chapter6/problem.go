package main

import "fmt"

func main(){
	x := []int{
		48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,
	}
	var min int = 0
	for i :=0; i<= 15 ; i++{
		if min > x[i]{
			min = x[i]
		}
		for j := 0; j <= 15; j++{
			if x[i] < x[j] {
				min = x[i]
			} else{
				min = x[j]
			}
		}
	}
fmt.Println(min)
}