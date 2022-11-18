package command

import "github.com/bwmarrin/discordgo"

func init() {
	Test = append(Test, &discordgo.ApplicationCommand{
		Name:        "stats",
		Description: "View various bot stats.",
		Type:        discordgo.ChatApplicationCommand,
	})
}
