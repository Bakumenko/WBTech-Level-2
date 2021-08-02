package pkg

type director struct {
	builder iBuilder
}

func NewDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) SetBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) BuildHouse() House {
	d.builder.setWallType()
	d.builder.setDoorType()
	return d.builder.getHouse()
}
