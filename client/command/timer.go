package command

import "github.com/bwmarrin/discordgo"

func init() {
	Test = append(Test, &discordgo.ApplicationCommand{
		Name:        "timer",
		Description: "Display reset times.",
		Type:        discordgo.ChatApplicationCommand,
	})
}
