package socketCommunication

type CommandStruct struct {
	Name    string
	Command string
	UUID    string
}

type CommandFeedback struct {
	Name    string
	Command string
	Msg     string
}
