package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/LennertDefauw1/go-trading/src/auth"
)

func QueryPrivateEndpoint(endPointName,
	inputParameters,
	apiPublicKey,
	apiPrivateKey string) string {
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

	jsonData := string(body)
	return jsonData
}
