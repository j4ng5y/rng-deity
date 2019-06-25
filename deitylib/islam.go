package deitylib

import (
	"github.com/google/uuid"
)

func createIslamDeities() {
	islamDeities := map[string]map[string]interface{}{
		"Allah": {
			"Religion": "Islamic",
			"Sex":      "Male",
			"Oversight": []string{
				"Everything",
			},
		},
	}

	for k, v := range islamDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
