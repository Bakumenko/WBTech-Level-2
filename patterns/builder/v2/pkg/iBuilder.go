package pkg

import "errors"

type iBuilder interface {
	setWallType()
	setDoorType()
	getHouse() House
}

func GetBuilder(builderType string) (iBuilder, error) {
	if builderType == "wood" {
		return &woodHouseBuilder{}, nil
	}

	if builderType == "stone" {
		return &stoneHouseBuilder{}, nil
	}

	return nil, errors.New("unknown builder")
}
