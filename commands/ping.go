package commands

import "github.com/pazuzu156/atlas"

// Ping command.
type Ping struct{ Command }

// InitPing initializes the ping command.
func InitPing() Ping {
	return Ping{Init(&CommandItem{
		Name:        "ping",
		Description: "Ping/pong command",
		Aliases:     []string{},
		Usage:       "ping",
		Parameters:  []Parameter{},
	})}
}

// Register registers and runs the ping command.
func (c Ping) Register() *atlas.Command {
	c.CommandInterface.Run = func(c atlas.Context) {
		c.Message.Reply(c.Context, c.Atlas, "Pong!")
	}
	return c.CommandInterface
}
