package deitylib

import (
	"github.com/google/uuid"
)

func createBahaiDeities() {
	bahaiDeities := map[string]map[string]interface{}{
		"God": {
			"Religion": "Baha'i",
			"Sex":      "Male",
			"Oversight": []string{
				"Everything",
			},
		},
	}

	for k, v := range bahaiDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
