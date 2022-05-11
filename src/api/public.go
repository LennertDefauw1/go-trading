package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func QueryPublicEndpoint(endPointName, inputParameters string) string {
	baseDomain := "https://api.kraken.com"
	publicPath := "/0/public/"
	apiEndpointFullURL := baseDomain + publicPath + endPointName + "?" + inputParameters

	resp, err := http.Get(apiEndpointFullURL)
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

	return string(body)
}
