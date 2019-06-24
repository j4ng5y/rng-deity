package deitylib

import (
	"github.com/google/uuid"
)

func createDruidDeities() {
	druidDeities := map[string]map[string]interface{}{
		"Aine": {
			"Religion": "Druid",
			"Sex":      "Female",
			"Oversight": []string{
				"Love",
				"Summer",
				"Sovreignty",
			},
		},
		"Amaethon": {
			"Religion": "Druid",
			"Sex":      "Female",
			"Oversight": []string{
				"Agreculture",
				"Animal Husbandry",
			},
		},
		"Arawn": {
			"Religion": "Druid",
			"Sex":      "Male",
			"Oversight": []string{
				"the Dead",
				"Hunting",
				"Revenge",
				"the Underworld",
			},
		},
		"Arianrhod": {
			"Religion": "Druid",
			"Sex":      "Female",
			"Oversight": []string{
				"Beauty",
				"Fertility",
				"Reincarnation",
				"the Sky",
				"Weaving",
				"Enchantment",
			},
		},
		"Blodeuwedd": {
			"Religion": "Druid",
			"Sex":      "Female",
			"Oversight": []string{
				"Flowers",
				"Lunar Mysteries",
				"Wisdom",
			},
		},
		"Brigit": {
			"Religion": "Druid",
			"Sex":      "Female",
			"Oversight": []string{
				"Fire",
				"Healing",
				"Motherhood",
				"Agriculture",
				"Inspiration",
				"Learning",
				"Divination",
				"Poetry",
				"Prophecy",
				"the Forge",
			},
		},
		"Cailleach": {
			"Religion": "Druid",
			"Sex":      "Female",
			"Oversight": []string{
				"Disease",
				"Plague",
				"Sorcery",
			},
		},
		"Cernunnos": {
			"Religion": "Druid",
			"Sex":      "Male",
			"Oversight": []string{
				"Hunting",
				"Animals",
				"Fertility",
				"Warriors",
				"Nature",
				"Commerce",
				"Love",
				"The Underworld",
			},
		},
		"Cerridwen": {
			"Religion": "Druid",
			"Sex":      "Female",
			"Oversight": []string{
				"Enchantment",
				"Death",
				"Initiation",
				"Wisdom",
				"Inspiration",
				"Regeneration",
				"Dark Prophehecy",
				"the Moon",
				"Grain",
			},
		},
		"Dagda": {
			"Religion": "Druid",
			"Sex":      "Male",
			"Oversight": []string{
				"Art",
				"Knowledge",
				"Magic",
				"Music",
				"Prophecy",
				"Prosperity",
				"Regeneration",
				"Fatherhood",
				"Protection",
			},
		},
		"Danu": {
			"Religion": "Druid",
			"Sex":      "Female",
			"Oversight": []string{
				"the Elements",
				"Water",
				"Magic",
				"Wisdom",
				"the Earth",
				"Cattle",
			},
		},
		"Diancecht": {
			"Religion": "Druid",
			"Sex":      "Male",
			"Oversight": []string{
				"Healing",
				"Resurection",
			},
		},
		"Druantia": {
			"Religion": "Druid",
			"Sex":      "Female",
			"Oversight": []string{
				"Fir Trees",
				"Passion",
				"Protection",
				"Knowledge",
				"Creativity",
			},
		},
		"Epona": {
			"Religion": "Druid",
			"Sex":      "Female",
			"Oversight": []string{
				"Healing",
				"Prosperity",
				"Maternity",
			},
		},
		"Goibniu": {
			"Religion": "Druid",
			"Sex":      "Male",
			"Oversight": []string{
				"the Forge",
				"Brewing",
				"Thunder",
			},
		},
		"Gwyddion": {
			"Religion": "Druid",
			"Sex":      "Male",
			"Oversight": []string{
				"Enchantment",
				"Illusion",
				"Magic",
				"Music",
				"Shapeshifting",
				"Learning",
			},
		},
		"Gwynn ap Nudd": {
			"Religion": "Druid",
			"Sex":      "Male",
			"Oversight": []string{
				"War",
				"Death",
				"Fallen Warriors",
				"Hunting",
			},
		},
		"Llyr": {
			"Religion": "Druid",
			"Sex":      "Male",
			"Oversight": []string{
				"the Sea",
				"Water",
				"the Underworld",
			},
		},
		"Lugh": {
			"Religion": "Druid",
			"Sex":      "Male",
			"Oversight": []string{
				"the Sun",
				"Crafts",
				"Art",
				"Healing",
				"Journeys",
				"Prophecy",
			},
		},
		"Manannan": {
			"Religion": "Druid",
			"Sex":      "Male",
			"Oversight": []string{
				"the Sea",
				"Weather",
				"the Underworld",
			},
		},
		"Morrigan": {
			"Religion": "Druid",
			"Sex":      "Female",
			"Oversight": []string{
				"Magic",
				"Prophecy",
				"Revenge",
				"War",
				"Death",
			},
		},
		"Niamh": {
			"Religion": "Druid",
			"Sex":      "Female",
			"Oversight": []string{
				"Light",
			},
		},
		"Nuada": {
			"Religion": "Druid",
			"Sex":      "Male",
			"Oversight": []string{
				"Harpers",
				"Healing",
				"Historians",
				"Magic",
				"Poets",
				"Warfare",
				"Writing",
			},
		},
		"Ogma": {
			"Religion": "Druid",
			"Sex":      "Male",
			"Oversight": []string{
				"Eloquence",
				"Inspiration",
				"Language",
				"Magic",
				"Music",
				"Strength",
				"Poets",
				"Writers",
			},
		},
		"Taliesin": {
			"Religion": "Druid",
			"Sex":      "Male",
			"Oversight": []string{
				"Magic",
				"Music",
				"Poetry",
				"Wisdom",
				"Writing",
			},
		},
	}

	for k, v := range druidDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
