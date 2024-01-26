Another OAuth for GOPL


  func Token(scopes ...string) (token *oauth2.Token) 

    OAUTH TOKEN application service requires downloaded
    [CLIENT] in
 
      ~/.goauth/client.json


Command line interface

  Run "make" to build a CLI to print the access token, as
  for caching.


References

  [UTIL] https://github.com/google/oauth2l

    A copy of OAUTH2L/UTIL has been truncated to its
    necessary and essential employment.  Note that GOOG
    OAUTH2L UTIL has an APACHE license.

  [CLIENT] https://console.cloud.google.com/apis/credentials

