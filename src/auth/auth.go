package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
)

func CreateAuthenticationSignature(apiPrivateKey,
	apiPath,
	endPointName,
	nonce,
	apiPostBodyData string) string {

	apiPost := nonce + apiPostBodyData
	secret, _ := base64.StdEncoding.DecodeString(apiPrivateKey)
	apiEndpointPath := apiPath + endPointName
	sha := sha256.New()
	sha.Write([]byte(apiPost))
	shaSum := sha.Sum(nil)
	allBytes := append([]byte(apiEndpointPath), shaSum...)
	mac := hmac.New(sha512.New, secret)
	mac.Write(allBytes)
	macSum := mac.Sum(nil)
	signatureString := base64.StdEncoding.EncodeToString(macSum)
	return signatureString
}
