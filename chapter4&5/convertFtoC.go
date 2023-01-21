package main

import "fmt"

func main(){
	fmt.Print("Entrer une valeur à convertir en C")
	var input float64
	fmt.Scanf("%f F", &input)

	output := (input - 32)* 5/9

	fmt.Println("La valeur en Celsius est", output)
}