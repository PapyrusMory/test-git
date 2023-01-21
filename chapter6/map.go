package main

import "fmt"

func main(){
	/*var x map[string]int
	x["key"] = 10
	fmt.Println(x)*/

	x := make(map[string]int)
	y := make(map[int]int)
	y[1] = 11
	x["key"] = 10

	fmt.Println(x["key"])

	
	fmt.Println(y[1])

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
elements["He"] = "Helium"
elements["Li"] = "Lithium"
elements["Be"] = "Beryllium"
elements["B"] = "Boron"
elements["C"] = "Carbon"
elements["N"] = "Nitrogen"
elements["O"] = "Oxygen"
elements["F"] = "Fluorine"
elements["Ne"] = "Neon"

fmt.Println(elements["Li"])
fmt.Println(elements["Un"])

if name, ok := elements["Un"]; ok {
fmt.Println(name, ok)
}

}