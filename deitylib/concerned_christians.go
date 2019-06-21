package deitylib

import (
	"github.com/google/uuid"
)

func createConcernedChristianDeities() {
	concernedChristianDeities := map[string]map[string]interface{}{
		"God": {
			"Religion": "Concerned Christian",
			"Sex":      "Male",
			"Oversight": []string{
				"Everything",
			},
		},
	}

	for k, v := range concernedChristianDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
