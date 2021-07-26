package pkg

import "fmt"

type House struct {
	windowType   string
	doorType     string
	floor        int
	swimmingPool SwimmingPool
	garden       bool
}

func (h House) String() string {
	return fmt.Sprintf("windowType: %v\n"+
		"doorType: %v\n"+
		"floor: %v\n"+
		"swimmingPool: %v\n"+
		"hasGarden: %v\n",
		h.windowType,
		h.doorType,
		h.floor,
		h.swimmingPool,
		h.garden)
}
