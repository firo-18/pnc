package event

import (
	"github.com/bwmarrin/discordgo"
	"github.com/firo-18/pnc/info"
)

var (
	Meta    info.MetaData
	DollsMu = info.DollsMutex{Dolls: map[string]*info.DollProfile{}}
	Index   = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){}
)

func init() {
	Meta.Update()
	fetchDolls()
}

func fetchDolls() {
	for name := range Meta.Dolls {
		go fetchDoll(name)
	}
}

func fetchDoll(name string) {
	doll := info.NewDoll()
	doll.Lookup(name + ".yaml")
	DollsMu.Write(doll)
}
