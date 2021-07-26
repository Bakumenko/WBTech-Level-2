package pkg

type state interface {
	addItem(int) error
	buyItem() error
	inputMoney(money int) error
	deliveryOfMoney() error
}
