package main

import "fmt"

func findBiggest(args ...int) int {
	biggest := 0

	for _, v := range args {
		if biggest < v {
			biggest = v
		}
	}
return biggest
}

func main(){
	xs:= []int{78,65,69,7888998,411221,1002215,222,3,2,1,1,785999}
	fmt.Println(findBiggest(xs...))
}