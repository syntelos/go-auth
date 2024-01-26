/*
 * Another OAuth for GOPL
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package auth

import (
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)
/*
 * The Google OAUTH2 Client Credentials is encoded and
 * employed distinctly, requiring a dance around the
 * packaging situation.
 */
func ConfigToClient(config *oauth2.Config) (client []byte) {
	var str string = fmt.Sprintf(`{
    "installed": {
        "client_id": "%s",
        "auth_uri": "%s",
        "token_uri": "%s",
        "client_secret": "%s",
        "redirect_uris": ["%s"]
    }
}`,config.ClientID,config.Endpoint.AuthURL,config.Endpoint.TokenURL,config.ClientSecret,config.RedirectURL)

	return []byte(str)
}

func ConfigFromClient(client []byte, scopes []string) (config *oauth2.Config) {
	switch len(scopes) {
	case 0:
		config, _ = google.ConfigFromJSON(client)
	case 1:
		config, _ = google.ConfigFromJSON(client,scopes[0])
	case 2:
		config, _ = google.ConfigFromJSON(client,scopes[0],scopes[1])
	case 3:
		config, _ = google.ConfigFromJSON(client,scopes[0],scopes[1],scopes[2])
	default:
		config, _ = google.ConfigFromJSON(client,scopes[0],scopes[1],scopes[2],scopes[3])
	}

	return config
}
