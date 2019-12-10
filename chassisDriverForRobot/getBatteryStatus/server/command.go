package server

func GenerateGetCommand() (name, cmd string) {
	name = "/api/get_power_status"
	cmd = name
	return cmd, name
}
