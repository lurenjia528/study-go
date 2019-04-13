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