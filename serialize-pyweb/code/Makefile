all:
	$(error please pick a target)

protobuf:
	protoc trade.proto --go_out=pb

plugin:
	go build -buildmode=plugin -o config.so config.go
