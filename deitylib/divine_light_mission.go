package deitylib

import (
	"github.com/google/uuid"
)

func createDLMDeities() {
	dlmDeities := map[string]map[string]interface{}{
		"Vishnu": {
			"Religion": "Hindu",
			"Sex":      "Male",
			"Oversight": []string{
				"Protection",
				"Preservation of Good",
				"Karma Restoration",
				"Moksha",
			},
		},
		"Shiva": {
			"Religion": "Hindu",
			"Sex":      "Male",
			"Oversight": []string{
				"Divine Energy",
				"Meditation",
				"Yoga",
				"the Arts",
				"Time",
				"Destruction",
				"Dance",
				"the Destruction of Evil",
				"the Devas",
			},
		},
		"Brahma": {
			"Religion": "Hindu",
			"Sex":      "Male",
			"Oversight": []string{
				"Creation",
				"Wisdom",
				"Moksha",
			},
		},
		"Ganesha": {
			"Religion": "Hindu",
			"Sex":      "Male",
			"Oversight": []string{
				"New Beginnings",
				"Success",
				"Wisdom",
				"Obsticle Removal",
			},
		},
		"Kartikeya": {
			"Religion": "Hindu",
			"Sex":      "Male",
			"Oversight": []string{
				"War",
				"Victory",
			},
		},
		"Parvati": {
			"Religion": "Hindu",
			"Sex":      "Female",
			"Oversight": []string{
				"Creative Power",
				"Divine Energy",
			},
		},
		"Lakshmi": {
			"Religion": "Hindu",
			"Sex":      "Female",
			"Oversight": []string{
				"Fortune",
				"Wealth",
				"Prosperity",
			},
		},
		"Saraswati": {
			"Religion": "Hindu",
			"Sex":      "Female",
			"Oversight": []string{
				"Knowledge",
				"Music",
				"Art",
				"Speech",
				"Wisdom",
				"Learning",
			},
		},
		"Durga": {
			"Religion": "Hindu",
			"Sex":      "Female",
			"Oversight": []string{
				"War",
			},
		},
		"Kali": {
			"Religion": "Hindu",
			"Sex":      "Female",
			"Oversight": []string{
				"Time",
				"Creation",
				"Destruction",
				"Violence",
				"Power",
			},
		},
		"Mariamman": {
			"Religion": "Hindu",
			"Sex":      "Female",
			"Oversight": []string{
				"Rain",
			},
		},
	}

	for k, v := range dlmDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
