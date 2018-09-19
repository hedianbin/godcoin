package cli

import (
	"github.com/hedianbin/godcoin/chain"
	"log"
)

func (cli *CLI) printChain(nodeID string) {
	bc, err := chain.LoadBlockChain(nodeID)
	if err != nil{
		log.Panic("ERROR: Load blockchain failed", err)
	}
	bc.ListBlockHashs()
}

