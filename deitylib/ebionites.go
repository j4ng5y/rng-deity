package deitylib

import (
	"github.com/google/uuid"
)

func createEbioniteDeities() {
	ebioniteDeities := map[string]map[string]interface{}{
		"Yahweh": {
			"Religion": "Ebionite",
			"Sex":      "Male",
			"Oversight": []string{
				"Everything",
			},
		},
	}

	for k, v := range ebioniteDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
