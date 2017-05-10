# gRPC Gateway Sample
This is a sample of gRPC Gateway.

## Pre installed
It assumes that you have read the [Project Root README](../) and followed the instruction.  

Install another plugins for `protoc`.
```shell
$ go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
$ go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
```

## How to convert
### Stub
```shell
$ protoc -I. \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --go_out=plugins=grpc:. \
    *.proto
```
This will generate a stub `profile.pb.go`.
### Gateway
```shell
$ protoc -I. \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --grpc-gateway_out=logtostderr=true:. \
    *.proto
```
This will generate a stub `profile.pb.gw.go`.

## Implementation
### gRPC Server
* [server/main.go](server/main.go)
* [gateway/main.go](gateway/main.go)

## Test
```shell
# Launch.
$ go run server/main.go
$ go run gateway/main.go

# Test.
$ curl -d '{"value":"aaa"}' http://localhost:8080/v1/example/echo
{"value":"aaa"}
```

## References
https://github.com/grpc-ecosystem/grpc-gateway