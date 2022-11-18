package event

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func init() {
	Index["update"] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		Meta.Update()
		fetchDolls()
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "```fix\nMeta data was updated successfully.```",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		if err != nil {
			log.Fatalln("interaction-respond:", err)
		}

		log.Printf("%v updated %v's memory data.", i.Member.User, s.State.User)

	}
}
