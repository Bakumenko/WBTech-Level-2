package pkg

type HouseBuilder struct {
	windowType   string
	doorType     string
	floor        int
	swimmingPool SwimmingPool
	hasGarden    bool
}

func NewHouseBuilder() *HouseBuilder {
	return &HouseBuilder{}
}

func (b *HouseBuilder) SetWindowType(wtype string) *HouseBuilder {
	b.windowType = wtype
	return b
}

func (b *HouseBuilder) SetDoorType(dtype string) *HouseBuilder {
	b.doorType = dtype
	return b
}

func (b *HouseBuilder) SetNumFloor(fnum int) *HouseBuilder {
	b.floor = fnum
	return b
}

func (b *HouseBuilder) SetSwimmingPool(pool SwimmingPool) *HouseBuilder {
	b.swimmingPool = pool
	return b
}

func (b *HouseBuilder) SetHasGarden(f bool) *HouseBuilder {
	b.hasGarden = f
	return b
}

func (b *HouseBuilder) GetHouse() House {
	return House{
		windowType:   b.windowType,
		doorType:     b.doorType,
		floor:        b.floor,
		swimmingPool: b.swimmingPool,
		garden:       b.hasGarden,
	}
}
