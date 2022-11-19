package event

import (
	"github.com/bwmarrin/discordgo"
	"github.com/firo-18/pnc/info"
)

var (
	Index   = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){}
	DMu     = info.DollsMutex{Dolls: map[string]*info.DollProfile{}}
	Classes = map[string][]*info.DollProfile{}
	Meta    info.MetaData
)

func init() {
	go Setup()
}

func Setup() {
	Meta.Update()
	fetchDolls()
}

func fetchDolls() {
	for _, name := range Meta.Dolls {
		doll := info.NewDoll()
		doll.Lookup(name)
		go DMu.Write(doll)
		Classes[doll.Class] = append(Classes[doll.Class], doll)
	}
}
