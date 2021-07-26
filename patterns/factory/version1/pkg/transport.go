package pkg

import "fmt"

type transport struct {
	name  string
	speed int
}

func (t *transport) Deliver() {
	fmt.Printf("Deliver by %v with speed %v\n", t.name, t.speed)
}
