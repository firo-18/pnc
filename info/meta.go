package info

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/firo-18/pnc/api"
	"gopkg.in/yaml.v3"
)

// DBPath defines directory structure for this database.
type DBPath struct {
	Root     string
	DollData string
	DollIcon string
}

var (
	path = DBPath{
		Root:     "https://raw.githubusercontent.com/firo-18/pnc-db/main/",
		DollData: "data/dolls/",
		DollIcon: "asset/dolls/icons/",
	}
)

// Meta defines meta file structure
type MetaData struct {
	Dolls     map[string]string            `yaml:"dolls"`
	Classes   map[string]map[string]string `yaml:"classes"`
	Companies map[string]map[string]string `yaml:"companies"`
}

// NewMetaData creates a new MetaData type with initial local data.
func NewMetaData() *MetaData {
	meta := MetaData{
		Dolls:     map[string]string{},
		Classes:   map[string]map[string]string{},
		Companies: map[string]map[string]string{},
	}

	// Read dir for data json files
	files, err := os.ReadDir(path.DollData)
	if err != nil {
		log.Fatalln("read-dir:", err)
	}

	// Loop over all files found
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".yaml") {
			data, err := os.ReadFile(path.DollData + file.Name())
			if err != nil {
				log.Fatalln("read-file:", err)
			}

			doll := NewDoll()
			err = yaml.Unmarshal(data, &doll)
			if err != nil {
				log.Fatalln("unmarshal:", err)
			}
			if doll.Verify() {
				releaseState := "Upcoming"
				if doll.Release.Before(time.Now()) {
					releaseState = "Released"
				}
				meta.Dolls[doll.Name] = releaseState
				if _, ok := meta.Classes[doll.Class]; !ok {
					meta.Classes[doll.Class] = make(map[string]string)
				}
				if _, ok := meta.Companies[doll.Manufacturer]; !ok {
					meta.Companies[doll.Manufacturer] = make(map[string]string)
				}
				meta.Classes[doll.Class][doll.Name] = doll.Analysis.Rating
				meta.Companies[doll.Manufacturer][doll.Name] = doll.Model
			}
		}
	}

	return &meta
}

// Update implements update function for meta data to the latest version from API.
func (m *MetaData) Update() {
	err := api.GetDecodeYAML(path.Root+"meta.yaml", m)
	if err != nil {
		log.Fatalln("decode:", err)
	}

	log.Println("Meta file was updated successfully.")
}

// WriteYAML implements a function to write meta data to a yaml file.
func (m MetaData) WriteYAML() {
	metaYAML, err := yaml.Marshal(m)
	if err != nil {
		log.Fatalln("marshal:", err)
	}

	err = os.WriteFile("meta.yaml", metaYAML, 0600)
	if err != nil {
		log.Fatalln("write-file:", err)
	}

	log.Println("Meta file was written successfully.")
}
