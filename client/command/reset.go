package command

import "github.com/bwmarrin/discordgo"

func init() {
	Index = append(Index, &discordgo.ApplicationCommand{
		Name:                     "reset",
		Description:              "Reset refetches and reinitialises to the latest version from database. Access level: 'ADMIN'.",
		Type:                     discordgo.ChatApplicationCommand,
		DefaultMemberPermissions: &permissionManageServer,
	})
}
