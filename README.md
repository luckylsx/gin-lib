### run container:

docker run -d -p 9001:9001 --env APP_ENV=dev --name go-docker -v /opt/wwwroot/tob/web/gin-lib:/opt/wwwroot/go/gin-lib docker-go:latest

**运行**

docker exec -it go-docker /bin/bash /opt/wwwroot/go/gin-lib/run.sh

**运行main 文件 默认启动http 服务**

go run main.go

**默认运行http服务**

go run main.go

**运行 crontab 服务**

go run main.go -isCron=true

**查看包文档**

godoc -http :8888

访问: http://localhost:8888/pkg/
gin-lib 下 为项目包文档




