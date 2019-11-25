package main

import (
	"runtime"

	"go.353solutions.com/config"
)

var Config = config.Config{
	Port:        9091,
	AuthServers: []string{"auth1.local", "auth2.local"},
}

func init() {
	if runtime.GOOS == "windows" {
		Config.MaxWorkers = 100
	} else {
		Config.MaxWorkers = 1000
	}
}
