package main

import (
	"fmt"
	"visitor/pkg"
)

func main() {
	circle := &pkg.Circle{Radius: 3}
	rectangle := &pkg.Rectangle{Length: 2, Width: 3}

	squareCalculator := &pkg.SquareCalculator{}

	circle.Accept(squareCalculator)
	rectangle.Accept(squareCalculator)

	fmt.Println()
	perimeterCalculator := &pkg.PerimeterCalculator{}
	circle.Accept(perimeterCalculator)
	rectangle.Accept(perimeterCalculator)
}
