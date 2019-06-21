package deitylib

import (
	"github.com/google/uuid"
)

func createFamilyOfGodDeities() {
	familyOfGodDeities := map[string]map[string]interface{}{
		"God": {
			"Religion": "Family of God",
			"Sex":      "Male",
			"Oversight": []string{
				"Everything",
			},
		},
	}

	for k, v := range familyOfGodDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
