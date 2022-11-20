package info

// DBPath defines directory structure for this database.
type DBPath struct {
	Root      string
	DollData  string
	DollIcon  string
	ClassData string
}

var (
	path = DBPath{
		Root:      "https://raw.githubusercontent.com/firo-18/pnc-db/main/",
		DollData:  "data/dolls/",
		DollIcon:  "asset/dolls/icons/",
		ClassData: "data/classes/",
	}
)
