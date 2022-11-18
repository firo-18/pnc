package info

import (
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/firo-18/pnc/api"
	"gopkg.in/yaml.v3"
)

// Doll defines information about a Doll.
type DollProfile struct {
	Model        string        `json:"model" yaml:"model"`
	Name         string        `json:"name" yaml:"name"`
	Class        string        `json:"class" yaml:"class"`
	Birthday     string        `json:"birthday" yaml:"birthday"`
	Release      time.Time     `json:"release" yaml:"release"`
	Manufacturer string        `json:"manufacturer" yaml:"manufacturer"`
	Career       string        `json:"career" yaml:"career"`
	Voice        string        `json:"voice" yaml:"voice"`
	Skills       DollSkills    `json:"skills" yaml:"skills"`
	Algorithm    DollAlgorithm `json:"algorithm,omitempty" yaml:"algorithm,omitempty"`
	Analysis     DollAnalysis  `json:"analysis,omitempty" yaml:"analysis,omitempty"`
	Links        DollLink      `json:"links,omitempty" yaml:"links,omitempty"`
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
	Wiki     string `json:"wiki,omitempty" yaml:"wiki,omitempty"`
	Ultimate string `json:"ultimate,omitempty" yaml:"ultimate,omitempty"`
}

type DollsMutex struct {
	mu    sync.Mutex
	Dolls map[string]*DollProfile
}

func NewDoll() *DollProfile {
	doll := &DollProfile{
		Birthday:     "Classified",
		Manufacturer: "Classified",
		Career:       "Classified",
		Voice:        "Classified",
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

// ReadYAML implements a function to read a yaml file using Doll's name.
func (d *DollProfile) ReadYAML(name string) {
	file, err := os.ReadFile(path.DollData + name + ".yaml")
	if err != nil {
		log.Fatalln("read-file:", err)
	}
	err = yaml.Unmarshal(file, d)
	if err != nil {
		log.Fatalln("yaml-unmarshal:", err)
	}
}

// WriteYAML implements a function to write/update to a yaml file.
func (d DollProfile) WriteYAML() {
	if d.Name == "" {
		log.Println("Cannot create a doll file with blank name.")
	} else {
		dataYAML, err := yaml.Marshal(d)
		if err != nil {
			log.Fatalln("marshal:", err)
		}

		err = os.WriteFile(path.DollData+d.Name+".yaml", dataYAML, 0600)
		if err != nil {
			log.Fatalln("write-file:", err)
		}

		log.Printf("[%v]'s file was written successfully.", d.Name)
	}
}

// Lookup queries doll by filename.
func (d *DollProfile) Lookup(filename string) error {
	url := path.Root + path.DollData + filename

	err := api.GetDecodeYAML(url, d)

	return err
}

// Verify implements function to verify if DollProfile has the minimum information required.
func (d DollProfile) Verify() bool {
	if d.Name == "" || d.Model == "" || d.Class == "" {
		return false
	}
	if d.Skills.Auto.Name == "" || d.Skills.Passive.Name == "" || d.Skills.Ultimate.Name == "" {
		return false
	}
	return true
}

// UpdateStructure loops through all dolls' data and re-write to local an updated json.
func UpdateStructure() {
	files, err := os.ReadDir(path.DollData)
	if err != nil {
		log.Fatalln("read-dir:", err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".yaml") {
			data, err := os.ReadFile(path.DollData + file.Name())
			if err != nil {
				log.Fatalln("read-file:", err)
			}

			doll := DollProfile{}
			err = yaml.Unmarshal(data, &doll)
			if err != nil {
				log.Fatalln("unmarshal:", err)
			}
			doll.WriteYAML()
		}
	}
}

func (dm *DollsMutex) Write(doll *DollProfile) {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	dm.Dolls[doll.Name] = doll
}

func (dm *DollsMutex) Read(field string) (*DollProfile, bool) {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	v, ok := dm.Dolls[field]

	return v, ok
}
