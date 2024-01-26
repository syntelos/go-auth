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
	"time"
)

const ScopePrefix string = "https://www.googleapis.com/auth/"
const DriveScope string = ScopePrefix + "drive"


/*
 * Fetch OAuth Token via credentials cache.
 */
func Token() (token *oauth2.Token) {
	var json []byte
	var er error

	json, er = ioutil.ReadFile(util.CacheFile)
	if nil == er {

		var config *oauth2.Config
		config, er = google.ConfigFromJSON(json,DriveScope)
		if nil == er {
			var redirect string = config.RedirectURL

			var consentPageSettings util.ConsentPageSettings = util.ConsentPageSettings{DisableAutoOpenConsentPage: false, InteractionTimeout: (time.Duration(2) * time.Minute)}

			var authCodeServer util.AuthorizationCodeServer = &util.AuthorizationCodeLocalhost {
				ConsentPageSettings: consentPageSettings,
				AuthCodeReqStatus: util.AuthorizationCodeStatus{Status: util.WAITING, Details: "Authorization code not yet set."},				
			}
			authCodeServer.ListenAndServe(redirect)

			defer authCodeServer.Close()

			var settings util.Settings = util.Settings{
				CredentialsJSON: string(json),
				Scope:           DriveScope,
				AuthHandler:     util.Get3LOAuthorizationHandler("state", consentPageSettings, &authCodeServer),
				State:           "state",
				Audience:        "",
				QuotaProject:    "",
				Sts:             false,
				ServiceAccount:  "",
				Email:           "",
				AuthType:        util.AuthTypeOAuth,
			}
			var taskSettings util.TaskSettings = util.TaskSettings{
				AuthType:  "oauth",
				Format:    "bare",
				CurlCli:   "",
				Url:       "",
				ExtraArgs: []string{"drive"},
				SsoCli:    "",
				Refresh:   false,
			}

			return util.FetchToken(&settings,&taskSettings)
		}
	}
	return nil
}
