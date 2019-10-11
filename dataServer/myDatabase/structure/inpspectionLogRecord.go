package structure

import (
	"Robot2019/myUtil"
	"encoding/json"
	"strconv"
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

type InspectionLogRecord_string struct {
	InspectionID string
	RecordID     string
	Level        string
	DateTime     string
	TextContent  string
	ImageUrl     string
}

func (ilr *InspectionLogRecord) ConvertTo_string() InspectionLogRecord_string {
	return InspectionLogRecord_string{
		InspectionID: strconv.FormatInt(ilr.InspectionID, 10),
		RecordID:     strconv.FormatInt(ilr.RecordID, 10),
		Level:        ilr.Level,
		DateTime:     myUtil.FormatTime(ilr.DateTime),
		TextContent:  ilr.TextContent,
		ImageUrl:     ilr.ImageUrl,
	}
}

//strconv.FormatFloat(v, 'g', -1, 64)

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
