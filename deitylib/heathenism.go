package deitylib

import (
	"github.com/google/uuid"
)

func createHeathenismDeities() {
	heathenismDeities := map[string]map[string]interface{}{
		"Odin": {
			"Religion": "Heathenistic",
			"Sex":      "Male",
			"Oversight": []string{
				"Knowledge",
				"Wisdom",
			},
		},
		"Thor": {
			"Religion": "Heathenistic",
			"Sex":      "Male",
			"Oversight": []string{
				"Thunder",
				"Strength",
				"Crafts",
			},
		},
		"Balder": {
			"Religion": "Heathenistic",
			"Sex":      "Male",
			"Oversight": []string{
				"Eloquence",
				"Friendship",
				"Good Looks",
			},
		},
		"Tyr": {
			"Religion": "Heathenistic",
			"Sex":      "Male",
			"Oversight": []string{
				"Law",
				"Order",
				"Justice",
			},
		},
		"Heimdall": {
			"Religion": "Heathenistic",
			"Sex":      "Male",
			"Oversight": []string{
				"Vision",
				"Hearing",
			},
		},
		"Aegir": {
			"Religion": "Heathenistic",
			"Sex":      "Male",
			"Oversight": []string{
				"The Sea",
			},
		},
		"Bragl": {
			"Religion": "Heathenistic",
			"Sex":      "Male",
			"Oversight": []string{
				"Poetry",
				"Eloquence",
			},
		},
		"Forsetti": {
			"Religion": "Heathenistic",
			"Sex":      "Male",
			"Oversight": []string{
				"Law",
				"Justice",
			},
		},
		"Njord": {
			"Religion": "Heathenistic",
			"Sex":      "Male",
			"Oversight": []string{
				"Winds",
				"The Sea",
				"Harvests",
			},
		},
		"Freyr": {
			"Religion": "Heathenistic",
			"Sex":      "Male",
			"Oversight": []string{
				"Fertility",
			},
		},
		"Vall": {
			"Religion": "Heathenistic",
			"Sex":      "Male",
			"Oversight": []string{
				"Vengence",
				"Retribution",
			},
		},
		"Loki": {
			"Religion": "Heathenistic",
			"Sex":      "Male",
			"Oversight": []string{
				"Mischief",
				"Fraud",
			},
		},
		"Manni": {
			"Religion": "Heathenistic",
			"Sex":      "Male",
			"Oversight": []string{
				"The Moon",
				"Darkness",
			},
		},
		"Frigga": {
			"Religion": "Heathenistic",
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
			"Religion": "Heathenistic",
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
			"Religion": "Heathenistic",
			"Sex":      "Female",
			"Oversight": []string{
				"Health",
				"Harvest",
				"Eternal Life",
			},
		},
		"Hel": {
			"Religion": "Heathenistic",
			"Sex":      "Female",
			"Oversight": []string{
				"The Underworld",
				"Death",
			},
		},
		"Sif": {
			"Religion": "Heathenistic",
			"Sex":      "Female",
			"Oversight": []string{
				"Harvest",
			},
		},
		"Ran": {
			"Religion": "Heathenistic",
			"Sex":      "Female",
			"Oversight": []string{
				"The Sea",
			},
		},
		"Eir": {
			"Religion": "Heathenistic",
			"Sex":      "Female",
			"Oversight": []string{
				"Healing",
			},
		},
		"Sunna": {
			"Religion": "Heathenistic",
			"Sex":      "Female",
			"Oversight": []string{
				"The Sun",
			},
		},
		"Nanna": {
			"Religion": "Heathenistic",
			"Sex":      "Female",
			"Oversight": []string{
				"Gifts",
			},
		},
		"Ostara": {
			"Religion": "Heathenistic",
			"Sex":      "Female",
			"Oversight": []string{
				"Dawn",
				"Rebirth",
				"Spring",
			},
		},
		"Nerthus": {
			"Religion": "Heathenistic",
			"Sex":      "Female",
			"Oversight": []string{
				"The Earth",
			},
		},
		"Skadi": {
			"Religion": "Heathenistic",
			"Sex":      "Female",
			"Oversight": []string{
				"Winter",
			},
		},
	}
	

	for k, v := range heathenismDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
