package deitylib

import (
	"github.com/google/uuid"
)

func createGreekDeities() {
	greekDeities := map[string]map[string]interface{}{
		"Achelous": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"the Achelous River",
			},
		},
		"Aeolus": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Wind",
				"the Air",
			},
		},
		"Aether": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"the Upper Air",
				"Light",
				"the Atmosphere",
				"Space",
				"Heaven",
			},
		},
		"Alastor": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Family Feuds",
				"Vengence",
			},
		},
		"Apollo": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Music",
				"Poetry",
				"Art",
				"Oracles",
				"Archery",
				"Plagues",
				"Medicine",
				"the Sun",
				"Light",
				"Knowledge",
			},
		},
		"Ares": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"War",
			},
		},
		"Aristaeus": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Animal Husbandry",
				"Bee Keeping",
				"Fruit Trees",
			},
		},
		"Asclepius": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Medicine",
				"Health",
				"Healing",
				"Rejuvenation",
				"Physicians",
			},
		},
		"Atlas": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Astronomy",
			},
		},
		"Attis": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Vegetation",
				"Fruit",
				"Rebirth",
			},
		},
		"Boreas": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Wind",
				"Winter",
			},
		},
		"Caerus": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Opportunity",
				"Luck",
			},
		},
		"Castor": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Twins",
			},
		},
		"Cerus": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Bulls",
			},
		},
		"Chaos": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Nothingness",
				"Creation",
			},
		},
		"Charon": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Ferrymen",
			},
		},
		"Cronos": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Time",
			},
		},
		"Crios": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"the Stars",
			},
		},
		"Cronus": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Agriculture",
			},
		},
		"Dinlas": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"the City of Lamark",
			},
		},
		"Deimos": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Dread",
				"Terror",
			},
		},
		"Dionysus": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Grapes",
				"Winemaking",
				"Wine",
				"Ritual Madness",
				"Religion Ecstasy",
				"Theatre",
			},
		},
		"Erebus": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Darkness",
			},
		},
		"Eros": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Sexual Desire",
				"Attraction",
				"Love",
				"Procreation",
			},
		},
		"Eurus": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Wind",
			},
		},
		"Galucus": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"the Sea",
			},
		},
		"Hades": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"the Dead",
				"the Underworld",
				"Riches",
			},
		},
		"Helios": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"the Sun",
			},
		},
		"Hephaestus": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Fire",
				"Metalworking",
				"Stone Masonry",
				"Forges",
				"Sculpture",
			},
		},
		"Heracles": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Heros",
				"Sports",
				"Athletes",
				"Health",
				"Agriculture",
				"Fertility",
				"Trade",
				"Oracles",
				"the Protection of Mankind",
			},
		},
		"Hermes": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Trade",
				"Theives",
				"Travelers",
				"Sports",
				"Athletes",
				"Borders",
			},
		},
		"Hesperus": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"the Evening Star",
			},
		},
		"Hymenaios": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Marriage",
				"Inspiring Feats",
				"Song",
			},
		},
		"Hypnos": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Sleep",
			},
		},
		"Kratos": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Strength",
				"Power",
			},
		},
		"Momus": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Satire",
				"Mockery",
				"Censure",
				"Writers",
				"Poets",
				"Evil-Spirited Blame",
				"Unfair Criticism",
			},
		},
		"Morpheus": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Dreams",
				"Sleep",
			},
		},
		"Nereus": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"the Sea",
			},
		},
		"Notus": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Wind",
			},
		},
		"Oceanus": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Oceans",
			},
		},
		"Oneiroi": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Dreams",
			},
		},
		"Paean": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Physicians",
			},
		},
		"Pallas": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Warcraft",
				"Spring",
			},
		},
		"Pan": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Nature",
				"the Wild",
				"Shepherds",
				"Flocks",
				"Goats",
				"Mountain Wilds",
				"Sexuality",
			},
		},
		"Phosphorus": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"the Morning Star",
			},
		},
		"Plutus": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Wealth",
			},
		},
		"Pollux": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Twins",
			},
		},
		"Pontus": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"the Deep Sea",
			},
		},
		"Poseidon": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"the Sea",
				"Earthquakes",
				"Storms",
				"Horses",
			},
		},
		"Priapus": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Fertility",
				"Flocks",
				"Furit Plants",
				"Bees",
				"Gardens",
			},
		},
		"Pricus": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Sea Goats",
			},
		},
		"Prometheus": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Forethought",
				"Counsel",
				"Creation",
			},
		},
		"Tartarus": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"the Abyss",
				"the Underworld",
			},
		},
		"Thanatos": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Death",
			},
		},
		"Triton": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"the Sea",
			},
		},
		"Typhon": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Monsters",
				"Storms",
				"Volcanoes",
			},
		},
		"Uranus": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"the Sky",
				"the Heavens",
			},
		},
		"Zelus": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Dedication",
				"Emulation",
				"Rivalry",
				"Envy",
				"Jealousy",
				"Zeal",
			},
		},
		"Zephyrus": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Wind",
			},
		},
		"Zeus": {
			"Religion": "Greek",
			"Sex":      "Male",
			"Oversight": []string{
				"Lightning",
				"the Sky",
				"Thunder",
				"Law",
				"Order",
				"Justice",
			},
		},
		"Achelois": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"the Moon",
			},
		},
		"Alectrona": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"the Sun",
				"Morning",
			},
		},
		"Amphitrite": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"the Sea",
			},
		},
		"Antheia": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Gardens",
				"Flowers",
				"Swamps",
				"Marshes",
			},
		},
		"Apate": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Gardens",
				"Flowers",
				"Swamps",
				"Marshes",
			},
		},
		"Aphaea": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"the island of Aegina",
			},
		},
		"Aphrodite": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Love",
				"Beauty",
			},
		},
		"Artemis": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Hunting",
			},
		},
		"Astraea": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Stars",
			},
		},
		"Ate": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Mischeif",
				"Delusion",
				"Ruin",
				"Folly",
			},
		},
		"Athena": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Wisdom",
				"Peotry",
				"Art",
				"War",
			},
		},
		"Atropos": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Fate",
				"Destiny",
				"Death",
			},
		},
		"Bia": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Force",
				"Raw Energy",
			},
		},
		"Brizo": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Mariners",
				"Sailors",
				"Fishermen",
			},
		},
		"Calliope": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Epic Poetry",
			},
		},
		"Calypso": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Nymphs",
			},
		},
		"Ceto": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Sea Monsters",
			},
		},
		"Circe": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Magic",
			},
		},
		"Clio": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"History",
			},
		},
		"Clotho": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Fate",
				"Destiny",
				"Life",
			},
		},
		"Cybele": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Caverns",
				"Mountains",
				"Nature",
				"Wild Animals",
			},
		},
		"Demeter": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Agriculture",
				"Fertility",
				"Law",
				"the Harvest",
			},
		},
		"Doris": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"the Nereids",
			},
		},
		"Eileithyia": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Childbirth",
			},
		},
		"Elpis": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Hope",
			},
		},
		"Enyo": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"War",
				"Destruction",
			},
		},
		"Eos": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"the Dawn",
			},
		},
		"Erato": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Love Poetry",
				"Erotic Poetry",
			},
		},
		"Eris": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Chaos",
				"Strife",
				"Discord",
			},
		},
		"Euterpe": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Music",
			},
		},
		"Gaia": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"the Earth",
			},
		},
		"Harmonia": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Harmony",
				"Concord",
			},
		},
		"Hebe": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Eternal Youth",
			},
		},
		"Hecate": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Magic",
				"Crossroads",
				"the Moon",
				"Ghosts",
				"Witchcraft",
				"Necromancy",
			},
		},
		"Hemera": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"the Day",
				"Daytime",
				"Daylight",
			},
		},
		"Hera": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Goddesses",
				"Women",
				"Marriage",
			},
		},
		"Hestia": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"the Hearth",
				"the Home",
				"Architecture",
				"Domesitcity",
				"Family",
				"the State",
			},
		},
		"Hygea": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Good Health",
				"Cleanliness",
				"Sanitation",
			},
		},
		"Iris": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Rainbows",
				"the Sea",
				"the Sky",
			},
		},
		"Keres": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Night",
			},
		},
		"Kotys": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Lascivious Celebrations",
			},
		},
		"Lachesis": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Fate",
				"Destiny",
			},
		},
		"Maia": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Fields",
			},
		},
		"Mania": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Insanity",
				"Madness",
				"Frenzy",
				"the Dead",
			},
		},
		"Melpomene": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Singing",
				"Tragedy",
			},
		},
		"Metis": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Wisdom",
			},
		},
		"Nemesis": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Retribution",
				"Vengence",
			},
		},
		"Nike": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Victory",
			},
		},
		"Nyx": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"the Night",
			},
		},
		"Peitho": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Persuasion",
				"Seduction",
			},
		},
		"Persephone": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Vegetation",
				"Spring",
				"the Underworld",
			},
		},
		"Pheme": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Fame",
				"Gossip",
				"Renown",
			},
		},
		"Polyhymnia": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Poetry",
				"Hymn",
				"Dance",
				"Eloquence",
				"Agriculture",
				"Geometry",
				"Pantomime",
			},
		},
		"Rhea": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Nature",
			},
		},
		"Selene": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"the Moon",
				"Vampires",
			},
		},
		"Styx": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"the river Styx",
			},
		},
		"Terpsichore": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Dance",
				"Chorus",
			},
		},
		"Thalia": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Comedy",
			},
		},
		"Themis": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Order",
				"Law",
				"Customs",
			},
		},
		"Thetis": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Water",
			},
		},
		"Tyche": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Prosperity",
				"Fortune",
			},
		},
		"Urania": {
			"Religion": "Greek",
			"Sex":      "Female",
			"Oversight": []string{
				"Astrology",
				"Astronomy",
			},
		},
	}

	for k, v := range greekDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
