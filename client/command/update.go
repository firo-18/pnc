package command

import "github.com/bwmarrin/discordgo"

func init() {
	Test = append(Test, &discordgo.ApplicationCommand{
		Name:                     "update",
		Description:              "Update refreshes memory data and fetches from database the latest data version.",
		Type:                     discordgo.ChatApplicationCommand,
		DefaultMemberPermissions: &permissionManageServer,
	})
}
