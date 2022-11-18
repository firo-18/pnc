package event

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/firo-18/pnc/client/create"
)

func init() {
	Index["stats"] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "Stats",
						Description: fmt.Sprintf("Viewing %v's stats.", s.State.User.Username),
						Color:       create.EmbedColor,
						Fields: []*discordgo.MessageEmbedField{
							{
								Name:  "Presence In Guilds",
								Value: fmt.Sprint(len(s.State.Guilds)),
							},
						},
					},
				},
			},
		})
		if err != nil {
			log.Fatalln("interaction-respond:", err)
		}

		log.Printf("%v viewed %v's stats.", i.Member.User, s.State.User)
	}
}
