package info

import (
	"log"
	"os"

	"github.com/firo-18/pnc/api"
	"gopkg.in/yaml.v3"
)

var (
	Classes = [...]string{"Guard", "Medic", "Sniper", "Specialist", "Warrior"}
)

// NewClassData creates a new ClassData and return its address.
func NewClassData() *ClassData {
	return &ClassData{}
}

// Class defines Doll class information structure.
type ClassData struct {
	Name string `yaml:"name"`
	Desc string `yaml:"desc"`
}

// WriteYAML writes Class data to a local YAML file.
func (c ClassData) Write() {
	data, err := yaml.Marshal(c)
	if err != nil {
		log.Fatalln("marshal:", err)
	}

	err = os.WriteFile(Path.ClassData+c.Name+".yaml", data, 0600)
	if err != nil {
		log.Fatalln("write-file:", err)
	}
}

func (c *ClassData) Lookup(name string) error {
	url := Path.Root + Path.ClassData + name + ".yaml"

	err := api.GetDecodeYAML(url, c)

	return err
}
