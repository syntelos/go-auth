/*
 * Another OAuth for GOPL
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package auth

import (
	"context"
	"io/ioutil"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"github.com/syntelos/go-auth/util"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

/*
 * Fetch OAuth Token via credentials cache.
 */
func Token(scopes ...string) (token *oauth2.Token) {

	scopes = ReviewScopes(scopes)

	var client []byte
	var er error

	client, er = ioutil.ReadFile(ClientFile)
	if nil == er {

		var config *oauth2.Config
		config, er = google.ConfigFromJSON(client,scopes[0])
		if nil == er {
			var redirect string = config.RedirectURL

			var consentPageSettings util.ConsentPageSettings = util.ConsentPageSettings {
				DisableAutoOpenConsentPage: false,
				InteractionTimeout: (time.Duration(2) * time.Minute),
			}

			var authCodeServer util.AuthorizationCodeServer = &util.AuthorizationCodeLocalhost {
				ConsentPageSettings: consentPageSettings,
				AuthCodeReqStatus: util.AuthorizationCodeStatus{Status: util.WAITING, Details: "Authorization code not yet set."},				
			}

			redirect, er = authCodeServer.ListenAndServe(redirect)
			if nil == er {
				defer authCodeServer.Close()

				var src oauth2.TokenSource
				var tok *oauth2.Token

				var params google.CredentialsParams = google.CredentialsParams {
					Scopes: scopes,
					State: "state",
					AuthHandler: util.Get3LOAuthorizationHandler("state", consentPageSettings, &authCodeServer),
					PKCE: util.GeneratePKCEParams() }

				var creds *google.Credentials

				creds, er = google.CredentialsFromJSONWithParams(context.Background(), client, params)

				if er == nil {

					src = creds.TokenSource

					ts := oauth2.ReuseTokenSource(nil, src)

					tok, er = ts.Token()

					if nil == er {
						return tok
					}
				}
			}
		}
	}
	return nil
}

const ScopePrefix string = "https://www.googleapis.com/auth/"
const ScopeInfixUser string = "userinfo."

func ReviewScopes(input []string) (output []string) {
	if 0 == len(input) {
		var scope string = ScopePrefix+"drive"
		output = append(output,scope)
	} else {
		for _, ins := range input {
			switch ins {
			case "openId":
				output = append(output,ins)

			case "user", "email":
				var scope string = ScopePrefix+ScopeInfixUser+ins
				output = append(output,scope)
				
			default:
				if -1 < strings.IndexByte(ins,'/') {
					output = append(output,ins)
				} else {
					var scope string = ScopePrefix+ins
					output = append(output,scope)
				}
			}
		}
	}
	return output
}

const ClientDirectoryName = ".gdr"
var ClientDirectory string = filepath.Join(GuessUnixHomeDir(), ClientDirectoryName)
var ClientFile string = filepath.Join(ClientDirectory, "client.json")

func GuessUnixHomeDir() string {
	// Prefer $HOME over user.Current due to glibc bug: golang.org/issue/13470
	if v := os.Getenv("HOME"); v != "" {
		return v
	}
	// Else, fall back to user.Current:
	if u, err := user.Current(); err == nil {
		return u.HomeDir
	}
	return ""
}
