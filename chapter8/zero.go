package main

import "fmt"

/*func zero (x int) {
	x = 0
}

func main(){
	x:=5
	zero(x)
	fmt.Println(x)
}*/

func zero (x *int){
	*x = 0
}

func one(xPtr *int){
	*xPtr = 1
}

func main(){
	x := 5 
	xPtr := new(int)
	zero(&x)
	one(xPtr)
	fmt.Println(x)
	fmt.Println(*xPtr)

}