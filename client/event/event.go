package event

import (
	"github.com/bwmarrin/discordgo"
	"github.com/firo-18/pnc/info"
)

var (
	Index     = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){}
	DMu       = info.DollsMutex{Dolls: map[string]*info.DollProfile{}}
	Classes   = map[string][]*info.DollProfile{}
	ClassData = map[string]*info.ClassData{}
	Meta      info.MetaData
)

func init() {
	go Setup()
}

func Setup() {
	go fetchClassData()
	Meta.Update()
	fetchDolls()
}

func fetchDolls() {
	Classes = map[string][]*info.DollProfile{}
	for _, name := range Meta.Dolls {
		doll := info.NewDoll()
		doll.Lookup(name)
		go DMu.Write(doll)
		Classes[doll.Bio.Class] = append(Classes[doll.Bio.Class], doll)
	}
}

func fetchClassData() {
	for _, v := range info.Classes {
		cd := info.NewClassData()
		cd.Lookup(v)
		ClassData[cd.Name] = cd
	}
}
