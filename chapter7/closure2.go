package main

import "fmt"

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func () (ret uint){
		ret = i
		i += 2
		return
	}
}

func main(){
	nextEvent:=makeEvenGenerator()
	fmt.Println(nextEvent())
	fmt.Println(nextEvent())
	fmt.Println(nextEvent())
}