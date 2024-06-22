package main

import (
	"fmt"
	"time"
)

/*
go build -o testgo main.go
nohup ./testgo > log/a.log 2>&1 &
docker  rm promtail

docker run  --name promtail \
-v /webapp/xxx/hot_douyin_api/log/config.yml:/etc/promtail/config.yml \
-v /webapp/xxx/hot_douyin_api/log:/var/log \
grafana/promtail:3.0.0

*/
func generateLog() string {
	// 这里可以替换为你需要生成的日志信息的逻辑
	return fmt.Sprintf("Generated xxxx test log entry at %s", time.Now().Format("2006-01-02 15:04:05"))
}

func main() {
	for {
		logEntry := generateLog()
		fmt.Println(logEntry)
		time.Sleep(500 * time.Millisecond)
	}
}
