package deitylib

import (
	"github.com/google/uuid"
)

func createHareKrishnaDeities() {
	harekrishnaDeities := map[string]map[string]interface{}{
		"Krishna": {
			"Religion": "Hare Krishna",
			"Sex":      "Male",
			"Oversight": []string{
				"Everything",
			},
		},
	}

	for k, v := range harekrishnaDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
