package main

import (
	"character"
	"fmt"
)

func main() {
	fmt.Println("Begin character generation...")
	var newCharacter character.Character = character.BuildCharacter()
	fmt.Println(newCharacter)
	fmt.Println("Character generation complete...")
}
