package pkg

import (
	"errors"
	"fmt"
)

type waitMoneyState struct {
	vendingMachine *VendingMachine
}

func (w *waitMoneyState) addItem(count int) error {
	return errors.New("Item is already contains")
}

func (w *waitMoneyState) buyItem() error {
	return errors.New("Input money")
}

func (w *waitMoneyState) inputMoney(money int) error {
	w.vendingMachine.countMoney += money
	if w.vendingMachine.countMoney > w.vendingMachine.itemPrice {
		w.vendingMachine.currentState = w.vendingMachine.buyItemState
	}
	fmt.Println("current money: ", w.vendingMachine.countMoney)
	return nil
}

func (w *waitMoneyState) deliveryOfMoney() error {
	w.vendingMachine.countMoney = 0
	return nil
}
