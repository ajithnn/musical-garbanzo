package main

import (
	"fmt"
	pubnub "github.com/pubnub/go"
)

func main() {

	ch := make(chan struct{}, 1)
	config := pubnub.NewConfig()
	config.PublishKey = "pub-c-1d2e61a4-c6bf-423f-98a6-bddda16c9207"
	config.SubscribeKey = "sub-c-0186167a-3832-11ea-afe9-722fee0ed680"
	pn := pubnub.NewPubNub(config)
	listener := pubnub.NewListener()

	pn.AddListener(listener)
	go func() {
		for {
			select {
			case message := <-listener.Message:
				//Channel
				fmt.Println(message.Channel)
				//Subscription
				fmt.Println(message.Subscription)
				//Payload
				fmt.Println(message.Message)
				//Publisher ID
				fmt.Println(message.Publisher)
				//Timetoken
				fmt.Println(message.Timetoken)
			}
		}
	}()

	pn.Subscribe().
		Channels([]string{"blip_sirius"}). // subscribe to channels
		Execute()
	<-ch
}
