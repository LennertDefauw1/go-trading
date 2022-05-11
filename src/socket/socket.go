package socket

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/LennertDefauw1/go-trading/src/auth"
	"github.com/sacOO7/gowebsocket"
)

type WsTokenJsonStruct struct {
	Error  []interface{} `json:"error"`
	Result struct {
		Token   string `json:"token"`
		Expires int    `json:"expires"`
	} `json:"result"`
}

func OpenAndStreamWebSocketSubscription(connectionURL, webSocketSubscription string) {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	webSocketClient := gowebsocket.New(connectionURL)

	webSocketClient.OnConnectError = func(err error, socket gowebsocket.Socket) {
		fmt.Println("Received connect error - ", err)
	}

	webSocketClient.OnConnected = func(socket gowebsocket.Socket) {
		fmt.Println("Connected to server")
	}

	webSocketClient.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		fmt.Println(time.Now().Format("01-02-2006 15:04:05") + ": " + message)
	}

	webSocketClient.OnPingReceived = func(message string, socket gowebsocket.Socket) {
		fmt.Println("Received ping - " + message)
	}

	webSocketClient.OnPongReceived = func(message string, socket gowebsocket.Socket) {
		fmt.Println("Received pong - " + message)
	}

	webSocketClient.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		fmt.Println("Socket Closed")
	}

	webSocketClient.Connect()
	webSocketClient.SendText(webSocketSubscription)

	for {
		select {
		case <-interrupt:
			log.Println("interrupt")
			webSocketClient.Close()
			return
		}
	}
}

func GetWebSocketToken(endPointName,
	inputParameters,
	apiPublicKey,
	apiPrivateKey string) WsTokenJsonStruct {
	baseDomain := "https://api.kraken.com"
	privatePath := "/0/private/"

	apiEndpointFullURL := baseDomain + privatePath + endPointName + "?" + inputParameters
	nonce := fmt.Sprintf("%d", time.Now().Unix())
	apiPostBodyData := "nonce=" + nonce + "&" + inputParameters

	signature := auth.CreateAuthenticationSignature(apiPrivateKey,
		privatePath,
		endPointName,
		nonce,
		apiPostBodyData)

	httpOptions, err := http.NewRequest("POST", apiEndpointFullURL, strings.NewReader(apiPostBodyData))
	httpOptions.Header.Add("API-Key", apiPublicKey)
	httpOptions.Header.Add("API-Sign", signature)
	httpOptions.Header.Add("User-Agent", "GO Lang Client")

	if err != nil {
		fmt.Println("ERROR OCCURED: ", err)
		os.Exit(1)
	}

	resp, err := http.DefaultClient.Do(httpOptions)

	if err != nil {
		fmt.Println("ERROR OCCURED: ", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("ERROR OCCURED: ", err)
		os.Exit(1)
	}

	var jsonData WsTokenJsonStruct
	json.Unmarshal(body, &jsonData)

	return jsonData
}
