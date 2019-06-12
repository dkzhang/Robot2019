package robotStatus

import (
	"encoding/json"
	"fmt"
)

//JSON解析无抢救版，如果JSON不完整或有错误，则直接报错，不抢救其中可用信息片段
func UnmarshalJSON(strJSON string) (ct CallbackTopic, err error) {
	err = json.Unmarshal([]byte(strJSON), &ct)
	if err != nil {
		return ct, fmt.Errorf("json.Unmarshal error: %v", err)
	} else {
		ct.OriginalMessage = strJSON
		return ct, nil
	}
}
