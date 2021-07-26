package pkg

import (
	"fmt"
	"math"
)

type SquareCalculator struct {
	area int
}

func (a *SquareCalculator) VisitForCircle(s *Circle) {
	a.area = int(math.Pi * float64(s.Radius) * float64(s.Radius))
	fmt.Println("Calculating square for circle: ", a.area)
}

func (a *SquareCalculator) VisitForRectangle(s *Rectangle) {
	a.area = s.Length * s.Width
	fmt.Println("Calculating square for rectangle: ", a.area)
}
