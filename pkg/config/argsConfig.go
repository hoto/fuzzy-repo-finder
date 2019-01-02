package config

import (
	"os"
	"strings"
)

type argsConfig struct {
	query string
}

func newArgsConfig() argsConfig {
	return parseArguments()
}

func parseArguments() argsConfig {
	argsConfig := argsConfig{}
	args := os.Args[1:]
	argsConfig.query = strings.Join(args, "")
	return argsConfig
}
