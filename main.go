package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sdan/getrekt/contracts/aave"
)

func main() {
	client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/df20e305bfb94881bfa6931e0500ea88")
	if err != nil {
		log.Fatal("connection", err)
	}
	fmt.Println("we have a connection")
	address := common.HexToAddress("0x7d2768de32b0b80b7a3454c06bdac94a69ddc7a9")

	// ap, err := aave.NewAaveLPAP(
	// 	address, client,
	// )
	// if err != nil {
	// 	log.Fatal("ap", err)
	// }
	// pool, err := ap.GetLendingPool(nil)
	// if err != nil {
	// 	log.Fatal("pool", err)
	// }

	lp, err := aave.NewAaveLP(address, client)
	if err != nil {
		log.Fatal("new aave lp", err)
	}

	// incoming := make(chan *aave.AaveLPLiquidationCall)
	incoming := make(chan *aave.AaveLPDeposit)

	// sub, err := lp.WatchLiquidationCall(nil, incoming, nil, nil, nil)
	sub, err := lp.WatchDeposit(nil, incoming, nil, nil, nil)
	if err != nil {
		log.Fatal("sub aave liq call", err)
	}

	for {
		fmt.Println("looking", sub)
		select {
		case err := <-sub.Err():
			log.Fatal("sub aave chan", err)
		case vLog := <-incoming:
			fmt.Println("incoming", vLog)
		}
	}

}
