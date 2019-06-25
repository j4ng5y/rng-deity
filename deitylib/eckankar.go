package deitylib

import (
	"github.com/google/uuid"
)

func createEckankarDeities() {
	eckankarDeities := map[string]map[string]interface{}{
		"God": {
			"Religion": "Eckankar",
			"Sex":      "Male",
			"Oversight": []string{
				"Everything",
			},
		},
	}

	for k, v := range eckankarDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
