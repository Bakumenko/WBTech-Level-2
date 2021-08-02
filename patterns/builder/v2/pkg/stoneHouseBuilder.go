package pkg

type stoneHouseBuilder struct {
	wallType string
	doorType string
}

func (b *stoneHouseBuilder) setWallType() {
	b.wallType = "каменные"
}

func (b *stoneHouseBuilder) setDoorType() {
	b.doorType = ""
}

func (b *stoneHouseBuilder) getHouse() House {
	return House{
		wallType: b.wallType,
		doorType: b.doorType,
	}
}
