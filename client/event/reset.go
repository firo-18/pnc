package event

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/firo-18/pnc/client/create"
)

func init() {
	Index["reset"] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		go Setup()
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "Success",
						Description: fmt.Sprintf("Reinitializing database variables for **%v**...", s.State.User.String()),
						Color:       create.EmbedColor,
						Timestamp:   create.EmbedTimestamp,
						Footer:      create.EmbedFooter(s),
					},
				},
				Flags: discordgo.MessageFlagsEphemeral,
			},
		})
		if err != nil {
			log.Fatalln("interaction-respond:", err)
		}

		log.Printf("%v updated %v's memory data.", i.Member.User, s.State.User)
	}
}
