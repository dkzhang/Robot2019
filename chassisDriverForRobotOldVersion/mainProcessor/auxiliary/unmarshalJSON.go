package auxiliary

import (
	"encoding/json"
	"fmt"
)

func UnmarshalJSON(strJSON string) (rt ResultType, err error) {
	err = json.Unmarshal([]byte(strJSON), &rt)
	if err != nil {
		return rt, fmt.Errorf("json.Unmarshal error: %v", err)
	} else {
		return rt, nil
	}
}
