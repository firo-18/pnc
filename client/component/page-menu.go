package component

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/firo-18/pnc/client/create"
	"github.com/firo-18/pnc/client/event"
)

func init() {
	Index["page-menu"] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Message.Interaction.User.String() != i.Member.User.String() {
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Title:       "Error",
							Description: "This interaction is intended for the original user only.",
						},
					},
					Flags: discordgo.MessageFlagsEphemeral,
				},
			})

			if err != nil {
				log.Fatalln("interaction-respond:", err)
			}
		} else {
			data := i.MessageComponentData()
			name, page, _ := strings.Cut(data.Values[0], "__")
			doll, _ := event.DMu.Read(name)

			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
				Data: &discordgo.InteractionResponseData{
					Embeds:     create.DollEmbeds(s, doll, &page),
					Components: create.DollComponents(doll, &page),
				},
			})

			if err != nil {
				log.Fatalln("interaction-respond:", err)
			}
		}
	}
}
