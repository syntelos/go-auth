/*
 * Another OAuth for GOPL
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package main

import (
	auth "github.com/syntelos/go-auth"
	"fmt"
	oauth2 "golang.org/x/oauth2"
	"os"
)

/*
 * Fetch OAuth Token.
 */
func main() {

	var token *oauth2.Token = auth.Token(os.Args[1:])

	if nil != token && token.Valid() {

		fmt.Println(token.AccessToken)

		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
