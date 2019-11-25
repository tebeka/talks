package config

// Config holds configuration fields
type Config struct {
	MaxWorkers  int
	Port        int
	AuthServers []string
}