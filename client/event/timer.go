package event

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/firo-18/pnc/client/create"
)

var (
	timeServer = 8
)

func init() {
	Index["timer"] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		current := time.Now().UTC()
		server := time.Date(current.Year(), current.Month(), current.Day(), timeServer, 0, 0, 0, time.UTC)
		if current.After(server) {
			server = server.Add(time.Hour * 24)
		}

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:     "Timers",
						Timestamp: create.EmbedTimestamp,
						Color:     create.EmbedColor,
						Footer: &discordgo.MessageEmbedFooter{
							Text:    s.State.User.Username,
							IconURL: s.State.User.AvatarURL(""),
						},
						Fields: []*discordgo.MessageEmbedField{
							{
								Name:   "EN Server Reset",
								Value:  fmt.Sprintf("<t:%v:t>", server.Unix()),
								Inline: true,
							},
						},
					},
				},
			},
		})
		if err != nil {
			log.Fatalln("interaction-respond:", err)
		}

		log.Printf("%v used timer command.", i.Member.User)
	}
}
