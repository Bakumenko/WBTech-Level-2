package pkg

import "fmt"

type House struct {
	wallType string
	doorType string
}

func (h House) String() string {
	return fmt.Sprintf("wallType: %v\n"+"doorType: %v\n", h.wallType, h.doorType)
}
