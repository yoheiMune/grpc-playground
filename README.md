# gRPC Playground
This is a playground for gRPC Go which has several tutorials and sample implementations.

## Pre installed
```shell

# Install `protoc`
$ brew install protobuf

# Install protoc-gen-go which is a plugin for `protoc`
$ go get -u github.com/golang/protobuf/protoc-gen-go

# (May need) Add to $PATH
$ export $PATH=$GOHOME/bin:$PATH
```
* References
http://www.grpc.io/docs/quickstart/go.html


## Samples
* [helloworld](./helloworld) : gRPC Hello World.
* [basics](./basics) : gRPC Basic samples (simple / streaming).

