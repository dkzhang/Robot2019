package robotMove

import (
	"encoding/json"
	"fmt"
)

func UnmarshalJSON(strJSON string) (sr StructReturn, err error) {
	err = json.Unmarshal([]byte(strJSON), &sr)
	if err != nil {
		return sr, fmt.Errorf("json.Unmarshal error: %v", err)
	} else {
		return sr, nil
	}
}
