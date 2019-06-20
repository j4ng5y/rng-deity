package deitylib

import (
	"github.com/google/uuid"
)

func createRELIGIONDeities() {
	celticDeities := map[string]map[string]interface{}{
		"NAME": {
			"Religion": "RELIGION",
			"Sex":      "SEX",
			"Oversight": []string{
				"THING",
			},
		},
	}

	for k, v := range celticDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
