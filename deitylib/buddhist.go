package deitylib

import (
	"github.com/google/uuid"
)

func createBuddhistDeities() {
	buddhistDeities := map[string]map[string]interface{}{
		"Vairocana": {
			"Religion": "Buddhist",
			"Sex":      "Male",
			"Oversight": []string{
				"Light",
			},
		},
		"Aksobhya": {
			"Religion": "Buddhist",
			"Sex":      "Male",
			"Oversight": []string{
				"Water",
			},
		},
		"Ratnasambhava": {
			"Religion": "Buddhist",
			"Sex":      "Male",
			"Oversight": []string{
				"Earth",
				"Spring",
			},
		},
		"Amitabha": {
			"Religion": "Buddhist",
			"Sex":      "Male",
			"Oversight": []string{
				"Life",
			},
		},
		"Amoghasiddhi": {
			"Religion": "Buddhist",
			"Sex":      "Male",
			"Oversight": []string{
				"Air",
			},
		},
		"Bhaisajyaguru": {
			"Religion": "Buddhist",
			"Sex":      "Male",
			"Oversight": []string{
				"Medicine",
			},
		},
		"Nageshvara Raja": {
			"Religion": "Buddhist",
			"Sex":      "Male",
			"Oversight": []string{
				"The Nagas",
			},
		},
		"Tara": {
			"Religion": "Buddhist",
			"Sex":      "Female",
			"Oversight": []string{
				"Protection From Fear",
			},
		},
		"Vajrayogini": {
			"Religion": "Buddhist",
			"Sex":      "Female",
			"Oversight": []string{
				"Passion",
			},
		},
		"Nairatmya": {
			"Religion": "Buddhist",
			"Sex":      "Female",
			"Oversight": []string{
				"None",
			},
		},
		"Kurukulla": {
			"Religion": "Buddhist",
			"Sex":      "Female",
			"Oversight": []string{
				"Magnetization",
			},
		},
	}

	for k, v := range buddhistDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
