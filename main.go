package main

import (
	"fmt"

	"github.com/LennertDefauw1/go-trading/src/config"
	"github.com/LennertDefauw1/go-trading/src/socket"
)

func main() {

	fmt.Println("Hello, World!")

	config, _ := config.LoadConfig("./config")

	fmt.Println(config.PrivateApiKey)

	publicWebSocketURL := "wss://ws.kraken.com/"
	publicWebSocketSubscriptionMsg := "{ \"event\":\"subscribe\", \"subscription\":{\"name\":\"trade\"},\"pair\":[\"DOT/USD\"] }"

	/*
	   *MORE PUBLIC WEBSOCKET EXAMPLES

	   publicWebSocketSubscriptionMsg := "{ \"event\": \"subscribe\", \"subscription\": { \"interval\": 1440, \"name\": \"ohlc\"}, \"pair\": [ \"XBT/EUR\"]}"
	   publicWebSocketSubscriptionMsg := "{ \"event\": \"subscribe\", \"subscription\": { \"name\": \"spread\"}, \"pair\": [ \"XBT/EUR\",\"ETH/USD\" ]}"
	*/

	socket.OpenAndStreamWebSocketSubscription(publicWebSocketURL, publicWebSocketSubscriptionMsg)
}
