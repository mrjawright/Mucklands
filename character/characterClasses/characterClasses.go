package characterClasses

var ClassNames = [...]string{"Bard", "Dungeoneer", "Gnome", "Knight-errant", "Loyal Chum", "Rascal"}

type ClassStory struct {
	Backstory     []string
	Ideals        []string
	Flaws         []string
	PersonalQuest []string
	Relationships []string
}

type CharacterClass struct {
	Name            string
	CourageModifier string
	Dread           string
	Perks           []string
	Abilities       []string
	Story           ClassStory
}

var CharacterClasses = map[string]CharacterClass{
	"Bard": {Name: "Bard", CourageModifier: "12", Dread: "D4",
		Perks: []string{"The Bard's Instrument:\nSoothing, allows the bard to heal allies. Twice per session, roll 1d6 to heal that much courage to an ally.\n" +
			"Fighting Music, The bard's music can also be epic and thrilling, making for a great soundtrack to a fight. During combat, the Bard's instrument is a ranged weapon that deals dread.\n",
			"Choose your instrument:\nString Instrument, Your instrument does +1 Healing.\n" +
				"Percussion Instrument, Your instrument does +1 Dread\n" +
				"Brass/Woodwind Instrument, Once per session, gain Advantage on an Inspire check\n",
			"Choose your talent:\n Musician: when you heal someone roll the healing dice with  advantage.\n" +
				"Storyteller: Once per session, when you tell a story, a party member gets an extra Quest Point for the session.\n" +
				"Performer: Choose two Proficiencies: Acting, Impersonating, Dancing, Acrobatics"},
		Abilities: []string{"Narrator: Roll+Inspire Once per session, narrate a desired outcome of a specific action beginning with the phrase \"And then...\".You receive a +1 bonus for rhyming.",
			"Little Ditty: Once per session, Roll+Inspire and heal that much Courage to all allies outside of a Conflict."},
		Story: ClassStory{Backstory: []string{"A traveling troupe of bards visited our hometown each year.",
			"A mysterious wanderer in the wilderness with a golden voice made a lasting impression.",
			"I found an old, abandoned instrument in a pile of rubbish.",
			"My parents were musicians and instilled a love of music in me at a young age.",
			"I snuck into a play whne I was a child and dreamed of sharing the stage.",
			"It wasn't easy. Music was outlawed where I'm from, but I would  play each night in secret."},
			Ideals: []string{"Learn or tell an old story.", "Give someone hope.", "Make a new fan.",
				"Execute a convoluted plan.", "Defuse a tense situation with humor.", "Encourage someone's creativity."},
			Flaws: []string{"Get distracted at an inopportune moment.", "Let jealousy get the better of me.",
				"Fall in love with an NPC", "Insult an important NPC", "Blurt out a secret to someone",
				"Put my foot in my mouth."},
			PersonalQuest: []string{"Discover the truth about _____ and share it with the world.",
				"Inspire the people of _____ to fight back against _____.",
				"Journey to _____ and learn its forgotten tales.",
				"Tell the story of  _____ for all the world to hear about.",
				"Find _____, which was thought to be lost to time",
				"Finally escape my former life of _____, and become a hero."},
			Relationships: []string{"_____ is the only one who truly gets my genius",
				"_____ sure is great in a fight, but needs to learn how to let loose every once in a while.",
				"_____ can be critical and rude, but always listens quietly to my stories.",
				"_____ and I have a lot in common...we both enjoy stirring up trouble.",
				"_____ can't carry a tune to save their life, but they saved mine when it counted.",
				"_____ is always there to offer wise advice, but sometimes it gets a little old.",
				"_____ I can tell that _____ doesn't like me much but I'll charm them into friendship!",
				"_____  and I used to be more than friends, but I swear I don't like them like that anymore."}}},
	"Dungeoneer": {Name: "Dungeoneer", CourageModifier: "13", Dread: "D8",
		Perks: []string{"Choose your underling:\n You have a named NPC underling that follows you around and does stuff for you." +
			"The have +1 in one stat and -3 in everything else. They have 1 Courage and won't participate in Combat." +
			"If your underling takes Dread, they go \"on break\" for the rest of the session.\nBrawny: +1 Vigor\nSneaky: +1 Knack\nBrainy: +1Knowhow",
			"Choose your gear:\nEvery Dungeoneer packs the proper gear for adventuring and exploration. It's always best to be prepared.\n" +
				"Maps: You get +1 to Travel Checks\nMetal Detector: +20%% Treasure Hunting\nDelving Tools: At the start of the session, you always have delving gear that doesn't count against your inventory:\n" +
				"Rope, Torch, Ten-foot Pole, Shovel, Pickax"},
		Abilities: []string{"Boggle Crew: Roll+Inspide or Intimidate Once per session, you can boss them around for various purposes.",
			"Reconnoiter: Roll+Realms Once per session, make up a fact or rumor about a point of interest, or a group of people at a point of interest."},
		Story: ClassStory{Backstory: []string{"I want to make a name for myself as an adventurer, maybe even make the Dungeoneer Magazine Hall of Fame one day.",
			"What better way to discover last treasure and hit it rich? Sure beats gettgina real job!",
			"I worked for Subterranean Pits and Lairs, LLC for a time but got fed up with the rat race and decided to strike out on my own.",
			"There are many amazing sights to behold in the Land of Eem, and I want to see them with my own eyes.",
			"There is so much to be learned from the relics of the past, more than any plain old history book could teach me.",
			"I enrolled on Dungeoneer Academy whe I was ten years old, to master the adventure trade, and one day pass down my knowledge to a new generation."},
			Ideals: []string{"Solve a problem with ingenuity.", "Explore a dungeon completely.",
				"Put my body on the line for a party member.", "Negotiate a deal.", "Find a relic or treasure.",
				"Interact with a rare creature"},
			Flaws: []string{"Throw caution to the wind.", "Let greed get the better of me.", "Follow my curiosity at any expense.",
				"Make a rival out of someone.", "Drastically overestimate my abilities.", "Refuse to follow orders."},
			PersonalQuest: []string{"Journey to _____ and plant a flag in my name.", "Learn the forgotten knowledge of _____",
				"Find the lost treasure of _____", "Discover the long lost cuty of _____", "Make first contact with _____ and establish peace and trade",
				"Build a successful dungeoneering company from the ground up and call it _____"},
			Relationships: []string{"Even though we're often at odds, I actually really admire _____",
				"_____ thinks they're the leader of the group, but it's obviously me.",
				"_____ is a liability, and I'm convinced they'll be the death of me.",
				"_____ knew me when I was an awkward teenager at Dungeoneer Academy, and still treats me like a kid.",
				"I basically consider _____ one of my underlings.",
				"_____ owes me big time for saving them in the past.",
				"I would risk my neck to save _____ from any danger, and I expect the same in return.",
				"I hired _____ for a dangerous adventure a while back, and now we're best buds."}}},
	"Gnome": {Name: "Gnome", CourageModifier: "14", Dread: "D8",
		Perks: []string{"Choose your hat:\nSpring Hat: (Green, Yellow, Pink, or Light Blue) Once per session, mend a mundane item." +
			"Autumn Hat: (Brown, Gold, Orange, or Red) Once per session, change the color of something." +
			"Winter Hat: (Dark Blue, Purple, Grey or White) Once per session, warm something up by touch.",
			"Choose your bane:\nGnomes deal 1D12 Dread to their chosen Bane on an attack. Choose one creature to be your bane:\n"},
		Abilities: []string{"Chronicler: Roll+Lore Once per session, create an historical fact, or any but of ancient knowledge, or trivia.",
			"Magic Feet: Roll+Nimbleness Once per session, reroll a nimbleness check or performa spectacular acrobatic feat beyond the capabilities of normal creatures."},
		Story: ClassStory{Backstory: []string{"I was finishing my multi-volume manuscript on an esoteric subject.",
			"I was replating trees in the Used T'be Forest, but the task became impossible.",
			"I was searching for a reliquary of loast treasures, but the trail went cold.",
			"I was training to serve witht he gnomish Rainbow Brigade, but found another calling.",
			"I was living peacefully in the Dingledell when news came of troubled times ahead.",
			"I was watching over a small swath of land, tending to the plants and animals."},
			Ideals: []string{"Bring out the good in someone.",
				"Help someone lost of hurt.",
				"Defeat an evil creature or fiend.",
				"Learn a piece of ancient lore.",
				"Inspire someone to tur over a new leaf.",
				"Risk your life to save an animal or preserve nature."},
			Flaws: []string{"Tell the truth when it is very inconvenient.",
				"Put your trust in someone questionable.",
				"Give more than you can afford, to your own detriment.",
				"Follow the pursuit of knowledge at any expense.",
				"Poke your nose where it doesn't belong.",
				"Act \"holier than thou\" to someone."},
			PersonalQuest: []string{"Defend _____ from certain doom at the hands of ____",
				"Learn the forgotten knowledge of _____ and share it with the world",
				"Banish and their evil from the Land of Eem.",
				"Build a new _____ for the people without one.",
				"Help the _____ overcome their enormous challenge.",
				"Return _____ to a state of peace and prosperity."},
			Relationships: []string{"_____ and I are bound by a Friendship Puff, which means we're friends for life no matter what.",
				"_____ always relies on me to bailmthem out of trouble, but they need to be more cautious.",
				"I have watched _____ grow up from afar, and now it is my time to mentor them.",
				"After they tried to steal from me, I invited _____ to a hot meal and place to rest.",
				"There is a darkness cast over _____, and I am to light their way.",
				"_____ and I often while away the hours discussing literature, tea, and daffodils.",
				"_____ is a rapscallion, but they have earned my trust so far.",
				"_____ believes there is no good left in the world, but I will prove them wrong."}}},
	"Knight-errant": {Name: "Knight-errant", CourageModifier: "15", Dread: "D10",
		Perks: []string{"Choose your equipment:\nMagnificent Sword: [trusty] can't be fumbled or accidentally dropped\n" +
			"Magnificent Breastplate: [sturdy] 1 Block; cannot be broken\n" +
			"Magnificent Banner: [lucky] Once per session, reroll an Inspire Check\n",
			"Choose your steed: Horse: Swiftest. Once per session you can charge at a target from Faraway to deal +1 Dread per level on a successful hit.\n" +
				"Bogril Tortoise: floats. While mounted you get +1 Block and your tortoise provides +12 Inventory\n" +
				"Zozo Bird: Can leap upto 20 feet in any direction with a rider and can safely glide from heights."},
		Abilities: []string{"Wayfarer: Roll+Realms Once per session make something up about a place, landmark, or point of interest in the Land of Eem",
			"Inspiring Orders: Roll+Inspire Twice per session, on a 6+, you may choose to do one of the following:\n" +
				"Give +1 to all party members' rolls during one phase of Conflict\n" +
				"Give +2 to an ally's Check before it is rolled\n" +
				"Heal 1d6 Courage to an ally"},
		Story: ClassStory{Backstory: []string{"My house has sworn fealty to our order for generations. I am honor bound.",
			"I defended my town from bandits an took up the sword from then on to protect others.",
			"I happened upon a dying, old knight and was inspired to continue their quest.",
			"I am a strong fighter and was chosen by my village to represent them in the great conflict to come.",
			"I never wanted to be a knight, but was thrust into the responsibility by a promise.",
			"My mother an father told me stories of chivalrous knights as a child, and I knew then that the world needed heroes."},
			Ideals: []string{"Rescue someone from danger.", "Make a promise and honor my word.",
				"Inspire someone with my heroic deeds.", "Stand my groud against difficult odds.",
				"Broker peace between two sides.", "Defeat an evil creature or fiend."},
			Flaws: []string{"Never back down from a challenge.", "Muscle through a delicate situation.",
				"Resort to violence before I have to.", "Destroy something indiscriminately",
				"Throw caution to the wind.", "Refuse to rest even when it's prudent."},
			PersonalQuest: []string{"Find the lost relic of _____ and return it to _____.",
				"Defeat the dreaded _____ once and for all.",
				"Rescue _____ from perilous danger and uncertainty.",
				"Avenge _____ and prove their sacrifice was not in vain.",
				"Discover the truth about _____ and share it with the world.",
				"Restore _____ to its former glory."},
			Relationships: []string{"I first met _____ after saving them from a pack of hungry weorgs, and not they feel indebted to me.",
				"_____ wants to be my squire, but I don't have  time to babysit someone.",
				"I trust _____ with my life, and there is nothing that I wouldn't do for them.",
				"I respect _____ , but we often butt heads over the littlest things.",
				"I have a feeling that _____ has ulterior motives and I'm not sure i should trust them.",
				"_____ is a bit of a know-it-all, but hads never let me down.",
				"_____ and I crossed paths once as enemies, but now travel as companions.",
				"I worry that _____ has gone astray, so I'lll do my best to keep them on the straight and narrow."}}},
	"Loyal Chum": {Name: "Loyal Chum", CourageModifier: "13", Dread: "D6",
		Perks: []string{"Choose your best chum from the party. Choose one Best Chum Perk:\nLighten the Mood: Say something funny or kind to heal your best chum for 1d12\n" +
			"Take the Hit: Take Dread for your best Chum\n" +
			"Share the Load: Suffer the consequences of your Best Chum's failed action.",
			"Choose your handicraft:\nCooking: Party members heal a D10 (instead of D6) when you make camp, eat food and rest.\n" +
				"Building: You can reroll your Crafting Checks\n" +
				"Alchemy: Everything you craft with Alchemy has one extra use than Normal"},
		Abilities: []string{"Old Chums: Roll+Charm Once per session you can make up an old friend that you've run into who can help you or give advice.\n" +
			"Lend a Hand: Once per session, if an ally fails a roll you can attempt to make the Check instead and replace the ally's result.\n" +
			"Also, you roll with Advantage whenever you attempt to catch someone's fall, pull them from danger, or anything similar"},
		Story: ClassStory{Backstory: []string{"I rejected the pressure to take over my family's farm, and took to the road instead, to see my own fortunes.\n" +
			"I was run out of town by the local riffraff, and someday I'll return to run them out instead.\n" +
			"I plied my trade for years working as an apprentice, but the allure of an adventurous life was too strong.\n" +
			"I was chosen to venture into the wide world and save my poor village from prverty.\n" +
			"My peaceful village was destroyed by fiends and the last of us had nowhere to go.\n" +
			"I was wrongfully accused of a crime and escaped a perilous prison by the skin of my teeth."},
			Ideals: []string{"Make a new friend.", "Risk my life for a party member.",
				"Volunteer to go head first into danger.", "Deescalate a fight when violence is imminent.",
				"Take on someone else's burden.", "Stick up for someone weak."},
			Flaws: []string{"Say something blunt even when it's inconvenient.", "Fall for someone's lies or tricks.",
				"Share too much information with an adversary.", "Poke my nose where it doesn't belong.",
				"Fumble or trip at a bad time.", "Run away like a coward."},
			PersonalQuest: []string{"Prevent _____ from falling into the wrong hands.", "Save _____ from capture, even though it seems impossible.",
				"Follow _____ into danger, and take on their Personal Quest.", "Destroy _____ and end a centuries old curse on my family.",
				"Seek justice for _____ and return home as a hero.", "Retreive _____ and bring it back to its rightful owner."},
			Relationships: []string{"_____ and I have know each other since childhood, and nothing would get between us.",
				"_____ is almost certainly stealing from me, but always shares their food.",
				"_____ is a big grump, but I'm going to do my darndest to make them my friend.",
				"The night we met, _____ and I talked for hours and ended up falling asleep on each other's shoulders.",
				"_____ isn't the person I expected them to be, but they saved my life.",
				"I would die for _____, but I don't think the feeling is mutual.",
				"When we're alone, _____ seems to like me, but then makes fun of me in front of the others.",
				"_____ is too proud to ask for help, so I'll give it without thanks."}}},
	"Rascal": {Name: "Rascal", CourageModifier: "12", Dread: "D6",
		Perks: []string{"Choose your style:\nFootpad: You always have lockpicks on hand and gain Advantage when picking locks.\n" +
			"Charlatan: You have an elaborate alias, complete with a disguise kit that makes you look like another species.\n" +
			"Ruffian: Roll+Might Once per session, knock an unsuspecting Goon or Bruiser unconscious.",
			"Choose your hustle:\nFibbing: Once per session, you gain Advantage on a Trickery Check.\n" +
				"Gambling: Once per session, you can choose to reroll a Success with a Twist, but if the result is lower, the effects are worse than normal." +
				"Pickpocketing: Once per session, roll a Mundane Item."},
		Abilities: []string{"Disappearing Act: Roll+Sneak Once per session, hide in plain sight or disappear from view even when it might seem impossible.\n" +
			"Sticky Fingers: Roll+Nimbleness Once per session, pickpocket an NPC and create a Mundane Item that you have just stolen."},
		Story: ClassStory{Backstory: []string{"I was an orphan, picking pockets and polishing shoes in the slums.",
			"Toiling away in the mines, just like everyone else I grew up with but that life is behind me now.",
			"Unfortunately, if it weren't for the gangs and bandit crews, I'd probably be a goner. Too bad once you're in with them, they never let you out.",
			"Snooping around the city and selling secrets to the highest bidder.",
			"I was forced to work two jobs for corporate goblin goons, until one day I looted their payroll and skipped town, never to look back.",
			"I just got out of the slammer for stealing a loaf of bread when i was only a kid. Now, I do my best to give back to the folks that really need it."},
			Ideals: []string{"Steal something valuable", "Talk your way out of trouble", "Share precious loot with someone in need.",
				"Stick up for someone weak.", "Get the jump on someone.", "Learn a secret"},
			Flaws: []string{"Snitch on another person.", "Run away like a coward.", "Take a bet agains the odds",
				"Lie when you know you shouldn't.", "Take the road less travelled", "Hold out secrets or loot on an NPC"},
			PersonalQuest: []string{"Rob the dastartdly _____ and give back to the downtrodden who need it.",
				"Become my own boss and build a _____ empire.", "Payoff my gigantic debt to _____", "Provide _____ with what they need to get by.",
				"Prevent _____ from getting too powerful and oppressing the weak.", "Steal the valuable and heavily guarded _____ and make a fortune."},
			Relationships: []string{"_____ is good people, but they gotta learn the world ain't all rainbows and unicorns.",
				"_____ knows my darkest secrets, so I need to learn theirs.",
				"_____ ratted me out in the past, and I'm still not sure that I'm over it.",
				"_____ gets on my nerves so much I don't know whether to punch 'em or kiss 'em.",
				"_____ is a fool, but so am I. So, we'll be fools together.",
				"_____ thinks I still owe them, but when will enough be enough?",
				"_____ and I used to be as think as theives, but now we don't talk all that much, and I know know why.",
				"_____ is kind, and brave, and true...so why are they friends with someone like me?"}}}}
