# study-go
编译：
```bash
go build -ldflags "-X main.VERSION=0.0.1 -X main.ARCH=amd64" loadAndPushImage.go
GOOS=linux GOARCH=arm64 go build -ldflags "-X main.VERSION=0.0.1 -X main.ARCH=arm64" loadAndPushImage.go
```

递归遍历当前文件夹下的tar文件，执行`docker load` `docker tag` `docker rmi` `docker push` `docker save`
