https://www.dockercn.net/?id=172

```powershell
$env:CGO_ENABLED=0
$env:GOOS='linux'
$env:GOARCH='amd64'
go build hello.go
```

https://www.youtube.com/watch?v=SnSH8Ht3MIc

```powershell
docker info
```

```powershell
docker build .
```

Dockerfile注意大小写

```ps1
[+] Building 0.5s (5/5) FINISHED
 => [internal] load build definition from Dockerfile                                                                                           0.1s 
 => => transferring dockerfile: 112B                                                                                                           0.0s 
 => [internal] load .dockerignore                                                                                                              0.0s 
 => => transferring context: 2B                                                                                                                0.0s 
 => [internal] load build context                                                                                                              0.1s 
 => => transferring context: 1.77MB                                                                                                            0.1s 
 => [1/1] COPY hello /                                                                                                                         0.0s 
 => exporting to image                                                                                                                         0.1s 
 => => exporting layers                                                                                                                        0.1s 
 => => writing image sha256:f5fdfc693fe31bc719619bc97ee83a10f940069c96f4cd6ed5a0abe960a22554                                                   0.0s
```

```ps1
docker build . -t docker-test
```

```ps1
> docker images
REPOSITORY    TAG               IMAGE ID       CREATED              SIZE
docker-test   latest            f5fdfc693fe3   About a minute ago   1.77MB
redis         latest            9da089657551   4 days ago           117MB
rabbitmq      3.10-management   6b94498c1b2f   3 weeks ago          262MB
```

```ps1
docker run docker-test
```

## 流程

- 创建Dockerfile
  - 其中语法再看
- `docker build . -t [REPOSITORY]` 创建镜像
- `docker run [REPOSITORY]` 运行镜像 

参见：
- https://www.dockercn.net/?id=172
- https://www.youtube.com/watch?v=SnSH8Ht3MIc

docker内部怎么连外面网啊？
host模式？

curl 命令转golang
https://mholt.github.io/curl-to-go/

普通的运行能够访问网络。

```ps1
docker images  # 查看
$env:CGO_ENABLED=0 
$env:GOOS='linux'  
$env:GOARCH='amd64'
go build -o ./bin/hello hello.go  

docker build . -t docker-test
docker run docker-test       
```

记得开启docker desktop

`ENTRYPOINT ["/hello"]`是啥。