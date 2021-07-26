package pkg

import (
	"errors"
	"fmt"
)

type noItemState struct {
	vendingMachine *VendingMachine
}

func (n *noItemState) addItem(count int) error {
	n.vendingMachine.countItem += count
	n.vendingMachine.currentState = n.vendingMachine.buyItemState
	fmt.Println("Added item")
	return nil
}

func (n *noItemState) buyItem() error {
	return errors.New("Need to add item")
}

func (n *noItemState) inputMoney(money int) error {
	return errors.New("Need to add item")
}

func (n *noItemState) deliveryOfMoney() error {
	return errors.New("Need to add item")
}
