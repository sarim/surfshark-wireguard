package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const ApplicationJsonRequestType = "application/json"

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type authTokens struct {
	Token      string `json:"token"`
	RenewToken string `json:"renewToken"`
}

type pubKey struct {
	PubKey string `json:"pubKey"`
}

type pubKeyStatus struct {
	ExpiresAt string `json:"expiresAt"`
	PubKey    string `json:"pubKey"`
	Id        string `json:"id"`
}

func getLoginDetailsFromEnv() login {
	username := os.Getenv("SURFSHARK_USERNAME")
	password := os.Getenv("SURFSHARK_PASSWORD")

	loginDetails := login{Username: username, Password: password}

	return loginDetails

}

func createPubKey(key string) pubKey {
	// TODO: Check for valid pubKey and return err

	pKey := pubKey{PubKey: key}

	return pKey

}

func readConfigFile(filePath string) []pubKey {
	readFile, err := os.Open(filePath)

	if err != nil {
		log.Panic(err)
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

func authenticate(loginData login) authTokens {
	loginURL := "https://api.surfshark.com/v1/auth/login"

	jsonLoginData, err := json.Marshal(loginData)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(loginURL, ApplicationJsonRequestType, bytes.NewBuffer(jsonLoginData))

	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unknown Status: %q", resp.Status)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var tokens authTokens

	json.Unmarshal(body, &tokens)

	return tokens

}

func registerPubKey(tokens authTokens, pubKey pubKey) {
	pubKeyCreationURL := "https://api.surfshark.com/v1/account/users/public-keys"
	bearer := "Bearer " + tokens.Token

	jsonPubKey, _ := json.Marshal(pubKey)

	req, _ := http.NewRequest("POST", pubKeyCreationURL, bytes.NewBuffer(jsonPubKey))

	req.Header.Add("Authorization", bearer)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln("Error syncing pubkey:", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var keyStatus pubKeyStatus

	json.Unmarshal(body, &keyStatus)

	log.Println("Register: ", keyStatus)
}

func extendPubKeyValidity(tokens authTokens, pubKey pubKey) {
	pubKeyValidityExtendURL := "https://api.surfshark.com/v1/account/users/public-keys/validate"
	bearer := "Bearer " + tokens.Token

	jsonPubKey, _ := json.Marshal(pubKey)

	req, _ := http.NewRequest("POST", pubKeyValidityExtendURL, bytes.NewBuffer(jsonPubKey))

	req.Header.Add("Authorization", bearer)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln("Error syncing pubkey:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		registerPubKey(tokens, pubKey)
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var keyStatus pubKeyStatus

	json.Unmarshal(body, &keyStatus)

	log.Println("Extend: ", keyStatus)
}

func processKeys(token authTokens, pubKeys []pubKey) {
	for _, key := range pubKeys {
		extendPubKeyValidity(token, key)
	}
}

func main() {
	loginData := getLoginDetailsFromEnv()
	tokens := authenticate(loginData)

	configFilePath := os.Getenv("CONFIG_FILE_PATH")
	if configFilePath == "" {
		configFilePath = ".config"
	}
	pubKeys := readConfigFile(configFilePath)

	processKeys(tokens, pubKeys)

}
