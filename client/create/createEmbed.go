package create

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/firo-18/pnc/info"
)

var (
	dateFormat     = "2006-01-02T03:04:05-0700"
	EmbedTimestamp = time.Now().Format(dateFormat)
	EmbedColor     = 13292010
	EmbedFooter    = func(s *discordgo.Session) *discordgo.MessageEmbedFooter {
		return &discordgo.MessageEmbedFooter{
			Text:    s.State.User.Username,
			IconURL: s.State.User.AvatarURL(""),
		}
	}
)

// DollEmbeds creates a Doll embed discord respond on a specific Doll.
func DollEmbeds(s *discordgo.Session, doll *info.DollProfile, page *string) []*discordgo.MessageEmbed {
	embeds := []*discordgo.MessageEmbed{
		{
			Title:     fmt.Sprintf("(%v) %v - %v", doll.Bio.Class, doll.Name, *page),
			URL:       doll.Links.Wiki,
			Timestamp: EmbedTimestamp,
			Color:     EmbedColor,
			Footer:    EmbedFooter(s),
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL:    "https://raw.githubusercontent.com/firo-18/pnc-db/main/asset/dolls/icons/" + doll.Name + ".png",
				Width:  256,
				Height: 256,
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
				Name:  "Rarity",
				Value: DiscordStyle(doll.Bio.Rarity),
			},
			{
				Name:  "Model",
				Value: DiscordStyle(doll.Bio.Model),
			},
			{
				Name:  "Manufacturer",
				Value: DiscordStyle(doll.Bio.Manufacturer),
			},
			{
				Name:  "Career",
				Value: DiscordStyle(doll.Bio.Career),
			},
			{
				Name:  "Birthday",
				Value: DiscordStyle(doll.Bio.Birthday),
			},
			{
				Name:  "Voice",
				Value: DiscordStyle(doll.Bio.Voice),
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
func ClassEmbeds(s *discordgo.Session, classes map[string][]*info.DollProfile, classData *info.ClassData) []*discordgo.MessageEmbed {
	embeds := []*discordgo.MessageEmbed{
		{
			Title:       classData.Name,
			Description: classData.Desc,
			Timestamp:   EmbedTimestamp,
			Color:       EmbedColor,
			Footer:      EmbedFooter(s),
			Fields:      ClassEmbedFields(classes[classData.Name]),
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
