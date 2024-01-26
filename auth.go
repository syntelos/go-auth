/*
 * Another OAuth for GOPL
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package auth

import (
	"io/ioutil"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"github.com/syntelos/go-auth/util"
	"strings"
	"time"
)

/*
 * Fetch OAuth Token via credentials cache.
 */
func Token(scopes ...string) (token *oauth2.Token) {

	scopes = ReviewScopes(scopes)

	var json []byte
	var er error

	json, er = ioutil.ReadFile(util.ClientFile)
	if nil == er {

		var config *oauth2.Config
		config, er = google.ConfigFromJSON(json,scopes[0])
		if nil == er {
			var redirect string = config.RedirectURL

			var consentPageSettings util.ConsentPageSettings = util.ConsentPageSettings{DisableAutoOpenConsentPage: false, InteractionTimeout: (time.Duration(2) * time.Minute)}

			var authCodeServer util.AuthorizationCodeServer = &util.AuthorizationCodeLocalhost {
				ConsentPageSettings: consentPageSettings,
				AuthCodeReqStatus: util.AuthorizationCodeStatus{Status: util.WAITING, Details: "Authorization code not yet set."},				
			}

			redirect, er = authCodeServer.ListenAndServe(redirect)
			if nil == er {
				defer authCodeServer.Close()

				var settings util.Settings = util.Settings{
					CredentialsJSON: string(json),
					Scope:           scopes[0],
					AuthHandler:     util.Get3LOAuthorizationHandler("state", consentPageSettings, &authCodeServer),
					State:           "state",
					Sts:             false,
					AuthType:        util.AuthTypeOAuth,
				}
				var taskSettings util.TaskSettings = util.TaskSettings{
					AuthType:  "oauth",
					Format:    "bare",
					ExtraArgs: scopes,
					Refresh:   false,
				}

				return util.FetchToken(&settings,&taskSettings)
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
