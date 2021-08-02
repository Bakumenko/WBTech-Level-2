package pkg

type woodHouseBuilder struct {
	wallType string
	doorType string
}

func (b *woodHouseBuilder) setWallType() {
	b.wallType = "деревянные"
}

func (b *woodHouseBuilder) setDoorType() {
	b.doorType = "деревянные"
}

func (b *woodHouseBuilder) getHouse() House {
	return House{
		wallType: b.wallType,
		doorType: b.doorType,
	}
}
