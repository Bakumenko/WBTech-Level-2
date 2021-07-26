package pkg

type Rectangle struct {
	Length int
	Width  int
}

func (t *Rectangle) Accept(v Visitor) {
	v.VisitForRectangle(t)
}
