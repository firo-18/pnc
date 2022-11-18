package command

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/firo-18/pnc/client/config"
)

var (
	Index                        = []*discordgo.ApplicationCommand{}
	Test                         = []*discordgo.ApplicationCommand{}
	permissionManageServer int64 = discordgo.PermissionManageServer
	c                            = config.Load("config.json")
)

func DeployTest() {
	s, err := discordgo.New("Bot " + c.Token)
	if err != nil {
		log.Fatalln(`Error setting up discord client:`, err)
	}

	registeredTest, err := s.ApplicationCommandBulkOverwrite(c.ClientID, c.GuildID, Test)
	if err != nil {
		log.Fatalln("Cannot create commands:", err)
	}
	log.Printf("Deployed %v slash commands successfully to test server.", len(registeredTest))
}

func DeployIndex() {
	s, err := discordgo.New("Bot " + c.Token)
	if err != nil {
		log.Fatalln(`Error setting up discord client:`, err)
	}

	registeredIndex, err := s.ApplicationCommandBulkOverwrite(c.ClientID, "", Index)
	if err != nil {
		log.Fatalln("Cannot create commands:", err)
	}
	log.Printf("Deployed %v slash commands successfully to all servers.", len(registeredIndex))
}

func UndeployTest() {
	s, err := discordgo.New("Bot " + c.Token)
	if err != nil {
		log.Fatalln(`Error setting up discord client:`, err)
	}

	_, err = s.ApplicationCommandBulkOverwrite(c.ClientID, c.GuildID, []*discordgo.ApplicationCommand{})
	if err != nil {
		log.Fatalln("Cannot create commands:", err)
	}
	log.Print("Undeployed all slash commands successfully from test server.")
}

func UndeployIndex() {
	s, err := discordgo.New("Bot " + c.Token)
	if err != nil {
		log.Fatalln(`Error setting up discord client:`, err)
	}

	_, err = s.ApplicationCommandBulkOverwrite(c.ClientID, "", []*discordgo.ApplicationCommand{})
	if err != nil {
		log.Fatalln("Cannot create commands:", err)
	}
	log.Print("Undeployed all slash commands successfully from all server.")
}
