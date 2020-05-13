package commands

import (
	"strings"

	"github.com/pazuzu156/atlas"
)

// Parrot command.
type Parrot struct{ Command }

// InitParrot initializes the parrot command.
func InitParrot() Parrot {
	return Parrot{Init(&CommandItem{
		Name:        "parrot",
		Description: "Say the same message back",
		Aliases:     []string{},
		Usage:       "parrot ...",
		Parameters:  []Parameter{},
	})}
}

// Register registers and runs the parrot command.
func (c Parrot) Register() *atlas.Command {
	c.CommandInterface.Run = func(c atlas.Context) {
		msgParts := strings.Split(c.Message.Content, " ")
		if len(msgParts) < 1 {
			msgParts = append(msgParts, "")
		}
		message := strings.Join(msgParts[1:], " ")
		c.Message.Reply(c.Context, c.Atlas, message)
	}
	return c.CommandInterface
}
