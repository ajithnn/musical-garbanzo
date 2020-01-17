package main

import (
	"fmt"
	pubnub "github.com/pubnub/go"
)

func main() {
	config := pubnub.NewConfig()
	config.PublishKey = "pub-c-1d2e61a4-c6bf-423f-98a6-bddda16c9207"
	config.SubscribeKey = "sub-c-0186167a-3832-11ea-afe9-722fee0ed680"

	pn := pubnub.NewPubNub(config)

	msg := map[string]interface{}{
		"asset_id": "Amagi123",
		"status":   "uploaded",
	}

	resPub, statusPub, errPub := pn.Publish().
		Channel("blip_sirius").
		Message(msg).
		Execute()
	fmt.Println(resPub, statusPub, errPub)
}
