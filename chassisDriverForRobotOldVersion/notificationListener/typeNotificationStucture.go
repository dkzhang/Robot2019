package notificationListener

type Notification struct {
	Type string `json:"type"`
	Code string `json:"code"`

	Level       string `json:"level"`
	Description string `json:"description"`
	Data        Data   `json:"data"`

	OriginalMessage string
}

type Data struct {
	Target            string   `json:"target"`
	Distance          float64  `json:"distance"`
	Markers           []string `json:"markers"`
	Count             int      `json:"count"`
	DistanceTolerance float64  `json:"distance_tolerance"`

	MoveStatus string `json:"move_status"`
}
