module github.com/syntelos/go-auth

go 1.20

require (
	github.com/google/uuid v1.6.0
	golang.org/x/oauth2 v0.16.0
)

require (
	cloud.google.com/go/compute v1.20.1 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.20.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

retract v1.3.1

replace github.com/syntelos/go-auth => github.com/syntelos/go-auth v0.0.0-20240125195815-241fb77ebcd7
