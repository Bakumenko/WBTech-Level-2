package main

import (
	"builder/pkg"
	"fmt"
)

func main() {
	stoneHBuilder, err := pkg.GetBuilder("stone")
	if err != nil {
		fmt.Println(err)
	}
	woodHBuilder, err := pkg.GetBuilder("wood")
	if err != nil {
		fmt.Println(err)
	}

	director := pkg.NewDirector(stoneHBuilder)
	stoneHouse := director.BuildHouse()
	fmt.Println(stoneHouse)

	director.SetBuilder(woodHBuilder)
	woodHouse := director.BuildHouse()
	fmt.Println(woodHouse)
}
