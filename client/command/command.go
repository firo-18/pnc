package command

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/firo-18/pnc/client/config"
)

var (
	Index                        = []*discordgo.ApplicationCommand{}
	permissionManageServer int64 = discordgo.PermissionManageServer
)

// DeployProd deploys all commands into production discord bot.
func DeployProduction() {
	config := config.Load("config-hk416.json")

	s, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatalln(`Error setting up discord client:`, err)
	}

	registeredIndex, err := s.ApplicationCommandBulkOverwrite(config.ClientID, "", Index)
	if err != nil {
		log.Fatalln("Cannot create commands:", err)
	}
	log.Printf("Deployed %v slash production commands successfully to all servers.", len(registeredIndex))
}

// DeployProd deploys all commands into production discord bot.
func DeployTest() {
	config := config.Load("config.json")

	s, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatalln(`Error setting up discord client:`, err)
	}

	registeredIndex, err := s.ApplicationCommandBulkOverwrite(config.ClientID, config.GuildID, Index)
	if err != nil {
		log.Fatalln("Cannot create commands:", err)
	}
	log.Printf("Deployed %v slash test commands successfully to test server.", len(registeredIndex))
}
