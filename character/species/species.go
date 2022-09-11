package species

var SpeciesNames = [...]string{"Boggart", "Boggle", "Bogril", "Gelatinous Goo", "Gnome", "Goblin", "Human", "Imp", "Skeleton", "Shrym", "Wug"}

type Species struct {
	name          string
	homelands     []string
	bonusSkills   []string
	penaltySkills []string
	perk          string
}
