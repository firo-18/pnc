package info

import (
	"log"
	"os"
	"strings"

	"github.com/firo-18/pnc/api"
	"gopkg.in/yaml.v3"
)

// Meta defines meta data structure.
type MetaData struct {
	Dolls []string `yaml:"dolls"`
}

// NewMetaData creates a new MetaData type with initial local data.
func NewMetaData() *MetaData {
	// Read dir for data json files
	files, err := os.ReadDir(path.DollData)
	if err != nil {
		log.Fatalln("read-dir:", err)
	}

	var meta MetaData

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
				meta.Dolls = append(meta.Dolls, doll.Name)
			}
		}
	}

	return &meta
}

// Update fetches and read into memory the latest meta file in databse.
func (m *MetaData) Update() {
	err := api.GetDecodeYAML(path.Root+"meta.yaml", m)
	if err != nil {
		log.Fatalln("decode:", err)
	}

	log.Println("Meta file was updated successfully.")
}

// WriteYAML implements a function to write meta data to a local yaml file.
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
