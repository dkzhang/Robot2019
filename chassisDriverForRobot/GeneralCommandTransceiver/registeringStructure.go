package GeneralCommandTransceiver

type RegisteringStructure struct {
	command  string
	filter   func(string) bool
	callback func(string) error
}

func DefaultCommandFilter(uuid string) func(string) error {
	return func(s string) error {

		return nil
	}
}
