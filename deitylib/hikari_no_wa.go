package deitylib

import (
	"github.com/google/uuid"
)

func createHikariNoWaDeities() {
	hikariNoWaDeities := map[string]map[string]interface{}{
		"NAME": {
			"Religion": "Aleph",
			"Sex":      "SEX",
			"Oversight": []string{
				"THING",
			},
		},
	}

	for k, v := range hikariNoWaDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
