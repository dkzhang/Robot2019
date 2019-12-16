package client

import (
	"testing"
)

func TestLifterControl(t *testing.T) {
	//LifterControl(1000 * 10)
	//time.Sleep(time.Second * 3)
	LifterControl(-1000 * 35)
}

//docker run -p 50051:50071 --rm --device /dev/gpiomem:/dev/mem  -e theApp="./lifterControl/server/server.go" registry.cn-beijing.aliyuncs.com/dkzhang/robot2019-app-go-arm32v7

//docker run -p 50051:50071 -d --restart=always --device /dev/gpiomem:/dev/mem  -e theApp="./lifterControl/server/server.go" registry.cn-beijing.aliyuncs.com/dkzhang/robot2019-app-go-arm32v7
