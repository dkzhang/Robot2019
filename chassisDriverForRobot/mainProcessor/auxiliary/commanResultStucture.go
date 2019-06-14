package auxiliary

type ResultType struct {
	Type string `json:"type"`
}

type BasicCommandResult struct {
	Type         string `json:"type"`
	Command      string `json:"command"`
	ErrorMessage string `json:"error_message"`
	Status       string `json:"status"`
	UUID         string `json:"uuid"`
}

type CommandResult struct {
	basicInfo BasicCommandResult
	strJSON   string
}
