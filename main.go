package main

import (
	"bufio"
	"fmt"
	"os"
)

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type authTokens struct {
	token      string
	renewToken string
}

type pubKey struct {
	pubKey string
}

type pubKeyStatus struct {
	expiresAt string
	pubKey    string
	id        string
}

func getLoginDetailsFromEnv() login {
	username := os.Getenv("SURFSHARK_USERNAME")
	password := os.Getenv("SURFSHARK_PASSWORD")

	loginDetails := login{Username: username, Password: password}

	return loginDetails

}

func createPubKey(key string) pubKey {
	// TODO: Check for valid pubKey and return err

	pKey := pubKey{pubKey: key}

	return pKey

}

func readConfigFile(filePath string) []pubKey {
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileScanner := bufio.NewScanner((readFile))
	fileScanner.Split(bufio.ScanLines)
	pubKeys := make([]pubKey, 0)

	for fileScanner.Scan() {
		pubKeyText := fileScanner.Text()
		pKey := createPubKey(pubKeyText)
		pubKeys = append(pubKeys, pKey)
	}

	readFile.Close()

	return pubKeys

}

func main() {
	fmt.Println(getLoginDetailsFromEnv())
	configs := readConfigFile(".config")

	fmt.Println(configs)
}
