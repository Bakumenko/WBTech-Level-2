package pkg

import (
	"errors"
	"fmt"
)

type deliveryOfMoneyState struct {
	vendingMachine *VendingMachine
}

func (d *deliveryOfMoneyState) addItem(count int) error {
	return errors.New("Delivery of money")
}

func (d *deliveryOfMoneyState) buyItem() error {
	return errors.New("Delivery of money")
}

func (d *deliveryOfMoneyState) inputMoney(money int) error {
	return errors.New("Delivery of money")
}

func (d *deliveryOfMoneyState) deliveryOfMoney() error {
	fmt.Println("delivery is: ", d.vendingMachine.countMoney)
	d.vendingMachine.countMoney = 0
	if d.vendingMachine.countItem == 0 {
		d.vendingMachine.currentState = d.vendingMachine.noItemState
	} else {
		d.vendingMachine.currentState = d.vendingMachine.waitMoneyState
	}
	return nil
}
