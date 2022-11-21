package info

import (
	"log"
	"os"
	"time"

	"github.com/firo-18/pnc/api"
	"gopkg.in/yaml.v3"
)

// Doll defines information about a Doll.
type DollProfile struct {
	Name      string        `json:"name" yaml:"name"`
	Bio       DollBio       `json:"bio" yaml:"bio"`
	Skills    DollSkills    `json:"skills" yaml:"skills"`
	Algorithm DollAlgorithm `json:"algorithm,omitempty" yaml:"algorithm,omitempty"`
	Analysis  DollAnalysis  `json:"analysis,omitempty" yaml:"analysis,omitempty"`
	Links     DollLink      `json:"links,omitempty" yaml:"links,omitempty"`
}

// DollBio lits all the relevant bio data on a Doll.
type DollBio struct {
	Rarity       byte      `json:"rarity" yaml:"rarity"`
	Class        string    `json:"class" yaml:"class"`
	Model        string    `json:"model" yaml:"model"`
	Manufacturer string    `json:"manufacturer" yaml:"manufacturer"`
	Career       string    `json:"career" yaml:"career"`
	Birthday     string    `json:"birthday" yaml:"birthday"`
	Release      time.Time `json:"release" yaml:"release"`
	Voice        string    `json:"voice" yaml:"voice"`
}

// Skillset lists Doll's skills. Each Doll has 3 skills; a passive, an auto, and an ultimate.
type DollSkills struct {
	Passive  DollSkill `json:"passive" yaml:"passive"`
	Auto     DollSkill `json:"auto" yaml:"auto"`
	Ultimate DollSkill `json:"ultimate" yaml:"ultimate"`
}

// Skill stores skill's name and description.
type DollSkill struct {
	Name string `json:"name" yaml:"name"`
	Desc string `json:"desc" yaml:"desc"`
}

// Analysis stores analysis information for doll.
type DollAnalysis struct {
	Rating string `json:"rating,omitempty" yaml:"rating,omitempty"`
	Detail string `json:"detail,omitempty" yaml:"detail,omitempty"`
}

// DollAlgorithm lists the recommended Alorithm set and stats for a doll.
type DollAlgorithm struct {
	Set   string `json:"set,omitempty" yaml:"set,omitempty"`
	Main  string `json:"main,omitempty" yaml:"main,omitempty"`
	Sub   string `json:"sub,omitempty" yaml:"sub,omitempty"`
	Image string `json:"image,omitempty" yaml:"image,omitempty"`
}

// DollLink lists resources URL relavant to a doll.
type DollLink struct {
	Wiki  string `json:"wiki,omitempty" yaml:"wiki,omitempty"`
	Video string `json:"video,omitempty" yaml:"video,omitempty"`
}

// NewDoll creates a new Doll with initial values, then return its address.
func NewDoll() *DollProfile {
	doll := &DollProfile{
		Bio: DollBio{
			Model:        "Classified",
			Birthday:     "Classified",
			Manufacturer: "Classified",
			Career:       "Classified",
			Voice:        "Classified",
		},
		Algorithm: DollAlgorithm{
			Set:  "Classified",
			Main: "Classified",
			Sub:  "Classified",
		},
		Analysis: DollAnalysis{
			Rating: "Classified",
			Detail: "Analysis data of this unit has not been declassified. Top level clearance is required to view this data.",
		},
	}
	return doll
}

// Lookup searchs for a Doll in database by name and return error, if any.
func (d *DollProfile) Lookup(name string) error {
	url := Path.Root + Path.DollData + name + ".yaml"

	err := api.GetDecodeYAML(url, d)

	return err
}

// Verify implements function to verify if DollProfile has the minimum information required.
func (d DollProfile) Verify() bool {
	if d.Name == "" || d.Bio.Class == "" {
		return false
	}
	if d.Skills.Auto.Name == "" || d.Skills.Passive.Name == "" || d.Skills.Ultimate.Name == "" {
		return false
	}
	return true
}

// ReadYAML implements a function to read a local yaml file using Doll's name.
func (d *DollProfile) ReadYAML(filename string, path string) {
	file, err := os.ReadFile(path + filename)
	if err != nil {
		log.Fatalln("read-file:", err)
	}
	err = yaml.Unmarshal(file, d)
	if err != nil {
		log.Fatalln("yaml-unmarshal:", err)
	}
}

// WriteYAML implements a function to write/update to a YAML file.
func (d DollProfile) WriteYAML(Path string) {
	if d.Name == "" {
		log.Println("Cannot create a doll file with blank name.")
	} else {
		data, err := yaml.Marshal(d)
		if err != nil {
			log.Fatalln("marshal:", err)
		}

		err = os.WriteFile(Path+d.Name+".yaml", data, 0600)
		if err != nil {
			log.Fatalln("write-file:", err)
		}

		log.Printf("[%v]'s YAML file was written successfully.", d.Name)
	}
}
