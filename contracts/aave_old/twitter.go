package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func sendTweet(amt *big.Float, coin string, tx string) {
	fmt.Println("tweet start")
	config := oauth1.NewConfig("UqrmoFCcV8cIYMjG7GKrNspJw", "HBFXoAAOScCko9XSc9Yk3O1WKI4TnrAOjr4FdkrlqJYP5zzUvv")
	token := oauth1.NewToken("1248497327125303297-", "G2nkE4ItUES0RnpSKxWtNA7kCcra8OW0cMzxYCLia9XZ2")
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Send a Tweet
	tweet, resp, err := client.Statuses.Update(fmt.Sprintf("Liquidated %.2f of %s on Aave https://etherscan.io/tx/%s", amt, coin, tx), nil)
	fmt.Println("send it!", tweet, resp)
	if err != nil {
		log.Fatal(err)
	}
}
