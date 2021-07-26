package pkg

type department interface {
	Execute(*Patient)
	SetNext(department)
}
