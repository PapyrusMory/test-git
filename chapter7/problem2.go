package main

import "fmt"

func main() {

	var input int

	fmt.Println("Entrer un nombre")

	fmt.Scanf("%f", &input)
	 

	if input % 2 == 0 {
		fmt.Println("Le nombre entré est pair")
	} else if input % 2 == 1 {
		fmt.Println("Le nombre entré est impair")
	}

}