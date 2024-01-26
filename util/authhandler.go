// Copyright 2020 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Contains authorization handler functions.
package util

import (
	"errors"
	"fmt"
	"golang.org/x/oauth2/authhandler"
	"log"
	"os/exec"
	"runtime"
	"time"
)

/*
 * 3LO authorization handler.
 *
 * Note that the "state" parameter is used to prevent CSRF attacks.
 */
func Get3LOAuthorizationHandler(state string, authCodeServer *AuthorizationCodeServer) authhandler.AuthorizationHandler {

	return func(authCodeURL string) (string, string, error) {

		const (
			maxWaitForListenAndServe time.Duration = 10 * time.Second
		)

		if started, _ := (*authCodeServer).WaitForListeningAndServing(maxWaitForListenAndServe); started {

			if ber := OpenURL(authCodeURL); ber != nil {

				log.Fatal(ber)
			}

			(*authCodeServer).WaitForConsentPageToReturnControl()
		}

		code, err := (*authCodeServer).GetAuthenticationCode()

		return code.Code, code.State, err
	}
}

func OpenURL(url string) error {
	var err error

	switch runtime.GOOS {
	case "darwin":
		err = exec.Command("open", url).Start()
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	default:
		err = errors.New("Unsupported runtime")
	}

	if err != nil {

		return fmt.Errorf("Unable to open browser window (%s): %v", runtime.GOOS, err)
	} else {
		return nil
	}
}
