package pkg

type Factory struct {
}

func (f *Factory) CreateTransport(gunType string) *transport {
	if gunType == "ship" {
		return newShip()
	}
	if gunType == "truck" {
		return newTruck()
	}
	return nil
}
