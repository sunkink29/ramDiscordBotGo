package commands

import (
	"strings"

	"github.com/pazuzu156/atlas"
)

// ListCurGuilds command.
type ListCurGuilds struct{ Command }

// InitListCurGuilds initializes the listcurguilds command.
func InitListCurGuilds() ListCurGuilds {
	return ListCurGuilds{Init(&CommandItem{
		Name:        "listcurguilds",
		Description: "Lists all the guilds the bot is in",
		Aliases:     []string{},
		Usage:       "listcurguilds ...",
		Parameters:  []Parameter{},
	})}
}

// Register registers and runs the listcurguilds command.
func (c ListCurGuilds) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {
		guilds, err := ctx.Atlas.GetGuilds(ctx.Context, nil)
		if err != nil {
			ctx.Atlas.Logger.Error(err)
			ctx.Message.Reply(ctx.Context, ctx.Atlas, "Error Retreving guilds")
			return
		}

		var guildNames []string
		for _, guild := range guilds {
			guildNames = append(guildNames, guild.Name)
		}
		message := strings.Join(guildNames, ", ")
		ctx.Message.Reply(ctx.Context, ctx.Atlas, message)
	}
	return c.CommandInterface
}
