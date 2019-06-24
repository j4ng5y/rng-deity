package deitylib

import (
	"github.com/google/uuid"
)

func createDudeistDeities() {
	dudeistDeities := map[string]map[string]interface{}{
		"Jeffrey Lebowski": {
			"Religion": "Dudeist",
			"Sex":      "Male",
			"Oversight": []string{
				"Dudes",
			},
		},
	}

	for k, v := range dudeistDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
