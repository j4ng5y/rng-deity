package deitylib

import (
	"github.com/google/uuid"
)

func createBrahamKumariDeities() {
	brahmaKumariDeities := map[string]map[string]interface{}{
		"The Supreme Soul": {
			"Religion": "Brahma Kumaris",
			"Sex":      "Male",
			"Oversight": []string{
				"Everything",
			},
		},
	}

	for k, v := range brahmaKumariDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
