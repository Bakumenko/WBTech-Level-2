package pkg

type Circle struct {
	Radius int
}

func (c *Circle) Accept(v Visitor) {
	v.VisitForCircle(c)
}
