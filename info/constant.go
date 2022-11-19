package info

// DBPath defines directory structure for this database.
type DBPath struct {
	Root        string
	DollData    string
	CompanyData string
	DollIcon    string
}

var (
	path = DBPath{
		Root:        "https://raw.githubusercontent.com/firo-18/pnc-db/main/",
		DollData:    "data/dolls/",
		CompanyData: "data/companies/",
		DollIcon:    "asset/dolls/icons/",
	}
)

var (
	Classes   = [...]string{"Guard", "Medic", "Sniper", "Specialist", "Warrior"}
	Companies = [...]string{"42Lab", "Cyber Media", "Svarog Heavy Industries", "Ultimate Life Holdings", "Universal Anything Services"}
)
