package main

import (
	"builderV1/pkg"
	"fmt"
)

func main() {
	builder := pkg.NewHouseBuilder()
	builder.SetDoorType("стеклопакет").SetWindowType("деревянные").SetNumFloor(1).SetHasGarden(false)
	builder.SetSwimmingPool(pkg.SwimmingPool{123, 5})
	house := builder.GetHouse()
	fmt.Println(house)
}
