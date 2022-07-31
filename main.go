package main

import (
	"fmt"
	"os"
)

type login struct {
	username string
	password string
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

	loginDetails := login{username: username, password: password}

	return loginDetails

}

func main() {
	fmt.Println(getLoginDetailsFromEnv())
}
