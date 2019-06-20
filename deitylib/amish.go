package deitylib

import (
	"github.com/google/uuid"
)

func createAmishDeities() {
	amishDeities := map[string]map[string]interface{}{
		"God": {
			"Religion": "Amish Christianity",
			"Sex":      "Male",
			"Oversight": []string{
				"Everything",
			},
		},
	}

	for k, v := range amishDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
