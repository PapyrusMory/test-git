package main

import "fmt"

func main(){
	fmt.Print("Convertisson ft en m. Entrez une valeur ")

	var input float64

	fmt.Scanf("%f ", &input)

	output := input * 0.3048

	fmt.Println("Le resultat est ", output)
}