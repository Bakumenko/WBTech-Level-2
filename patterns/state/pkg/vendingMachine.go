package pkg

type VendingMachine struct {
	waitMoneyState       state
	buyItemState         state
	noItemState          state
	deliveryOfMoneyState state

	currentState state
	countItem    int
	itemPrice    int
	countMoney   int
}

func NewVendingMachine(countItem, itemPrice int) *VendingMachine {
	vm := &VendingMachine{
		countItem: countItem,
		itemPrice: itemPrice,
	}

	wms := &waitMoneyState{vm}
	bis := &buyItemState{vm}
	nis := &noItemState{vm}
	doms := &deliveryOfMoneyState{vm}

	vm.waitMoneyState = wms
	vm.buyItemState = bis
	vm.noItemState = nis
	vm.deliveryOfMoneyState = doms

	if vm.countItem == 0 {
		vm.currentState = nis
	} else {
		vm.currentState = wms
	}

	return vm
}

func (v *VendingMachine) AddItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *VendingMachine) BuyItem() error {
	return v.currentState.buyItem()
}

func (v *VendingMachine) InputMoney(money int) error {
	return v.currentState.inputMoney(money)
}

func (v *VendingMachine) GetDelivery() error {
	return v.currentState.deliveryOfMoney()
}
