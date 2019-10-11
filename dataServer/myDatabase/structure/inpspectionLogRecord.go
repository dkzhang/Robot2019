package structure

import (
	"encoding/json"
	"time"
)

type InspectionLogRecord struct {
	InspectionID int64
	RecordID     int64
	Level        string
	DateTime     time.Time
	TextContent  string
	ImageUrl     string
}

func MyMakeJSON() (string, error) {
	ilr := InspectionLogRecord{
		InspectionID: 1,
		RecordID:     2,
		Level:        "Normal",
		DateTime:     time.Now(),
		TextContent:  "xxx",
		ImageUrl:     "http://xxx.yyy.jpg",
	}

	b, err := json.Marshal(ilr)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
