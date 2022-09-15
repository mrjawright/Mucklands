package character

import (
	"characterClasses"
	"dice"
	"equipment"
	"fmt"
	"homelands"
	"slices"
	"species"
)

var attributeNames = [...]string{"Vim", "Vigor", "Knack", "Knowhow"}
var attributeSkills = map[string][]string{"Vim": {"Charm", "Inspire", "Mettle", "Perception"},
	"Vigor":   {"Athletics", "Intimidate", "Might", "Vitality"},
	"Knack":   {"Nimbleness", "Search", "Sneak", "Trickery"},
	"Knowhow": {"Lore", "Realms", "Tinker", "Wilderness"}}

type Stats struct {
	Courage     int
	Attack      int
	Defense     int
	QuestPoints int
}

func (s Stats) String() string {
	var retVal string
	retVal = fmt.Sprint("Stats:\n")
	retVal += fmt.Sprintf("\tCourage: %d\n", s.Courage)
	retVal += fmt.Sprintf("\tAttack: %d\n", s.Attack)
	retVal += fmt.Sprintf("\tDefense: %d\n", s.Defense)
	retVal += fmt.Sprintf("\tQuestPoints: %d\n", s.QuestPoints)
	return retVal
}

type Skill struct {
	Name     string
	Modifier int
}

type Character struct {
	Name           string
	Class          string
	Species        string
	Homeland       string
	Attributes     map[string]int
	InventorySlots int
	Proficiencies  []string
	Dread          string
	Stats
	Skills        map[string]map[string]int
	Abilities     []string
	Perks         []string
	Backstory     string
	Ideals        string
	Flaws         string
	PersonalQuest string
	Relationships []string
	Inventory     []equipment.Equipment
}

func (c Character) strAbilities() string {
	var retVal string
	for _, ability := range c.Abilities {
		retVal += fmt.Sprintf("%s\n", ability)
	}
	return retVal
}

func (c Character) strAttributes() string {
	var retVal string
	retVal = "Attributes:\n"
	//to for the order they're listed
	for _, attr := range attributeNames {
		retVal += fmt.Sprintf("\t%-8s: %-3d", attr, c.Attributes[attr])
		retVal += "("
		for s, m := range c.Skills[attr] {
			retVal += fmt.Sprintf("%-10s:%-3d", s, m)
		}
		retVal += ")\n"
	}
	return retVal
}

func (c Character) strInventory() string {
	var retVal string
	retVal = fmt.Sprintf("%-30s%-6s%-20s", "Name", "Slots", "Cost")
	for _, i := range c.Inventory {
		retVal += fmt.Sprintf("%s", i.String())
	}
	return retVal
}

func (c Character) strPerks() string {
	var retVal string
	retVal = "Perks:\n"
	for _, p := range c.Perks {
		retVal += fmt.Sprintf("%s\n", p)
	}
	return retVal
}

func (c Character) strProficiencies() string {
	var retVal string
	retVal = "Proficiencies:\n"
	for _, p := range c.Proficiencies {
		retVal += fmt.Sprintf("\t%s\n", p)
	}
	return retVal
}

func (c Character) strRelationships() string {
	var retVal string
	retVal = "Relationships:\n"
	var padLen int = 0
	for _, r := range c.Relationships {
		if padLen < len(r) {
			padLen = len(r)
		}
	}
	for idx, r := range c.Relationships {
		retVal += fmt.Sprintf("(%d)%*s", idx+1, -padLen, r)
		if idx%2 == 1 {
			retVal += "\n"
		} else {
			retVal += " "
		}
	}
	return retVal
}

func (c Character) String() string {
	return fmt.Sprintf("Name: %s\nClass: %s Species:%s Homeland:%s Dread:%s\n%s\n%s\n%s\n%s\n%s\nBackstory:\n%s\nIdeals:%s\nFlaws:%s\nPersonal Quest:%s\n%s",
		c.Name, c.Class, c.Species, c.Homeland, c.Dread, c.strAttributes(), c.strProficiencies(),
		c.Stats.String(), c.strPerks(), c.strInventory(), c.Backstory, c.Ideals, c.Flaws, c.PersonalQuest, c.strRelationships())
}

func setStats(character *Character) {
	character.QuestPoints = 3 + character.Attributes["Knowhow"]
	character.Attack = character.Attributes["Vigor"]
	character.Defense = -1 * character.Attributes["Knack"]
}

func setSkillModifiers(character *Character) {
	modifiers := map[int][]int{2: {2, 2, 2, 1}, 1: {1, 1, 1, 0}, 0: {0, 0, 0, 1}, -1: {-1, -1, -1, 0}}
	inventoryModifier := 0
	character.Skills = make(map[string]map[string]int)
	for a, m := range character.Attributes {
		var assignedSkill []string = make([]string, 0)
		var skillName string
		skillList := make(map[string]int)
		for skillMod := range modifiers[m] {
			idx := -2
			for idx != -1 {
				skillIdx := dice.GetRandomIndex(4)
				skillName = attributeSkills[a][skillIdx]
				idx = slices.IndexFunc(assignedSkill, func(skill string) bool { return skill == skillName })
			}
			assignedSkill = append(assignedSkill, skillName)
			skillList[skillName] = skillMod
			if skillName == "Might" || skillName == "Vitality" {
				inventoryModifier += skillMod
			}
		}
		character.Skills[a] = skillList
	}
	character.InventorySlots = 20 + inventoryModifier
}

func initCharacter(character *Character) {
	var assignedAttributes []string = make([]string, 0)
	character.Attributes = make(map[string]int)
	for _, x := range []int{2, 1, 0, -1} {
		idx := -2
		var attrIdx int = -1
		for idx != -1 {
			attrIdx = dice.GetRandomIndex(4)
			var attrName string = attributeNames[attrIdx]
			idx = slices.IndexFunc(assignedAttributes, func(attr string) bool { return attr == attrName })
		}
		assignedAttributes = append(assignedAttributes, attributeNames[attrIdx])
		fmt.Println(assignedAttributes)
		character.Attributes[attributeNames[attrIdx]] = x
		fmt.Println(character.Attributes)
	}
}

func selectClass(character *Character) {
	fmt.Println("Selecting character class...")
	idx := dice.GetRandomIndex(len(characterClasses.ClassNames))
	className := characterClasses.ClassNames[idx]
	selectedClass := characterClasses.CharacterClasses[className]
	character.Class = className
	character.Stats.Courage = selectedClass.CourageModifier + character.Attributes["Vim"]
	character.Dread = selectedClass.Dread
	character.Abilities = selectedClass.Abilities
	character.Perks = selectedClass.Perks
	character.Relationships = selectedClass.Story.Relationships
	backstoryIdx := dice.GetRandomIndex(len(selectedClass.Story.Backstory))
	character.Backstory = selectedClass.Story.Backstory[backstoryIdx]
	flawsIdx := dice.GetRandomIndex(len(selectedClass.Story.Flaws))
	character.Flaws = selectedClass.Story.Flaws[flawsIdx]
	idealsIdx := dice.GetRandomIndex(len(selectedClass.Story.Ideals))
	character.Ideals = selectedClass.Story.Ideals[idealsIdx]
	questIdx := dice.GetRandomIndex(len(selectedClass.Story.PersonalQuest))
	character.PersonalQuest = selectedClass.Story.PersonalQuest[questIdx]
	fmt.Println("...got", character.Class)
}

func selectSpecies(character *Character) {
	speciesIdx := dice.GetRandomIndex(len(species.SpeciesNames))
	selectedSpecies := species.SpeciesTable[species.SpeciesNames[speciesIdx]]
	character.Species = selectedSpecies.Name
	character.Perks = append(character.Perks, selectedSpecies.Perk)
	if len(selectedSpecies.BonusSkills) > 1 {
		bonusSkill := selectedSpecies.BonusSkills[dice.GetRandomIndex(2)]
		fmt.Println("Bonus to " + bonusSkill)
		for attr, skills := range character.Skills {
			for skill, _ := range skills {
				if skill == bonusSkill {
					character.Skills[attr][skill] += 1
				}
			}
		}
	}
	if len(selectedSpecies.PenaltySkills) > 1 {
		penaltySkill := selectedSpecies.PenaltySkills[dice.GetRandomIndex(2)]
		fmt.Println("Penalty to " + penaltySkill)
		for attr, skills := range character.Skills {
			for skill, _ := range skills {
				if skill == penaltySkill {
					character.Skills[attr][skill] -= 1
				}
			}
		}
	}
	homelandIdx := dice.GetRandomIndex(len(selectedSpecies.Homelands))
	selectedSpeciesHomeland := selectedSpecies.Homelands[homelandIdx]
	selectedHomeland := homelands.Homelands[selectedSpeciesHomeland]
	character.Homeland = selectedHomeland.Name
	character.Inventory = selectedHomeland.Equipment
	proficencyIdx := dice.GetRandomIndex(len(selectedHomeland.Proficiencies))
	character.Proficiencies = append(character.Proficiencies, selectedHomeland.Proficiencies[proficencyIdx])
	var idx int = -1
	for idx != proficencyIdx {
		idx = dice.GetRandomIndex(len(selectedHomeland.Proficiencies))
		if idx != proficencyIdx {
			character.Proficiencies = append(character.Proficiencies, selectedHomeland.Proficiencies[idx])
			break
		}
	}
}

func BuildCharacter() Character {
	var character Character = Character{}
	fmt.Println("Init character...")
	initCharacter(&character)
	fmt.Println("Set character stats...")
	setStats(&character)
	fmt.Println("Select class...")
	selectClass(&character)
	fmt.Println("Got " + character.Class)
	fmt.Println("Set skill modifiers...")
	setSkillModifiers(&character)
	fmt.Println("Select species and homeland...")
	selectSpecies(&character)
	fmt.Println("Got " + character.Species + " from " + character.Homeland)
	equipmentIdx := dice.Roll(1, 200, 0)
	randomItem := equipment.EquipmentTable[equipmentIdx]
	character.Inventory = append(character.Inventory, randomItem)
	return character
}
