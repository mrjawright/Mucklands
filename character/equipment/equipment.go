package equipment

import (
	"dice"
)

type Equipment struct {
	Name    string
	Cost    string
	Slots   int
	Carried bool
}

var EquipmentTable = map[int]Equipment{
	1:   {"Axe (Bladed)", "Cheap", 2, true},
	2:   {"Ball and chain (Flexible)", "Cheap", 2, true},
	3:   {"Bat (Blunt)", "Cheap", 1, true},
	4:   {"Blowgun (Ranged)", "Cheap", 1, true},
	5:   {"Bolas (Entangle)", "Cheap", 1, true},
	6:   {"Boomerang (Ranged)", "Cheap", 1, true},
	7:   {"Claws (Bladed)", "Cheap", 1, true},
	8:   {"Cleaver (Bladed)", "Cheap", 2, true},
	9:   {"Club (Blunt)", "Cheap", 1, true},
	10:  {"Cookware (Tool)", "Cheap", 1, true},
	11:  {"Crampons (Tool)", "Cheap", 1, true},
	12:  {"Dagger (Bladed)", "Cheap", 1, true},
	13:  {"Darts (Ranged)", "Cheap", 1, true},
	14:  {"Javelin (Ranged)", "Cheap", 1, true},
	15:  {"Knife (Bladed)", "Cheap", 1, true},
	16:  {"Lash (Flexible)", "Cheap", 1, true},
	17:  {"Main Gauche", "Cheap", 1, true},
	18:  {"Mallet (Blunt)", "Cheap", 2, true},
	19:  {"Quarterstaff (Blunt)", "Cheap", 2, true},
	20:  {"Sap (Blunt)", "Cheap", 1, true},
	21:  {"Scythe (Bladed)", "Cheap", 3, true},
	22:  {"Short Sword (Bladed)", "Cheap", 1, true},
	23:  {"Sickle (Bladed)", "Cheap", 1, true},
	24:  {"Sling (Ranged)", "Cheap", 1, true},
	25:  {"Slingshot (Ranged)", "Cheap", 1, true},
	26:  {"Spear (Polearm)", "Cheap", 2, true},
	27:  {"Staff (Blunt)", "Cheap", 2, true},
	28:  {"Throwing Knives (Ranged)", "Cheap", 1, true},
	29:  {"Truncheon (Blunt)", "Cheap", 2, true},
	30:  {"Whip (Flexible)", "Cheap", 1, true},
	31:  {"Abacus (Tool)", "Cheap", 1, true},
	32:  {"Alarm (Tool)", "Cheap", 1, true},
	33:  {"Bag (Tool)", "Cheap", 0, true},
	34:  {"Ball and chain (Tool)", "Cheap", 3, true},
	35:  {"Bandolier (Tool)", "Cheap", 1, true},
	36:  {"Barrel (Tool)", "Cheap", 3, true},
	37:  {"Bedroll (Tool)", "Cheap", 1, true},
	38:  {"Bell (Tool)", "Cheap", 0, true},
	39:  {"Bellows (Tool)", "Cheap", 1, true},
	40:  {"Blanket (Tool)", "Cheap", 1, true},
	41:  {"Bottle (Tool)", "Cheap", 1, true},
	42:  {"Broom (Tool)", "Cheap", 2, true},
	43:  {"Bucket (Tool)", "Cheap", 1, true},
	44:  {"Cage (Tool)", "Cheap", 2, true},
	45:  {"Cattle prod (Tool)", "Cheap", 1, true},
	46:  {"Cauldron (Tool)", "Cheap", 3, true},
	47:  {"Chain (Tool)", "Cheap", 2, true},
	48:  {"Chalk (Tool)", "Cheap", 0, true},
	49:  {"Dye (Tool)", "Cheap", 0, true},
	50:  {"Fanny Pack (Tool)", "Cheap", 0, true},
	51:  {"Flask (Tool)", "Cheap", 1, true},
	52:  {"Fishing Pole (Tool)", "Cheap", 2, true},
	53:  {"Frying Pan (Tool)", "Cheap", 1, true},
	54:  {"Goggles (Tool)", "Cheap", 1, true},
	55:  {"Gong (Tool)", "Cheap", 1, true},
	56:  {"Galoshes (Tool)", "Cheap", 1, true},
	57:  {"Hammer (Tool)", "Cheap", 1, true},
	58:  {"Hook (Tool)", "Cheap", 1, true},
	59:  {"Lantern (Tool)", "Cheap", 1, true},
	60:  {"Lasso (Tool)", "Cheap", 1, true},
	61:  {"Magnet (Tool)", "Cheap", 0, true},
	62:  {"Mirror (Tool)", "Cheap", 1, true},
	63:  {"Net (Tool)", "Cheap", 2, true},
	64:  {"Paint (Tool)", "Cheap", 1, true},
	65:  {"Parchment (Tool)", "Cheap", 1, true},
	66:  {"Pickaxe (Tool)", "Cheap", 2, true},
	67:  {"Prybar (Tool)", "Cheap", 1, true},
	68:  {"Quill and Inkpot (Tool)", "Cheap", 0, true},
	69:  {"Rope (50') (Tool)", "Cheap", 2, true},
	70:  {"Rucksack (Tool)", "Cheap", 0, true},
	71:  {"Satchel (Tool)", "Cheap", 0, true},
	72:  {"Saw (Tool)", "Cheap", 2, true},
	73:  {"Scissors (Tool)", "Cheap", 1, true},
	74:  {"Sewing Kit (Tool)", "Cheap", 1, true},
	75:  {"Shackles (Tool)", "Cheap", 2, true},
	76:  {"Shovel (Tool)", "Cheap", 2, true},
	77:  {"Stilts (Tool)", "Cheap", 2, true},
	78:  {"String (Tool)", "Cheap", 0, true},
	79:  {"Tankard (Tool)", "Cheap", 3, true},
	80:  {"Tent (Tool)", "Cheap", 3, true},
	81:  {"Tinderbox (Tool)", "Cheap", 1, true},
	82:  {"Tongs (Tool)", "Cheap", 1, true},
	83:  {"Torch (Tool)", "Cheap", 1, true},
	84:  {"Trip Wire (Tool)", "Cheap", 1, true},
	85:  {"Tuning Fork (Tool)", "Cheap", 1, true},
	86:  {"Umbrella (Tool)", "Cheap", 1, true},
	87:  {"Wax (Tool)", "Cheap", 0, true},
	88:  {"Wheelbarrow (Tool)", "Cheap", 0, false},
	89:  {"Whistle (Tool)", "Cheap", 0, true},
	90:  {"Bardiche (Polearm)", "Pricey", 3, true},
	91:  {"Battleaxe (Bladed)", "Pricey", 3, true},
	92:  {"Bec de Corbin (Pricey)", "Pricey", 3, true},
	93:  {"Broadsword (Bladed)", "Pricey", 3, true},
	94:  {"Bow (Ranged)", "Pricey", 2, true},
	95:  {"Chain sickle (Flexible)", "Pricey", 2, true},
	96:  {"Crossbow (Ranged)", "Pricey", 2, true},
	97:  {"Cutlass (Bladed)", "Pricey", 2, true},
	98:  {"Dual Axes (Bladed)", "Pricey", 2, true},
	99:  {"Dual Daggers (Bladed)", "Pricey", 1, true},
	100: {"Falchion (Bladed)", "Pricey", 2, true},
	101: {"Flail (Blunt)", "Pricey", 2, true},
	102: {"Glaive (Bladed)", "Pricey", 3, true},
	103: {"Greataxe (Bladed)", "Pricey", 3, true},
	104: {"Greateclub (Blunt)", "Pricey", 3, true},
	105: {"Greatsword (Bladed)", "Pricey", 3, true},
	106: {"Halberd (Polearm)", "Pricey", 3, true},
	107: {"Khopesh (Bladed)", "Pricey", 2, true},
	108: {"Kukri (Bladed)", "Pricey", 1, true},
	109: {"Lance (Polearm)", "Pricey", 3, true},
	110: {"Longbow (Ranged)", "Pricey", 3, true},
	111: {"Longsword (Bladed)", "Pricey", 2, true},
	112: {"Mace (Blunt)", "Pricey", 2, true},
	113: {"Machete (Bladed)", "Pricey", 2, true},
	114: {"Mancatcher (Polearm)", "Pricey", 3, true},
	115: {"Maul (Blunt)", "Pricey", 3, true},
	116: {"Morning Star (Blunt)", "Pricey", 2, true},
	117: {"Pike (Polearm)", "Pricey", 3, true},
	118: {"Ranseur (Polearm)", "Pricey", 3, true},
	119: {"Rapier (Bladed)", "Pricey", 1, true},
	120: {"Saber (Bladed)", "Pricey", 2, true},
	121: {"Scimitar (Bladed)", "Pricey", 2, true},
	122: {"Scourge (Flexible)", "Pricey", 1, true},
	123: {"Shortbow (Ranged)", "Pricey", 1, true},
	124: {"Throwing Axes (Ranged)", "Pricey", 1, true},
	125: {"Trident (Polearm)", "Pricey", 3, true},
	126: {"Two-Handed Sword (Bladed)", "Pricey", 3, true},
	127: {"War Axe (Bladed)", "Pricey", 2, true},
	128: {"War Pick (Blunt)", "Pricey", 2, true},
	129: {"Warhammer (Blunt)", "Pricey", 3, true},
	130: {"Quilted (Armor)", "Pricey", 1, true},
	131: {"Padded (Armor)", "Pricey", 1, true},
	132: {"Leather (Armor)", "Pricey", 1, true},
	133: {"Ring Mail (Armor)", "Pricey", 2, true},
	134: {"Pauldrons (Armor)", "Pricey", 2, true},
	135: {"Gambeson (Armor)", "Pricey", 1, true},
	136: {"Buckler (Armor)", "Pricey", 1, true},
	137: {"Greaves (Armor)", "Pricey", 1, true},
	138: {"Kite Shield (Armor)", "Pricey", 3, true},
	139: {"Gauntlets (Armor)", "Pricey", 1, true},
	140: {"Jousting Helm (Armor)", "Pricey", 1, true},
	141: {"Kettle Helm (Armor)", "Pricey", 1, true},
	142: {"Goblin Helm (Armor)", "Pricey", 1, true},
	143: {"Chain Mail Coif (Armor)", "Pricey", 1, true},
	144: {"Acid (Tool)", "Pricey", 0, true},
	145: {"Astrolabe (Tool)", "Pricey", 2, true},
	146: {"Blasting Charge (Tool)", "Pricey", 1, true},
	147: {"Caltrops (Tool)", "Pricey", 0, true},
	148: {"Cart (Tool)", "Pricey", 0, false},
	149: {"Chemistry Set (Tool)", "Pricey", 2, true},
	150: {"Compass (Tool)", "Pricey", 1, true},
	151: {"Deck of Huzzah Cards (Tool)", "Pricey", 0, true},
	152: {"Disguise Kit (Tool)", "Pricey", 2, true},
	153: {"Divining Rod (Tool)", "Pricey", 1, true},
	154: {"Explosive Powder (Tool)", "Pricey", 1, true},
	155: {"Grappling Hook (Tool)", "Pricey", 2, true},
	156: {"Lock (Tool)", "Pricey", 1, true},
	157: {"Lockpicks (Tool)", "Pricey", 0, true},
	158: {"Journal (Tool)", "Pricey", 1, true},
	159: {"Crafting Tools (Tools)", "Pricey", 3, true},
	160: {"Medicine (Tool)", "Pricey", 1, true},
	161: {"Parachute (Tool)", "Pricey", 3, true},
	162: {"Perfume (Tool)", "Pricey", 0, true},
	163: {"Raft (Tool)", "Pricey", 0, false},
	164: {"Rapelling Harness (Tool)", "Pricey", 1, true},
	165: {"Saddlebags (Tool)", "Pricey", 0, false},
	166: {"Sail (Tool)", "Pricey", 0, false},
	167: {"Snare (Tool)", "Pricey", 2, true},
	168: {"Strongbox (Tool)", "Pricey", 2, true},
	169: {"Wig (Tool)", "Pricey", 1, true},
	170: {"Anvil (Tool)", "Expensive", 0, false},
	171: {"Bear Skin (Armor)", "Expensive", 2, true},
	172: {"Blasting Machine (Tool)", "Expensive", 1, true},
	173: {"Blunderbuss (Ranged)", "Expensive", 2, true},
	174: {"Breastplate (Armor)", "Expensive", 3, true},
	175: {"Brigandine Mail (Armor)", "Expensive", 2, true},
	176: {"Canoe (Tool)", "Expensive", 0, false},
	177: {"Chain Mail (Armor)", "Expensive", 2, true},
	178: {"Chariot (Tool)", "Expensive", 0, false},
	179: {"Detonator (Tool)", "Expensive", 1, true},
	180: {"Great Helm (Armor)", "Expensive", 2, true},
	181: {"Heater Shield (Armor)", "Expensive", 3, true},
	182: {"Horned Helm (Armor)", "Expensive", 2, true},
	183: {"Lamellar (Armor)", "Expensive", 2, true},
	184: {"Paraglider (Tool)", "Expensive", 3, true},
	185: {"Poison (Tool)", "Expensive", 0, true},
	186: {"Powder Keg (Tool)", "Expensive", 3, true},
	187: {"Scale Mail (Armor)", "Expensive", 2, true},
	188: {"Spitfire (Ranged)", "Expensive", 3, true},
	189: {"Targe (Armor)", "Expensive", 2, true},
	190: {"Telescope (Tool)", "Expensive", 1, true},
	191: {"Timepiece (Tool)", "Expensive", 0, true},
	192: {"Tower Shield (Armor)", "Expensive", 3, true},
	193: {"Toxin (Tool)", "Expensive", 0, true},
	194: {"Hunting Dog", "Expensive", 0, false},
	195: {"Pony", "Really Expensive", 0, false},
	196: {"War Horse", "Really Expensive", 0, false},
	197: {"Boat", "Really Expensive", 0, false},
	198: {"Wagon", "Really Expensive", 0, false},
	199: {"Ballistae", "Really Expensive", 0, false},
	200: {"Broken Flying Machine", "Really Expensive", 0, false}}

func RollEquipment() Equipment {
	roll := dice.Roll(1, 200, 0)
	return EquipmentTable[roll]
}
