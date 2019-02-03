package config

import (
	"os"
	"strings"
)

type argsConfig struct {
	projectNameFilter string
}

func newArgsConfig() argsConfig {
	return parseArguments()
}

func parseArguments() argsConfig {
	argsConfig := argsConfig{}
	args := os.Args[1:]
	argsConfig.projectNameFilter = strings.Join(args, "")
	return argsConfig
}
