package create

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/firo-18/pnc/info"
)

var (
	EmbedTimestamp = time.Now().Format("2006-01-02T03:04:05-0700")
	EmbedColor     = 13292010
)

// DollEmbeds creates a Doll embed discord respond on a specific Doll.
func DollEmbeds(s *discordgo.Session, doll *info.DollProfile, page *string) []*discordgo.MessageEmbed {
	embeds := []*discordgo.MessageEmbed{
		{
			Title:     fmt.Sprintf("(%v) %v - %v", doll.Class, doll.Name, *page),
			URL:       doll.Links.Wiki,
			Timestamp: EmbedTimestamp,
			Color:     EmbedColor,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL:    "https://raw.githubusercontent.com/firo-18/pnc-db/main/asset/dolls/icons/" + doll.Name + ".png",
				Width:  256,
				Height: 256,
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text:    s.State.User.Username,
				IconURL: s.State.User.AvatarURL(""),
			},
			Fields: DollEmbedFields(doll, page),
		},
	}

	return embeds
}

// DollEmbedFields populates and return a slice of embed fields for a specific Doll page.
func DollEmbedFields(doll *info.DollProfile, page *string) []*discordgo.MessageEmbedField {
	fields := []*discordgo.MessageEmbedField{}

	switch *page {
	case "Bio":
		fields = []*discordgo.MessageEmbedField{
			{
				Name:  "Model",
				Value: DiscordStyle(doll.Model),
			},
			{
				Name:  "Birthday",
				Value: DiscordStyle(doll.Birthday),
			},
			{
				Name:  "Manufacturer",
				Value: DiscordStyle(doll.Manufacturer),
			},
			{
				Name:  "Career",
				Value: DiscordStyle(doll.Career),
			},
			{
				Name:  "Voice",
				Value: DiscordStyle(doll.Voice),
			},
		}
	case "Skills":
		fields = []*discordgo.MessageEmbedField{
			{
				Name:  "Passive - " + doll.Skills.Passive.Name,
				Value: DiscordStyle(doll.Skills.Passive.Desc),
			},
			{
				Name:  "Auto - " + doll.Skills.Auto.Name,
				Value: DiscordStyle(doll.Skills.Auto.Desc),
			},
			{
				Name:  "Ultimate - " + doll.Skills.Ultimate.Name,
				Value: DiscordStyle(doll.Skills.Ultimate.Desc),
			},
		}
	case "Algorithm":
		fields = []*discordgo.MessageEmbedField{
			{
				Name:  "Set Name",
				Value: DiscordStyle(doll.Algorithm.Set),
			},
			{
				Name:  "Main Stat",
				Value: DiscordStyle(doll.Algorithm.Main),
			},
			{
				Name:  "Sub Stats",
				Value: DiscordStyle(doll.Algorithm.Sub),
			},
		}
	case "Analysis":
		fields = []*discordgo.MessageEmbedField{
			{
				Name:  "Rating",
				Value: DiscordStyle(doll.Analysis.Rating),
			},
			{
				Name:  "Detail",
				Value: DiscordStyle(doll.Analysis.Detail),
			},
		}

	}

	return fields
}

// ClassEmbeds creates a class embed discord respond on a specific Doll class.
func ClassEmbeds(s *discordgo.Session, classes map[string][]*info.DollProfile, class *string) []*discordgo.MessageEmbed {
	desc := ""
	switch *class {
	case "Guard":
		desc = "Guards are your frontline tanks. They are your shields and your your most trustworthy allies. They also trust you to have their backs, so do keep them alive."
	case "Medic":
		desc = "Medics are your lifelines. They are the core of the group, so make sure you bring at least 1 Medic along, and make sure to pretect them well."
	case "Sniper":
		desc = "Snipers are range damage dealers. They are mainly stay at the very back focusing on defeating any enemies within their ranges."
	case "Specialist":
		desc = "Specialists are... well, specialists. Their specialities lie in buff, debuff, crowd-control, and status effect. Their arsenals are indispensable in some battles."
	case "Warrior":
		desc = "Warriors are frontline attackers. They adept at killing enemies at close-ranged. They are highly mobile, and are handy to quickly take out prioritized targets."
	}
	embeds := []*discordgo.MessageEmbed{
		{
			Title:       *class,
			Description: desc,
			Timestamp:   EmbedTimestamp,
			Color:       EmbedColor,
			Footer: &discordgo.MessageEmbedFooter{
				Text:    s.State.User.Username,
				IconURL: s.State.User.AvatarURL(""),
			},
			Fields: ClassEmbedFields(classes[*class]),
		},
	}

	return embeds
}

// ClassEmbedFields populates and return a slice of embed fields for a specific Doll class.
func ClassEmbedFields(dolls []*info.DollProfile) []*discordgo.MessageEmbedField {
	fields := []*discordgo.MessageEmbedField{}

	for _, doll := range dolls {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   doll.Name,
			Value:  DiscordStyle(doll.Analysis.Rating),
			Inline: true,
		})
	}

	return fields
}
