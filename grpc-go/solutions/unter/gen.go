package unter

/* Get the Go gRPC tools by running

$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

See more at https://grpc.io/docs/languages/go/quickstart/
*/

//go:generate mkdir -p pb
//go:generate protoc --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative unter.proto
