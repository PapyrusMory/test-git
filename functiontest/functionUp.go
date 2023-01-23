package main

import (
	"fmt"
	"strconv"
)

func getUp(str string) int {
	getStrElt := (string(str[5]))
	getNumber, err := strconv.Atoi(getStrElt)

	if err != nil{
		panic(err)
	}
	return getNumber
}

func main(){
	fmt.Println(getUp("(low,3)"))
	for i := 1; i<= getUp("(low,3)"); i++{
		fmt.Println(i)
	}
}
