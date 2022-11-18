package event

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/firo-18/pnc/client/create"
)

func init() {
	Index["set"] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		data := i.ApplicationCommandData()
		switch data.Options[0].Name {
		case "server-time":
			timeServer = int(data.Options[0].Options[0].IntValue())
			utc := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), timeServer, 0, 0, 0, time.UTC)

			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Title:       "Success",
							Description: fmt.Sprintf("Server reset time has been successfully set to <t:%v:t>", utc.Unix()),
							Color:       create.EmbedColor,
						},
					},
					Flags: discordgo.MessageFlagsEphemeral,
				},
			})
			if err != nil {
				log.Fatalln("interaction-respond:", err)
			}
			log.Printf("%v set server time to %v.", i.Member.User, utc)
		}
	}
}
