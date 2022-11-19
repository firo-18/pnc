package event

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/firo-18/pnc/client/create"
)

var (
	defaultDollPage = "Skills"
)

func init() {
	Index["info"] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			data := i.ApplicationCommandData()
			switch data.Options[0].Name {
			case "doll":
				name := data.Options[0].Options[0].StringValue()
				respondDollInfo(s, i, name)
			case "class":
				if len(data.Options[0].Options) > 1 {
					name := data.Options[0].Options[1].StringValue()
					respondDollInfo(s, i, name)
				} else {
					class := data.Options[0].Options[0].StringValue()
					log.Printf("%v queried for class: %v.", i.Member.User, class)

					err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Embeds: create.ClassEmbeds(s, &Meta, &class),
						},
					})
					if err != nil {
						log.Fatalln("interaction-respond:", err)
					}
				}
			}
		case discordgo.InteractionApplicationCommandAutocomplete:
			data := i.ApplicationCommandData()
			choices := []*discordgo.ApplicationCommandOptionChoice{}
			choice := data.Options[0].Options[0].StringValue()

			switch data.Options[0].Name {
			case "doll":
				for k := range DMu.Dolls {
					if ok, _ := regexp.MatchString("(?i)"+choice, k); ok {
						choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
							Name:  k,
							Value: k,
						})
					}
				}
			case "class":
				if len(data.Options[0].Options) > 1 {
					choice = data.Options[0].Options[1].StringValue()
					for k, v := range DMu.Dolls {
						if v.Class == data.Options[0].Options[0].StringValue() {
							if ok, _ := regexp.MatchString("(?i)"+choice, k); ok {
								choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
									Name:  k,
									Value: k,
								})
							}
						}
					}
				}
			}

			if len(choices) > 25 {
				choices = choices[:25]
			}

			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionApplicationCommandAutocompleteResult,
				Data: &discordgo.InteractionResponseData{
					Choices: choices,
				},
			})
			if err != nil {
				log.Fatalln("autocomplete-result:", err)
			}
		}
	}
}

// respondDollInfo sends a discord interaction respond about a Doll's data.
func respondDollInfo(s *discordgo.Session, i *discordgo.InteractionCreate, name string) {
	log.Printf("%v queried for doll: %v.", i.Member.User, name)

	if doll, ok := DMu.Read(name); !ok {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "404 NOT FOUND",
						Description: fmt.Sprintf("Data on requested Doll '%v' cannot be found.", name),
						Color:       create.EmbedColor,
					},
				},
				Flags: discordgo.MessageFlagsEphemeral,
			}})

		if err != nil {
			log.Fatalln("interaction-respond:", err)
		}
	} else {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds:     create.DollEmbeds(s, doll, &defaultDollPage),
				Components: create.DollComponents(doll, &defaultDollPage),
			},
		})
		if err != nil {
			log.Fatalln("interaction-respond:", err)
		}

		time.Sleep(time.Second * 60)

		_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Components: &[]discordgo.MessageComponent{},
		})

		if err != nil {
			log.Fatalln("interaction-respond:", err)
		}
	}
}
