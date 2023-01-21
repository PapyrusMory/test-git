package main

import "fmt"

func main(){
	for i:=1; i<=300; i++{
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}