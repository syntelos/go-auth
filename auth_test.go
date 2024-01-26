/*
 * Another OAuth for GOPL
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package auth

import (
	"fmt"
	"golang.org/x/oauth2"
	"testing"
)

func TestAuthToken(t *testing.T){

	var token *oauth2.Token = Token()

	if nil != token && token.Valid() {

		fmt.Println("[TestAuthToken] Success")

	} else {
		t.Fatal("[TestAuthToken] Failure")
	}
}
