package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/firo-18/pnc/info"
)

func init() {
	Index = append(Index, &discordgo.ApplicationCommand{

		Name:        "info",
		Description: "Information databse for PNC.",
		Type:        discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "doll",
				Description: "Information on a specific Doll.",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:         "name",
						Description:  "Select a doll.",
						Type:         discordgo.ApplicationCommandOptionString,
						Required:     true,
						Autocomplete: true,
					},
				},
			},
			{
				Name:        "class",
				Description: "Information on a specific class.",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "name",
						Description: "Select a class.",
						Type:        discordgo.ApplicationCommandOptionString,
						Required:    true,
						Choices:     classesChoice(),
					},
					{
						Name:         "doll",
						Description:  "Select a doll.",
						Type:         discordgo.ApplicationCommandOptionString,
						Autocomplete: true,
					},
				},
			},
		},
	})
}

func classesChoice() []*discordgo.ApplicationCommandOptionChoice {
	choices := []*discordgo.ApplicationCommandOptionChoice{}

	for _, class := range info.Classes {
		choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
			Name:  class,
			Value: class,
		})
	}

	return choices
}
