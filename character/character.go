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
	Proficiences   []string
	Dread          string
	Stats
	Skills        map[string][]Skill
	Abilities     []string
	Perks         []string
	Backstory     string
	Ideals        string
	Flaws         string
	PersonalQuest string
	Relationships []string
	Inventory     []equipment.Equipment
}

func setStats(character *Character) {
	character.QuestPoints = 3 + character.Attributes["Knowhow"]
	character.Attack = character.Attributes["Vigor"]
	character.Defense = -1 * character.Attributes["Knack"]
}

func setSkillModifiers(character *Character) {
	modifiers := map[int][]int{2: {2, 2, 2, 1}, 1: {1, 1, 1, 0}, 0: {0, 0, 0, 1}, -1: {-1, -1, -1, 0}}
	inventoryModifier := 0
	character.Skills = make(map[string][]Skill)
	for a, m := range character.Attributes {
		var assignedSkill []string = make([]string, 0)
		var skillName string
		skillList := []Skill{}
		for skillMod := range modifiers[m] {
			idx := -2
			for idx != -1 {
				skillIdx := dice.GetRandomIndex(4)
				skillName = attributeSkills[a][skillIdx]
				idx = slices.IndexFunc(assignedSkill, func(skill string) bool { return skill == skillName })
			}
			assignedSkill = append(assignedSkill, skillName)
			skillList = append(skillList, Skill{Name: skillName, Modifier: skillMod})
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
	backstoryIdx := dice.GetRandomIndex(6)
	character.Backstory = selectedClass.Story.Backstory[backstoryIdx]
	flawsIdx := dice.GetRandomIndex(6)
	character.Flaws = selectedClass.Story.Flaws[flawsIdx]
	idealsIdx := dice.GetRandomIndex(6)
	character.Ideals = selectedClass.Story.Ideals[idealsIdx]
	questIdx := dice.GetRandomIndex(6)
	character.PersonalQuest = selectedClass.Story.PersonalQuest[questIdx]
	fmt.Println("...got", character.Class)
}

func selectSpecies(character *Character) {
	speciesIdx := dice.GetRandomIndex(len(species.SpeciesNames))
	selectedSpecies := species.SpeciesTable[species.SpeciesNames[speciesIdx]]
	character.Species = selectedSpecies.Name
	character.Perks = append(character.Perks, selectedSpecies.Perk)
	bonusSkill := selectedSpecies.BonusSkills[dice.GetRandomIndex(2)]
	for attr, skills := range character.Skills {
		for idx, skill := range skills {
			if skill.Name == bonusSkill {
				character.Skills[attr][idx].Modifier += 1
			}
		}
	}
	penaltySkill := selectedSpecies.PenaltySkills[dice.GetRandomIndex(2)]
	for attr, skills := range character.Skills {
		for idx, skill := range skills {
			if skill.Name == penaltySkill {
				character.Skills[attr][idx].Modifier -= 1
			}
		}
	}
	homelandIdx := dice.GetRandomIndex(len(selectedSpecies.Homelands))
	selectedSpeciesHomeland := selectedSpecies.Homelands[homelandIdx]
	selectedHomeland := homelands.Homelands[selectedSpeciesHomeland]
	character.Homeland = selectedHomeland.Name
	character.Inventory = selectedHomeland.Equipment
	proficencyIdx := dice.GetRandomIndex(len(selectedHomeland.Proficiencies))
	character.Proficiences = append(character.Proficiences, selectedHomeland.Proficiencies[proficencyIdx])
	var idx int = -1
	for idx != proficencyIdx {
		idx = dice.GetRandomIndex(len(selectedHomeland.Proficiencies))
		if idx != proficencyIdx {
			character.Proficiences = append(character.Proficiences, selectedHomeland.Proficiencies[idx])
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
