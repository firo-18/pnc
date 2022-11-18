package command

import "github.com/bwmarrin/discordgo"

func init() {
	Test = append(Test, &discordgo.ApplicationCommand{
		Name:        "set",
		Description: "Set various bot parameters.",
		Type:        discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "server-time",
				Description: "Set server reset time if the default time is no longer correct.",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "hour",
						Description: "Enter reset hour in UTC time. For example, if reset is at 1 PM UTC time, enter 13.",
						Type:        discordgo.ApplicationCommandOptionInteger,
						Required:    true,
					},
				},
			},
		},
	})
}
