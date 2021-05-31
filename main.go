package main

import (
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sdan/getrekt/contracts/aave"
	"github.com/sdan/getrekt/contracts/token"
	"github.com/sdan/getrekt/gateway/twitter"
)

func main() {
	client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/df20e305bfb94881bfa6931e0500ea88")
	if err != nil {
		log.Fatal("connection", err)
	}
	fmt.Println("we have a connection")
	address := common.HexToAddress("0x7d2768de32b0b80b7a3454c06bdac94a69ddc7a9")

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
		select {
		case err := <-sub.Err():
			log.Fatal("sub aave chan", err)
		case vLog := <-incoming:
			token, err := token.NewToken(vLog.Reserve, client)
			if err != nil {
				log.Fatal("token retrieve", err)
			}
			tokenName, err := token.Name(nil)
			if err != nil {
				log.Fatal("token name", err)
			}
			fbalance := new(big.Float)
			fbalance.SetString(vLog.Amount.String())
			tokenAmt := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
			fmt.Printf("amt:%.2f coin: %s, tx: %s\n", tokenAmt, tokenName, vLog.Raw.TxHash.Hex())
			// sendTweet(tokenAmt, tokenName, vLog.Raw.TxHash.Hex())
			twitter.SendTweet(tokenAmt, tokenName, vLog.Raw.TxHash.Hex())
		}
	}

}
