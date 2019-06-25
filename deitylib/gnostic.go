package deitylib

import (
	"github.com/google/uuid"
)

func createGnosticDeities() {
	gnosticDeities := map[string]map[string]interface{}{
		"the Monad": {
			"Religion": "Gnostic",
			"Sex":      "Male",
			"Oversight": []string{
				"Everything",
			},
		},
	}

	for k, v := range gnosticDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
