package main

import (
	"fmt"
	"plugin"

	"go.353solutions.com/config"
)

func main() {
	plug, _ := plugin.Open("config.so")
	v, _ := plug.Lookup("Config")
	cfg := v.(*config.Config)
	fmt.Printf("%+v\n", cfg)
	// &{MaxWorkers:1000 Port:9091 AuthServers:[auth1.local auth2.local]}
}
