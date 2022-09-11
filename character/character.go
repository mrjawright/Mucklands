package character

import (
	"characterClasses"
	"dice"
	"equipment"
	"fmt"
	"slices"
)

var attributeNames = [...]string{"Vim", "Vigor", "Knack", "Knowhow"}

type Attribute struct {
	Name     string
	Modifier int
}

type Stats struct {
	Courage     int
	Attack      int
	Defense     int
	QuestPoints int
}

type Skill struct {
	Attribute string
	Name      string
	Bonus     int
	Penalty   int
}

func setQuestPoints(character *Character) {
	idx := slices.IndexFunc(character.Attributes, func(a Attribute) bool { return a.Name == "Knowhow" })
	character.QuestPoints = 3 + character.Attributes[idx].Modifier
}

type Character struct {
	Name       string
	Class      characterClasses.CharacterClass
	Species    string
	Attributes []Attribute
	Stats
	Skills        []Skill
	Abilities     []string
	Perks         []string
	Backstory     string
	Ideals        string
	Flaws         string
	PersonalQuest string
	Relationships []string
	Inventory     []equipment.Equipment
}

func initCharacter(character *Character) {
	for x := 0; x < len(attributeNames); x++ {
		character.Attributes = append(character.Attributes, Attribute{attributeNames[x], 0})
	}
}

func selectClass(character *Character) {
	fmt.Println("Selecting character class...")
	idx := dice.GetRandomIndex(len(characterClasses.ClassNames))
	className := characterClasses.ClassNames[idx]
	selectedClass := characterClasses.CharacterClasses[className]
	newClass := characterClasses.CharacterClass{}
	newClass.Name = selectedClass.Name
	newClass.Abilities = selectedClass.Abilities
	newClass.Perks = selectedClass.Perks
	newClass.Story.Relationships = selectedClass.Story.Relationships
	backstoryIdx := dice.GetRandomIndex(6)
	newClass.Story.Backstory = []string{selectedClass.Story.Backstory[backstoryIdx]}
	fmt.Println("...got", character.Class)
}

func assignAttributeModifiers(character *Character) {
	modifiers := []int{-1, 0, 1, 2}
	for _, attrName := range attributeNames {
		index := dice.GetRandomIndex(len(modifiers))
		attr := Attribute{attrName, modifiers[index]}
		modifiers[index] = modifiers[len(modifiers)-1]
		modifiers[len(modifiers)-1] = 0
		modifiers = modifiers[:len(modifiers)-1]
		character.Attributes = append(character.Attributes, attr)
	}
}

func BuildCharacter(character *Character) {
	initCharacter(character)
	assignAttributeModifiers(character)
	selectClass(character)
}
