package commands

import (
	"strings"

	"github.com/pazuzu156/atlas"
)

// ListCurRoles command.
type ListCurRoles struct{ Command }

// InitListCurRoles initializes the listcurroles command.
func InitListCurRoles() ListCurRoles {
	return ListCurRoles{Init(&CommandItem{
		Name:        "listcurroles",
		Description: "List all the roles in the current server",
		Aliases:     []string{},
		Usage:       "listcurroles ...",
		Parameters:  []Parameter{},
	})}
}

// Register registers and runs the listcurroles command.
func (c ListCurRoles) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {
		roles, err := ctx.Atlas.GetGuildRoles(ctx.Context, ctx.Message.GuildID)
		if err != nil {
			ctx.Atlas.Logger.Error(err)
			ctx.Message.Reply(ctx.Context, ctx.Atlas, "Error Retreving roles")
		}

		var roleNames []string
		for _, role := range roles {
			roleNames = append(roleNames, role.Name)
		}
		message := strings.Join(roleNames, ", ")
		ctx.Message.Reply(ctx.Context, ctx.Atlas, message)
	}
	return c.CommandInterface
}
