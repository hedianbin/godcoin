package cli

import (
	"fmt"
	"github.com/hedianbin/godcoin/wallet"
	"log"
)

func (cli *CLI) createWallet(nodeID string) {
	wallets, err := wallet.LoadWallets(nodeID)
	if err!=nil{
		log.Panic("ERROR: Load wallets failed", err)
	}
	address := wallets.CreateWallet()
	wallets.SaveToFile()

	fmt.Printf("Your new address: %s\n", address)
}

