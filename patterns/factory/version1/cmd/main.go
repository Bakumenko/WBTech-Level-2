package main

import (
	"factory_v1/pkg"
)

func main() {
	factoty := pkg.Factory{}
	ship := factoty.CreateTransport("ship")
	truck := factoty.CreateTransport("truck")
	ship.Deliver()
	truck.Deliver()
}
