package component

import "github.com/bwmarrin/discordgo"

var (
	Index = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))
)
