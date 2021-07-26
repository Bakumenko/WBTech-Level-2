package pkg

import "fmt"

type transport struct {
	Name  string
	Speed int
}

func (t *transport) Deliver() {
	fmt.Printf("Deliver by %v with speed %v\n", t.Name, t.Speed)
}
