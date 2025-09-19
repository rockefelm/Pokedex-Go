package structs


type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}