package main
import (
        "fmt"
        pubnub "github.com/pubnub/go"
    )


func main(){
        config := pubnub.NewConfig()
        config.PublishKey = "pub-c-1d2e61a4-c6bf-423f-98a6-bddda16c9207"
        config.SubscribeKey = "sub-c-0186167a-3832-11ea-afe9-722fee0ed680"

        pn := pubnub.NewPubNub(config)

        msg := map[string]interface{}{
            "asset_id" : "Amagi123",
            "status" : "uploaded",
        }


        resPub, statusPub, errPub := pn.Publish().
                Channel("blip_sirius").
                Message(msg).
                Execute()
        fmt.Println(resPub, statusPub, errPub)

        res, status, err := pn.History().
                Channel("blip_sirius").
                Count(10).
                IncludeTimetoken(true).
                Execute()

        if res != nil {
                fmt.Println(" --- HISTORY: ")
                if res.Messages != nil {
                        for _, v := range res.Messages {
                                fmt.Println(fmt.Sprintf("Message: %s, Timetoken: %d", v.Message, v.Timetoken))
                        }
                } else {
                        fmt.Println(fmt.Sprintf("Messages null"))
                }
                fmt.Println(fmt.Sprintf("EndTimetoken: %d", res.EndTimetoken))
                fmt.Println(fmt.Sprintf("StartTimetoken: %d", res.StartTimetoken))
                fmt.Println("")
        } else {
                fmt.Println(fmt.Sprintf("StatusResponse: %s %e", status.Error, err))
        }
}


