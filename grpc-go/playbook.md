- gRPC overview
- copy git URL to zoom
    - https://github.com/tebeka/talks/tree/master/grpc-go
- The Protocol buffers definition file
    - solutions/unter/unter.proto
    - https://developers.google.com/protocol-buffers/docs/proto3#scalar
    - start with location
    - `protoc --go_out=pb --go_opt=paths=source_relative unter.proto`
        - gen.go
    - go over sources
    - time
    - status (enum)
- Serializing data
    - solutions/unter/cmd/marshal/main.go
	- start with empty
	- xxd output file
	- add name, then plate
	- location as pointer
	- Updated: start with time.Time
	- jsonproto

- Writing a server
    - solutions/unter/cmd/server/main.go
    - Start with just Update
    - grpcurl --plaintext localhost:7689 list Unter
    - grpcurl --plaintext localhost:7689 describe Unter.Update
    - grpcurl --plaintext localhost:7689 describe .UpdateRequest
    - grpcurl --plaintext -d @ localhost:7689  Unter.Update < car.json
    - validate
- Writing a client
    - solutions/unter/cmd/client/main.go
- Error handling
    - validate location
- exercise: near

---

# Session 3: Beyond Request/Response

- Timeouts & cancellations
    - ctx
    - https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-metadata.md
- Streaming data
    - client stream updates
	- Update
- Writing interceptors (middleware)
    - timingInterceptor in server
    - https://github.com/grpc-ecosystem/go-grpc-middleware
- exercise: id key in metadata, log it in ctx all over
    - context.WithValue

# Session 4: Development & Deployment

- Sharing your protobuf definitions
    - discussion
	- monorepo
	- git submodule
	- pb in repo or different repo
- Testing your code
    - Regular test
- Security
    - solutions/carz/gen-cert.sh
    - JWT? Just API key
- Using grpc-gateway to generate REST API
    - https://grpc-ecosystem.github.io/grpc-gateway/
    - solutions/carz/gen-gateway.sh
    - solutions/carz/cmd/gateway
    - curl -d@car.json http://localhost:8080/Carz/Set
