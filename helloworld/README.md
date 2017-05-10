# gRPC Hello World for GO
This is a super-easy sample for beginners.

## How to Launch
```shell
# gRPC Server
$ go run server/main.go

# gRPC Client
$ go run client/main.go
```

## How to convert
```shell
$ protoc --go_out=plugins=grpc:. hello.proto  
```


## References
[http://www.grpc.io/docs/quickstart/go.html](http://www.grpc.io/docs/quickstart/go.html)