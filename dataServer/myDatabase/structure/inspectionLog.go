package structure

import "time"

type InspectionLog struct {
	InspectionID         int64
	DateTimeBegin        time.Time
	DateTimeBeginEnd     time.Time
	InspectionConclusion string
}

func test() {
	time.Now()
}
