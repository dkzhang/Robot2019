package socketCommunication

type CommandStruct struct {
	Name       string
	Command    string
	ChanResult chan CommandResultStruct
}

type CommandResultStruct struct {
	basicInfo BasicStructReturn
	strJSON   string
}

type CommandFeedback struct {
	Name    string
	Command string
	Msg     string
}
