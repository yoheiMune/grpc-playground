# gRPC のサンプル実装

## 起動方法
```shell
# gRPC サーバー
$ go run server/main.go

# gRPC クライアント
$ go run client/main.go
```

## proto の変換方法
```shell
$ protoc --go_out=plugins=grpc:. hello.proto  
```


## 参考URL
[http://www.grpc.io/docs/quickstart/go.html](http://www.grpc.io/docs/quickstart/go.html)