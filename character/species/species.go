package species

var SpeciesNames = [...]string{"Boggart", "Boggle", "Bogril", "Gelatinous Goo", "Gnome", "Goblin", "Human", "Imp", "Skeleton", "Shrym", "Wug"}

type Species struct {
	Name          string
	Homelands     []string
	BonusSkills   []string
	PenaltySkills []string
	Perk          string
}

var SpeciesTable = map[string]Species{"Boggart": {Name: "Boggart",
	Homelands:   []string{"The Used T'be Forest", "The Drippy Downs", "The Quagmash", "River Country", "Fleabag County", "The Underlands", "Scalawag Strand"},
	BonusSkills: []string{"Realms", "Charm"}, PenaltySkills: []string{"Mettle", "Lore"},
	Perk: "Start with 3 Proficiences from their Homeland"},
	"Boggle": {Name: "Boggle",
		Homelands:   []string{"The Used T'be Forest", "The Drippy Downs", "The Quagmash", "River Country", "Fleabag County", "The Underlands", "Scalawag Strand"},
		BonusSkills: []string{"Nimbleness", "Sneak"}, PenaltySkills: []string{"Might", "Mettle"},
		Perk: "When Parleying with a Boggle Adversary, a Boggle PC always rolls with  Advantage"},
	"Bogril": {Name: "Bogril",
		Homelands:   []string{"The Used T'be Forest", "The Drippy Downs", "The Quagmash", "River Country", "Fleabag County", "Scalawag Strand"},
		BonusSkills: []string{"Realms", "Vitality"}, PenaltySkills: []string{"Sneak", "Nimbleness"},
		Perk: "Bogrils can hold their breath for 5 minutes without trouble."},
	"Gelatinous Goo": {Name: "Gelatinous Goo",
		Homelands:   []string{"The Underlands"},
		BonusSkills: []string{"Search", "Vitality"}, PenaltySkills: []string{"Intimidate", "Mettle"},
		Perk: "Roll+Vitality to squish and sequuze through tight cracks and openings.\n" +
			"Gelatinous Goos can also excrete a slippery goop om the ground that lasts for several minutes.\n" +
			"Gelatinous Goos have Disadvantage when trying to communicate with someone by themselves."},
	"Gnome": {Name: "Gnome",
		Homelands:   []string{"The Used T'be Forest", "The Drippy Downs", "The Quagmash", "River Country", "The Dingledell"},
		BonusSkills: []string{"Lore", "inspire"}, PenaltySkills: []string{"Might", "Trickery"},
		Perk: "Gnomes can talk to and converse with animals."},
	"Goblin": {Name: "Goblin",
		Homelands:   []string{"The Used T'be Forest", "The Drippy Downs", "The Quagmash", "River Country", "Fleabag County", "The Underlands", "Scalawag Strand"},
		BonusSkills: []string{"Trickery", "Intimidate"}, PenaltySkills: []string{"Lore", "Inspire"},
		Perk: "A goblin may declare two Grudges against different enemies in their life. A grudge provides a +1 bonus to all actions against that enemy."},
	"Human": {Name: "Human",
		Homelands:   []string{"The Used T'be Forest", "The Drippy Downs", "The Quagmash", "River Country", "Fleabag County", "Scalawag Strand"},
		BonusSkills: []string{"Any Skill"}, PenaltySkills: []string{"Any Skill"},
		Perk: "None"},
	"Imp": {Name: "Imp",
		Homelands:   []string{"The Used T'be Forest", "The Drippy Downs", "The Quagmash", "The Underlands"},
		BonusSkills: []string{"Trickery", "Nimbleness"}, PenaltySkills: []string{"Charm", "Might"},
		Perk: "Roll+Trickery - Once per session create a disguise for yourself or an ally. The believability of the disguise varies by the die result, and the disguise is undone if the Imp sneezes."},
	"Skeleton": {Name: "Skeleton",
		Homelands:   []string{"The Underlands"},
		BonusSkills: []string{"Intimidate", "Trickery"}, PenaltySkills: []string{"Vitality", "Charm"},
		Perk: "Roll+Vitality, once per session to block that much Dread.\n" +
			"Skeletons have Disadvantage on Charm checks with strangers."},
	"Shrym": {Name: "Shrym",
		Homelands:   []string{"The Used T'be Forest", "The Drippy Downs", "The Quagmash", "River Country", "Fleabag County", "Scalawag Strand"},
		BonusSkills: []string{"Tinker", "Lore"}, PenaltySkills: []string{"Perception", "Intimidate"},
		Perk: "Shryms have Proficiency in Mechanics"},
	"Wug": {Name: "Wug",
		Homelands:   []string{"The Used T'be Forest", "The Drippy Downs", "The Quagmash", "River Country", "Fleabag County", "Scalawag Strand"},
		BonusSkills: []string{"Might", "Vitality"}, PenaltySkills: []string{"Lore", "Tinker"},
		Perk: "Wugs have a Proficiency in Resisting Charm"},
}
