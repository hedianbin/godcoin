package cli

import (
	"fmt"
	"github.com/hedianbin/godcoin/wallet"
	"log"
	"github.com/hedianbin/godcoin/chain"
)

func (cli *CLI) getBalance(address, nodeID string) {
	if !wallet.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc, err := chain.LoadBlockChain(nodeID)
	if err != nil{
		log.Panic("ERROR: Load blockchain failed", err)
	}
	balance := bc.GetBalance(address)

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}
