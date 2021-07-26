package pkg

type Visitor interface {
	VisitForCircle(*Circle)
	VisitForRectangle(*Rectangle)
}
