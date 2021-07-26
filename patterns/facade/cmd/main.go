package main

import (
	"facade/pkg"
	"fmt"
	"log"
)

func main() {
	fmt.Println()
	walletFacade := pkg.NewWalletFacade("123", 1234)
	fmt.Println()

	err := walletFacade.AddMoneyToWallet("123", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.DeductMoneyFromWallet("123", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
