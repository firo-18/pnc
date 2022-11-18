package create

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/firo-18/pnc/info"
)

func DollComponents(doll *info.DollProfile, page *string) []discordgo.MessageComponent {
	pages := []string{"Bio", "Skills", "Algorithm", "Analysis"}
	menuOptions := []discordgo.SelectMenuOption{}

	for _, p := range pages {
		menuOptions = append(menuOptions, SelectMenuOption(doll, &p, page))
	}

	components := []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.SelectMenu{
					CustomID: "page-menu",
					Options:  menuOptions,
				},
			},
		},
	}

	return components
}

func SelectMenuOption(doll *info.DollProfile, p, page *string) discordgo.SelectMenuOption {
	option := discordgo.SelectMenuOption{
		Label:       *p,
		Value:       doll.Name + "__" + *p,
		Description: fmt.Sprintf("View %v's %v.", doll.Name, strings.ToLower(*p)),
		Default:     *p == *page,
	}
	return option
}
