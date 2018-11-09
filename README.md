# Hands-On-Dependency-Injection-in-Go
Hands-On Dependency Injection in Go, published by Packt

## Quick Install/Setup

1. Install Go
1. Download this repo `go get github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/...`
1. Change to the code directory `cd $GOPATH/src/github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/`

### Running the sample service
1. Create a config file for the sample service by copying ./ch04/default-config.json to ./ch04/config.json 
`cp ./ch04/default-config.json ./ch04/config.json`
1. Customize the config -
1. Run the sample service (for a particular chapter) `ACME_CONFIG=$GOPATH/src/github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch04/config.json go run ./ch04/acme/main.go` 

### Running tests for the sample service
1. Create a config file for the sample service by copying ./ch04/default-config.json to ./ch04/config.json 
`cp ./ch04/default-config.json ./ch04/config.json`
1. Customize the config -
1. Run the sample service (for a particular chapter) `ACME_CONFIG=$GOPATH/src/github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch04/config.json go test ./ch04/...` 
