module serialize

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	go.353solutions.com/config v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v2 v2.2.7
)

replace go.353solutions.com/config => ./config
