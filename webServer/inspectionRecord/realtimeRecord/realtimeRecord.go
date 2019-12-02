package realtimeRecord

type RealTimeInfo struct {
	InspectionID int
	RecordID     int
	Level        string
	DateTime     string
	TextContent  string
	ImageUrl     string
}

type RealTimeRecords struct {
	Records []RealTimeInfo
}
