package socketCommunication

type CommandStruct struct {
	Name       string
	Command    string
	ChanResult chan CommandResult
}

type CommandFeedback struct {
	Name    string
	Command string
	Msg     string
}
