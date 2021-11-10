package main

import (
	"fmt"
)

type _ interface{
	getArea() float64
}
type triangle struct{
	base float64
	height float64
}
type square struct{
	side float64
}

func main(){
	t := triangle {base: 10, height: 4}
	fmt.Printf("Area of trianle is %v.\n", t.getArea())

	s:= square {10}
	fmt.Printf("Area of trianle is %v.\n", s.getArea())
}

func (t triangle) getArea() float64{
	return (t.base * t.height /2)
}

func (s square) getArea() float64{
	return (s.side * s.side)
}