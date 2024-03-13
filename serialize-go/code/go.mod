module serialize

go 1.18

replace go.353solutions.com/config => ./config

require (
	github.com/BurntSushi/toml v1.1.0
	github.com/golang/protobuf v1.5.2
	go.353solutions.com/config v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v2 v2.4.0
)

require google.golang.org/protobuf v1.33.0 // indirect
