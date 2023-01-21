package main

import (
	"fmt"
	"strconv"
)

func main(){

	hexadecimal_num := "1E"

	decimal_num, err := strconv.ParseInt(hexadecimal_num, 16, 64)

	if err != nil{
		panic(err)
	}

	fmt.Println("Le décimal de", hexadecimal_num, "est", decimal_num)
}