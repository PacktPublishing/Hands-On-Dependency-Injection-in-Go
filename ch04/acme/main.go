package main

import (
	"flag"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch04/acme/internal/config"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch04/acme/internal/rest"
)

// config file location (parsed from the command line)
var configFile string

func main() {
	// parse flags
	parseFlags()

	// load config
	err := config.Load(configFile)
	if err != nil {
		panic(err.Error())
	}

	// start REST server
	server := rest.New()
	server.Listen()
}

func parseFlags() {
	flag.StringVar(&configFile, "config", "", "JSON config file location")
	flag.Parse()
}
