package cli

import (
	"fmt"
	"github.com/michain/dotcoin/wallet"
	"log"
	"github.com/michain/dotcoin/server"
)

func (cli *CLI) startNode(nodeID string, isMining bool, minerAddress string, isGenesis bool, listenAddr, seedAddr string) {
	printLogo()

	fmt.Printf("Starting node %s\n", nodeID)
	//nodeID = "3eb456d086f34118925793496cd20945"
	if len(minerAddress) > 0 {
		if wallet.ValidateAddress(minerAddress) {
			fmt.Println("Mining is on. Address to receive rewards: ", minerAddress)
		} else {
			log.Panic("Wrong miner address!")
		}
	}
	if listenAddr == ""{
		listenAddr = tcpPort
	}

	server.StartServer(nodeID, isMining, minerAddress, listenAddr, seedAddr, isGenesis)
}

func printLogo(){
	fmt.Println(` ___   ___       |  __   ___   ___  ___`)
	fmt.Println(`/   \ /   \   ___| /  \ /   \   |  /   \`)
	fmt.Println(`\___||     | /   ||    |     |  |  |   |`)
	fmt.Println(`\___| \___/  \___/ \__/ \___/  _|_ |   |`)
}