package main

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sdan/getrekt/contracts/aave"
	"github.com/sdan/getrekt/contracts/compound"
	"github.com/sdan/getrekt/contracts/token"
	// "github.com/sdan/getrekt/gateway/twitter"
)

func main() {
	client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/df20e305bfb94881bfa6931e0500ea88")
	if err != nil {
		log.Fatal("connection", err)
	}
	fmt.Println("we have a connection")
	aaveV2address := common.HexToAddress("0x7d2768de32b0b80b7a3454c06bdac94a69ddc7a9")
	cETHaddress := common.HexToAddress("0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5")
	cBATaddress := common.HexToAddress("0x6c8c6b02e7b2be14d4fa6022dfd6d75921d90e4e")
	cCOMPaddress := common.HexToAddress("0x70e36f6bf80a52b3b46b3af8e106cc0ed743e8e4")
	cDAIaddress := common.HexToAddress("0x5d3a536e4d6dbd6114cc1ead35777bab948e3643")
	// cLINKaddress := common.HexToAddress("0xface851a4921ce59e912d19329929ce6da6eb0c7")
	// cREPaddress := common.HexToAddress("0x158079ee67fce2f58472a96584a73c7ab9ac95c1")
	// cSAIaddress := common.HexToAddress("0xf5dce57282a584d2746faf1593d3121fcac444dc")
	// cTUSDaddress := common.HexToAddress("0x12392f67bdf24fae0af363c24ac620a2f67dad86")
	// cUNIaddress := common.HexToAddress("0x35a18000230da775cac24873d00ff85bccded550")
	// cUSDCaddress := common.HexToAddress("0x39aa39c021dfbae8fac545936693ac917d5e7563")
	// cUSDTaddress := common.HexToAddress("0xf650c3d88d12db855b8bf7d11be6c55a4e07dcc9")
	// cWBTCaddress := common.HexToAddress("0xc11b1268c1a384e55c48c2391d8d480264a3a7f4")
	// cWBTC2address := common.HexToAddress("0xccf4429db6322d5c611ee964527d42e5d685dd6a")
	// cZRXaddress := common.HexToAddress("0xb3319f5d18bc0d84dd1b4825dcde5d5f7266d407")

	aavelp, err := aave.NewAaveLP(aaveV2address, client)
	if err != nil {
		log.Fatal("new aave lp", err)
	}
	compoundcETHClient, err := compound.NewCETH(cETHaddress, client)
	if err != nil {
		log.Fatal("new compound", err)
	}
	compoundcBATClient, err := compound.NewCBAT(cBATaddress, client)
	if err != nil {
		log.Fatal("new compound", err)
	}
	compoundcCOMPClient, err := compound.NewCCOMP(cCOMPaddress, client)
	if err != nil {
		log.Fatal("new compound", err)
	}
	compoundcDAIClient, err := compound.NewCDAI(cDAIaddress, client)
	if err != nil {
		log.Fatal("new compound", err)
	}

	aaveliq := make(chan *aave.AaveLPLiquidationCall)
	// aavedep := make(chan *aave.AaveLPDeposit)
	// aavebrw := make(chan *aave.AaveLPBorrow)

	compoundliq := make(chan *compound.CETHLiquidateBorrow)
	// compoundmin := make(chan *compound.CETHMint)

	compoundliqcBAT := make(chan *compound.CBATLiquidateBorrow)
	compoundliqcCOMP := make(chan *compound.CCOMPLiquidateBorrow)
	compoundliqcDAI := make(chan *compound.CDAILiquidateBorrow)

	go watchAaveLiquidation(aaveliq, aavelp, client)
	//go watchAaveDeposit(aavedep, aavelp, client)
	// go watchAaveBorrow(aavebrw, aavelp, client)
	go watchCompoundcETHLiquidation(compoundliq, compoundcETHClient)
	// go watchCompoundcETHMint(compoundmin, compoundcETHClient)
	go watchCompoundcBATLiquidation(compoundliqcBAT, compoundcBATClient)
	go watchCompoundcCOMPLiquidation(compoundliqcCOMP, compoundcCOMPClient)
	go watchCompoundcDAILiquidation(compoundliqcDAI, compoundcDAIClient)

	<-aaveliq
	// <-aavedep
	// <-aavebrw
	<-compoundliq
	// <-compoundmin
	<-compoundliqcBAT
	<-compoundliqcCOMP
	<-compoundliqcDAI

}

func watchAaveLiquidation(incoming chan *aave.AaveLPLiquidationCall, lp *aave.AaveLP, client *ethclient.Client) {
	sub, err := lp.WatchLiquidationCall(nil, incoming, nil, nil, nil)
	if err != nil {
		log.Fatal("sub aave liq call", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal("sub aave chan", err)
		case vLog := <-incoming:
			fmt.Println("tx hash", vLog.Raw.TxHash.Hex())
			token, err := token.NewToken(vLog.CollateralAsset, client)
			if err != nil {
				log.Fatal("token retrieve", err)
			}
			tokenName, err := token.Name(nil)
			if err != nil {
				tokenName = "Maker"
			}
			fbalance := new(big.Float)
			fbalance.SetString(vLog.LiquidatedCollateralAmount.String())
			tokenAmt := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
			if tokenName == "USDC" || tokenName == "cUSDC" || tokenName == "USDT" || tokenName == "cUSDT" {
				tokenAmt = new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(6)))
			} else if tokenName == "WBTC" || tokenName == "cWBTC" {
				tokenAmt = new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(8)))
			}
			fmt.Printf("Liquidated amt:%.2f coin: %s, tx: %s\n", tokenAmt, tokenName, vLog.Raw.TxHash.Hex())
			// twitter.SendTweet("Liquidated", tokenAmt, tokenName, vLog.Raw.TxHash.Hex(), "Aave")
		}
	}
}

func convertBytesToString(b []byte) string {
	s := make([]string, len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s, ",")
}

func watchAaveDeposit(incoming chan *aave.AaveLPDeposit, lp *aave.AaveLP, client *ethclient.Client) {
	sub, err := lp.WatchDeposit(nil, incoming, nil, nil, nil)
	if err != nil {
		log.Fatal("sub aave liq call", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal("sub aave chan DEPOSIT", err)
		case vLog := <-incoming:
			fmt.Println("tx hash", vLog.Raw.TxHash.Hex())
			token, err := token.NewToken(vLog.Reserve, client)
			if err != nil {
				log.Fatal("token retrieve", err)
			}
			tokenName, err := token.Name(nil)
			if err != nil {
				tokenName = "Maker"
			}
			fbalance := new(big.Float)
			fbalance.SetString(vLog.Amount.String())
			tokenAmt := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
			if tokenName == "USDC" || tokenName == "cUSDC" || tokenName == "USDT" || tokenName == "cUSDT" {
				tokenAmt = new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(6)))
			} else if tokenName == "WBTC" || tokenName == "cWBTC" {
				tokenAmt = new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(8)))
			}
			fmt.Printf("Deposit amt:%.2f coin: %s, tx: %s\n", tokenAmt, tokenName, vLog.Raw.TxHash.Hex())
			// twitter.SendTweet("Deposited", tokenAmt, tokenName, vLog.Raw.TxHash.Hex(), "Aave")
		}
	}
}

// func watchAaveBorrow(incoming chan *aave.AaveLPBorrow, lp *aave.AaveLP, client *ethclient.Client) {
// 	sub, err := lp.WatchBorrow(nil, incoming, nil, nil, nil)
// 	if err != nil {
// 		log.Fatal("sub aave liq call", err)
// 	}

// 	for {
// 		select {
// 		case err := <-sub.Err():
// 			log.Fatal("sub aave chan", err)
// 		case vLog := <-incoming:
// 			fmt.Println("tx hash", vLog.Raw.TxHash.Hex())
// 			token, err := token.NewToken(vLog.Reserve, client)
// 			if err != nil {
// 				log.Fatal("token retrieve", err)
// 			}
// 			tokenName, err := token.Name(nil)
// 			if err != nil {
// 				log.Fatal("token name", err)
// 			}
// 			fbalance := new(big.Float)
// 			fbalance.SetString(vLog.Amount.String())
// 			tokenAmt := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
// 			fmt.Printf("Borrowed amt:%.2f coin: %s, tx: %s\n", tokenAmt, tokenName, vLog.Raw.TxHash.Hex())
// 			// twitter.SendTweet("Borrowed", tokenAmt, tokenName, vLog.Raw.TxHash.Hex(), "Aave")
// 		}
// 	}
// }

func watchCompoundcETHLiquidation(incoming chan *compound.CETHLiquidateBorrow, lp *compound.CETH) {
	sub, err := lp.WatchLiquidateBorrow(nil, incoming)
	if err != nil {
		log.Fatal("sub aave liq call", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal("sub aave chan", err)
		case vLog := <-incoming:
			fmt.Println("tx hash", vLog.Raw.TxHash.Hex())
			fbalance := new(big.Float)
			fbalance.SetString(vLog.RepayAmount.String())
			tokenAmt := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
			fmt.Printf("Liquidated amt:%.2f coin: %s, tx: %s\n", tokenAmt, "cETH", vLog.Raw.TxHash.Hex())
			// twitter.SendTweet("Liquidated", tokenAmt, "cETH", vLog.Raw.TxHash.Hex(), "Compound")
		}
	}
}

func watchCompoundcETHMint(incoming chan *compound.CETHMint, lp *compound.CETH) {
	sub, err := lp.WatchMint(nil, incoming)
	if err != nil {
		log.Fatal("sub aave liq call", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal("sub aave chan", err)
		case vLog := <-incoming:
			fmt.Println("tx hash", vLog.Raw.TxHash.Hex())
			fbalance := new(big.Float)
			fbalance.SetString(vLog.MintAmount.String())
			tokenAmt := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
			fmt.Printf("Minted amt:%.2f coin: %s, tx: %s\n", tokenAmt, "cETH", vLog.Raw.TxHash.Hex())
			// twitter.SendTweet("Minted", tokenAmt, "cETH", vLog.Raw.TxHash.Hex(), "Compound")
		}
	}
}

func watchCompoundcDAILiquidation(incoming chan *compound.CDAILiquidateBorrow, lp *compound.CDAI) {
	sub, err := lp.WatchLiquidateBorrow(nil, incoming)
	if err != nil {
		log.Fatal("sub aave liq call", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal("sub aave chan", err)
		case vLog := <-incoming:
			fmt.Println("tx hash", vLog.Raw.TxHash.Hex())
			fbalance := new(big.Float)
			fbalance.SetString(vLog.RepayAmount.String())
			tokenAmt := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
			fmt.Printf("Liquidated amt:%.2f coin: %s, tx: %s\n", tokenAmt, "cDAI", vLog.Raw.TxHash.Hex())
			// twitter.SendTweet("Liquidated", tokenAmt, "cDAI", vLog.Raw.TxHash.Hex(), "Compound")
		}
	}
}

func watchCompoundcBATLiquidation(incoming chan *compound.CBATLiquidateBorrow, lp *compound.CBAT) {
	sub, err := lp.WatchLiquidateBorrow(nil, incoming)
	if err != nil {
		log.Fatal("sub aave liq call", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal("sub aave chan", err)
		case vLog := <-incoming:
			fmt.Println("tx hash", vLog.Raw.TxHash.Hex())
			fbalance := new(big.Float)
			fbalance.SetString(vLog.RepayAmount.String())
			tokenAmt := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
			fmt.Printf("Liquidated amt:%.2f coin: %s, tx: %s\n", tokenAmt, "cBAT", vLog.Raw.TxHash.Hex())
			// twitter.SendTweet("Liquidated", tokenAmt, "cBAT", vLog.Raw.TxHash.Hex(), "Compound")
		}
	}
}

func watchCompoundcCOMPLiquidation(incoming chan *compound.CCOMPLiquidateBorrow, lp *compound.CCOMP) {
	sub, err := lp.WatchLiquidateBorrow(nil, incoming)
	if err != nil {
		log.Fatal("sub aave liq call", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal("sub aave chan", err)
		case vLog := <-incoming:
			fmt.Println("tx hash", vLog.Raw.TxHash.Hex())
			fbalance := new(big.Float)
			fbalance.SetString(vLog.RepayAmount.String())
			tokenAmt := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
			fmt.Printf("Liquidated amt:%.2f coin: %s, tx: %s\n", tokenAmt, "cCOMP", vLog.Raw.TxHash.Hex())
			// twitter.SendTweet("Liquidated", tokenAmt, "cCOMP", vLog.Raw.TxHash.Hex(), "Compound")
		}
	}
}
