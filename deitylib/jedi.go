package deitylib

import (
	"github.com/google/uuid"
)

func createJediDeities() {
	jediDeities := map[string]map[string]interface{}{
		"The Force": {
			"Religion": "Jedi",
			"Sex":      "None",
			"Oversight": []string{
				"Everything",
			},
		},
	}

	for k, v := range jediDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
