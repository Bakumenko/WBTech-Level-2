package pkg

import (
	"errors"
	"fmt"
)

type buyItemState struct {
	vendingMachine *VendingMachine
}

func (b *buyItemState) addItem(count int) error {
	return errors.New("Buying item")
}

func (b *buyItemState) buyItem() error {
	fmt.Println("item price: ", b.vendingMachine.itemPrice)
	b.vendingMachine.countItem--
	b.vendingMachine.countMoney -= b.vendingMachine.itemPrice
	b.vendingMachine.currentState = b.vendingMachine.deliveryOfMoneyState
	fmt.Println("Сount of item: ", b.vendingMachine.countItem)
	return nil
}

func (b *buyItemState) inputMoney(money int) error {
	b.vendingMachine.countMoney += money
	fmt.Println("Сurrent money: ", b.vendingMachine.countMoney)
	return nil
}

func (b *buyItemState) deliveryOfMoney() error {
	b.vendingMachine.countMoney = 0
	fmt.Println("Refund. Current money: ", b.vendingMachine.countMoney)
	return nil
}
