# study-go
编译：
```bash
go build -ldflags "-X main.VERSION=0.0.1 -X main.ARCH=amd64" loadAndPushImage.go
GOOS=linux GOARCH=arm64 go build -ldflags "-X main.VERSION=0.0.1 -X main.ARCH=arm64" loadAndPushImage.go
```

递归遍历当前文件夹下的tar文件，执行`docker load` `docker tag` `docker rmi` `docker push` `docker save`

# gui

参考： https://github.com/lxn/walk

```bash
rsrc -manifest test.manifest –ico check.ico -o rsrc.syso

go build -ldflags="-s -w -H windowsgui"
```
生成exe文件

# statik

把静态资源打包进可执行文件

参考https://github.com/rakyll/statik/tree/master/example

# rpc 测试

添加rpc测试
```bash
protoc --stdrpc_out=. arith.proto
```
# grpc 测试

添加grpc

安装grpc
```bash
git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc
git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net
git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto
git clone https://github.com/golang/sys $GOPATH/src/golang.org/x/sys
git clone https://github.com/golang/net $GOPATH/src/golang.org/x/net
go install google.golang.org/grpc
```
```bash
 protoc --go_out=plugins=grpc:. helloworld.proto
// 我手动把生成的XXX_删除了
```

# 添加go聊天室

# 添加websocket聊天室

# cobra命令行工具

# redis

# mongodb

# HTTP/1.1 HTTP/2

# ants goroutines 协程池

# graphql

# docker-web-terminal