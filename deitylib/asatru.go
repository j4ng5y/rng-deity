package deitylib

import (
	"github.com/google/uuid"
)

func createAsatruDeities() {
	asartuDeities := map[string]map[string]interface{}{
		"Odin": {
			"Religion": "Asartu",
			"Sex":      "Male",
			"Oversight": []string{
				"Knowledge",
				"Wisdom",
			},
		},
		"Thor": {
			"Religion": "Asartu",
			"Sex":      "Male",
			"Oversight": []string{
				"Thunder",
				"Strength",
				"Crafts",
			},
		},
		"Balder": {
			"Religion": "Asartu",
			"Sex":      "Male",
			"Oversight": []string{
				"Eloquence",
				"Friendship",
				"Good Looks",
			},
		},
		"Tyr": {
			"Religion": "Asartu",
			"Sex":      "Male",
			"Oversight": []string{
				"Law",
				"Order",
				"Justice",
			},
		},
		"Heimdall": {
			"Religion": "Asartu",
			"Sex":      "Male",
			"Oversight": []string{
				"Vision",
				"Hearing",
			},
		},
		"Aegir": {
			"Religion": "Asartu",
			"Sex":      "Male",
			"Oversight": []string{
				"The Sea",
			},
		},
		"Bragl": {
			"Religion": "Asartu",
			"Sex":      "Male",
			"Oversight": []string{
				"Poetry",
				"Eloquence",
			},
		},
		"Forsetti": {
			"Religion": "Asartu",
			"Sex":      "Male",
			"Oversight": []string{
				"Law",
				"Justice",
			},
		},
		"Njord": {
			"Religion": "Asartu",
			"Sex":      "Male",
			"Oversight": []string{
				"Winds",
				"The Sea",
				"Harvests",
			},
		},
		"Freyr": {
			"Religion": "Asartu",
			"Sex":      "Male",
			"Oversight": []string{
				"Fertility",
			},
		},
		"Vall": {
			"Religion": "Asartu",
			"Sex":      "Male",
			"Oversight": []string{
				"Vengence",
				"Retribution",
			},
		},
		"Loki": {
			"Religion": "Asartu",
			"Sex":      "Male",
			"Oversight": []string{
				"Mischief",
				"Fraud",
			},
		},
		"Manni": {
			"Religion": "Asartu",
			"Sex":      "Male",
			"Oversight": []string{
				"The Moon",
				"Darkness",
			},
		},
		"Frigga": {
			"Religion": "Asartu",
			"Sex":      "Female",
			"Oversight": []string{
				"Knowledge",
				"Wisdom",
				"Marriage",
				"Cloth Weaving",
				"Health",
				"Home",
			},
		},
		"Freya": {
			"Religion": "Asartu",
			"Sex":      "Female",
			"Oversight": []string{
				"Beauty",
				"Strength",
				"Love",
				"Fertility",
				"Magic",
				"War",
				"Death",
			},
		},
		"Idunna": {
			"Religion": "Asartu",
			"Sex":      "Female",
			"Oversight": []string{
				"Health",
				"Harvest",
				"Eternal Life",
			},
		},
		"Hel": {
			"Religion": "Asartu",
			"Sex":      "Female",
			"Oversight": []string{
				"The Underworld",
				"Death",
			},
		},
		"Sif": {
			"Religion": "Asartu",
			"Sex":      "Female",
			"Oversight": []string{
				"Harvest",
			},
		},
		"Ran": {
			"Religion": "Asartu",
			"Sex":      "Female",
			"Oversight": []string{
				"The Sea",
			},
		},
		"Eir": {
			"Religion": "Asartu",
			"Sex":      "Female",
			"Oversight": []string{
				"Healing",
			},
		},
		"Sunna": {
			"Religion": "Asartu",
			"Sex":      "Female",
			"Oversight": []string{
				"The Sun",
			},
		},
		"Nanna": {
			"Religion": "Asartu",
			"Sex":      "Female",
			"Oversight": []string{
				"Gifts",
			},
		},
		"Ostara": {
			"Religion": "Asartu",
			"Sex":      "Female",
			"Oversight": []string{
				"Dawn",
				"Rebirth",
				"Spring",
			},
		},
		"Nerthus": {
			"Religion": "Asartu",
			"Sex":      "Female",
			"Oversight": []string{
				"The Earth",
			},
		},
		"Skadi": {
			"Religion": "Asartu",
			"Sex":      "Female",
			"Oversight": []string{
				"Winter",
			},
		},
	}

	for k, v := range asartuDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
