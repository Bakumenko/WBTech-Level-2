package pkg

import (
	"fmt"
	"math"
)

type PerimeterCalculator struct {
	perimeter int
}

func (p *PerimeterCalculator) VisitForCircle(s *Circle) {
	p.perimeter = int(math.Pi * float64(2) * float64(s.Radius))
	fmt.Println("Calculating perimeter for circle: ", p.perimeter)
}
func (p *PerimeterCalculator) VisitForRectangle(s *Rectangle) {
	p.perimeter = 2 * (s.Length + s.Width)
	fmt.Println("Calculating perimeter for rectangle: ", p.perimeter)
}
