package main

import (
	"fmt"
	"state/pkg"
)

func main() {
	machine := pkg.NewVendingMachine(1, 10)
	var err error
	err = machine.BuyItem()
	if err != nil {
		fmt.Println(err)
	}

	err = machine.InputMoney(7)
	if err != nil {
		fmt.Println(err)
	}

	err = machine.InputMoney(7)
	if err != nil {
		fmt.Println(err)
	}

	err = machine.BuyItem()
	if err != nil {
		fmt.Println(err)
	}

	err = machine.GetDelivery()
	if err != nil {
		fmt.Println(err)
	}

	err = machine.InputMoney(10)
	if err != nil {
		fmt.Println(err)
	}

	err = machine.AddItem(2)
	if err != nil {
		fmt.Println(err)
	}

	err = machine.InputMoney(10)
	if err != nil {
		fmt.Println(err)
	}

	err = machine.BuyItem()
	if err != nil {
		fmt.Println(err)
	}

	err = machine.GetDelivery()
	if err != nil {
		fmt.Println(err)
	}

}
