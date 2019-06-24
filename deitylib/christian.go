package deitylib

import (
	"github.com/google/uuid"
)

func createChristianDeities() {
	christianDeities := map[string]map[string]interface{}{
		"God": {
			"Religion": "Christian",
			"Sex":      "Male",
			"Oversight": []string{
				"Everything",
			},
		},
		"Jesus Christ": {
			"Religion": "Christian",
			"Sex":      "Male",
			"Oversight": []string{
				"Everything",
			},
		},
	}

	for k, v := range christianDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
