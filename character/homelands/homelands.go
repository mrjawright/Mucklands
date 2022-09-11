package homelands

import "equipment"

var HomelandsNames = [...]string{"The Drippy Downs", "Fleabag County", "The Quagmash", "River Country", "Scalawag Strand", "The Used T'be Forest", "The Dingledell", "The Underlands"}

type Homeland struct {
	name          string
	proficiencies []string
	equipment     []equipment.Equipment
}
