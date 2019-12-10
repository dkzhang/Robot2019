package client

import (
	"testing"
	"time"
)

func TestLifterControl(t *testing.T) {
	LifterControl(1000 * 20)
	time.Sleep(time.Second * 3)
	LifterControl(-1000 * 22)
}

//docker run -p 50051:50071 --rm --device /dev/gpiomem:/dev/mem  -e theApp="./lifterControl/server/server.gitcode_go" registry.cn-beijing.aliyuncs.com/dkzhang/robot2019-app-gitcode_go-arm32v7

//docker run -p 50051:50071 -d --restart=always --device /dev/gpiomem:/dev/mem  -e theApp="./lifterControl/server/server.gitcode_go" registry.cn-beijing.aliyuncs.com/dkzhang/robot2019-app-gitcode_go-arm32v7
