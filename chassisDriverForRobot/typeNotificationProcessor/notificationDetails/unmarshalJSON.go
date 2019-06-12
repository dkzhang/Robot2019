package notificationDetails

import (
	"encoding/json"
	"fmt"
)

//JSON解析无抢救版，如果JSON不完整或有错误，则直接报错，不抢救其中可用信息片段
func UnmarshalJSON(strJSON string) (nd NotificationDetails, err error) {
	err = json.Unmarshal([]byte(strJSON), &nd)
	if err != nil {
		return nd, fmt.Errorf("json.Unmarshal error: %v", err)
	} else {
		nd.OriginalMessage = strJSON
		return nd, nil
	}
}
