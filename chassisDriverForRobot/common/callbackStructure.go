package common

import (
	"encoding/json"
	"fmt"
)

type CallbackTopic struct {
	Type  string `json:"type"`
	Topic string `json:"topic"`
}

func CallbackTopicDetection(strJSON string, topic string) (pct *CallbackTopic, err error) {
	//先解析消息类型
	rt := ResultType{}
	err = json.Unmarshal([]byte(strJSON), &rt)
	if err != nil {
		return nil, fmt.Errorf("result type json unmarshal error: %v", err)
	}

	//判断消息类型是否为callback
	if rt.Type != "callback" {
		return nil, nil
	}

	//按callback格式进行详细解析
	ct := CallbackTopic{}
	err = json.Unmarshal([]byte(strJSON), &ct)
	if err != nil {
		return nil, fmt.Errorf("CommandResponse json unmarshal error: %v", err)
	}

	//判断topic是否相符
	if ct.Topic != topic {
		return nil, nil
	}

	return &ct, nil
}
