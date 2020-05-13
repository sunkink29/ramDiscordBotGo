package main

import (
	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/atlas"
	"github.com/sunkink29/ramDiscordBotGo/commands"
	"github.com/sunkink29/ramDiscordBotGo/lib"
)

var config = lib.Config()

func main() {
	client := atlas.New(&atlas.Options{
		DisgordOptions: disgord.Config{
			BotToken: config.Token,
			// Logger:   disgord.DefaultLogger(true), // uncomment for disgord logging
		},
	})

	client.Use(atlas.DefaultLogger())
	client.GetPrefix = func(m *disgord.Message) string {
		return config.Prefix
	}
	if err := client.Init(); err != nil {
		panic(err)
	}
}

func init() {
	atlas.Use(commands.InitPing().Register())
	atlas.Use(commands.InitParrot().Register())
	atlas.Use(commands.InitListCurGuilds().Register())
	atlas.Use(commands.InitListCurRoles().Register())
}
