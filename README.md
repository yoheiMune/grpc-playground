# grpc-playground
gRPC Playground for Go lang

## Pre install
```shell
$ brew install protobuf

$ go get -u github.com/golang/protobuf/protoc-gen-go

~/go/bin # パスを通す.
$ sudo cp ~/go/bin/protoc-gen-go /usr/local/go/bin/

$ protoc --go_out=plugins=grpc:. my.proto 
# => my.pb.go

google.golang.org/grpc

```



## 参考
[http://blog.fenrir-inc.com/jp/2016/10/grpc-go.html](http://blog.fenrir-inc.com/jp/2016/10/grpc-go.html)

