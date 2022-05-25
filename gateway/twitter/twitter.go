package twitter

import (
	"fmt"
	"log"
	"math/big"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func SendTweet(event string, amt *big.Float, coin string, tx string, protocol string) {
	fmt.Println("tweet start")
	config := oauth1.NewConfig("UqrmoFCcV8cIYMjG7GKrNspJ", "HBFXoAAOScCko9XSc9Yk3O1WKI4TnrAOjr4FdkrlqJYP5zzUv")
	token := oauth1.NewToken("1248497327125303297-9Pgqq4pWC3anVhPAB1WtW7IEfirWA", "G2nkE4ItUES0RnpSKxWtNA7kCcra8OW0cMzxYCLia9XZ")
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Send a Tweet
	_, _, err := client.Statuses.Update(fmt.Sprintf("%s %.2f of %s on %s https://etherscan.io/tx/%s", event, amt, coin, protocol, tx), nil)
	if err != nil {
		log.Fatal(err)
	}
}
