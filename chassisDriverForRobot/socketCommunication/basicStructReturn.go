package socketCommunication

type BasicStructReturn struct {
	Type         string `json:"type"`
	Command      string `json:"command"`
	ErrorMessage string `json:"error_message"`
	Status       string `json:"status"`
	UUID         string `json:"uuid"`
}
