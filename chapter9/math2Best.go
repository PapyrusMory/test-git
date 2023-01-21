package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	x1, y1, x2, y2 float64
}


func distance  (x1,y1,x2,y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1

	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
	l:= distance(r.x1, r.y1, r.x1, r.y2)
	w:= distance(r.x1, r.y1, r.x2, r.y1)
	return l*w
}

func main(){
	r := Rectangle{x1: 4, y1: 6, x2: 9, y2: 3}
	fmt.Println(r.area())
}