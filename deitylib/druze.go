package deitylib

import (
	"github.com/google/uuid"
)

func createDruzeDeities() {
	druzeDeities := map[string]map[string]interface{}{
		"Allah": {
			"Religion": "Druze",
			"Sex":      "Male",
			"Oversight": []string{
				"Everything",
			},
		},
	}

	for k, v := range druzeDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
