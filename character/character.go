package character

import (
	"characterClasses"
	"dice"
	"equipment"
	"fmt"
	"slices"
)

var attributeNames = [...]string{"Vim", "Vigor", "Knack", "Knowhow"}
var attributeSkills = map[string][]string{"Vim":[]string {"Charm", "Inspire", "Mettle", "Perception"}, 
"Vigor": []string {"Athletics", "Intimidate", "Might", "Vitality"}, 
"Knack":[]string {"Nimbleness", "Search", "Sneak", "Trickery"}, 
"Knowhow":[]string {"Lore", "Realms", "Tinker", "Wilderness"}}

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
	Name       string
	Class      string
	Species    string
	Attributes map[string]int
	Dread      string
	Stats
	Skills        map[string]Skill
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

func setSkillModifiers(character *Character){
	modifiers := map[int][]int{2:{2,2,2,1}, 1:{1,1,1,0}, 0:{0,0,0,1}, -1:{-1,-1,-1,0}}

	for a,m := range character.Attributes{
		var assignedSkill []string
		var skillName string
		for skillmod := range modifiers[m]{
			idx :=-1
			for idx == -1{
				skillIdx := dice.GetRandomIndex(4)
				skillName = attributeSkills[a][skillIdx]
				idx = slices.IndexFunc(assignedSkill, func(skill string) bool { return skill == skillName })
			}
			assignedSkill = append(assignedSkill, skillName)
			character.Skills[a]=Skill{Name: skillName, Modifier: skillmod}
		}
	}
}

func initCharacter(character *Character) {
	var assignedAttributes []string
	for _, x := range []int{2, 1, 0, -1} {
		idx := -1
		var attrIdx int = -1
		for idx == -1 {
			attrIdx = dice.GetRandomIndex(4)
			var attrName string = attributeNames[attrIdx]
			idx = slices.IndexFunc(assignedAttributes, func(attr string) bool { return attr == attrName })
		}
		assignedAttributes = append(assignedAttributes, attributeNames[attrIdx])
		character.Attributes[attributeNames[attrIdx]] = x
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


func BuildCharacter() Character {
	var character Character = Character{} 
	initCharacter(&character)
	setStats(&character)
	selectClass(&character)
	setSkillModifiers(&character)
	
	return character
}
