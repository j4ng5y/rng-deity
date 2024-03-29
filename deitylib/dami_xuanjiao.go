package deitylib

import (
	"github.com/google/uuid"
)

func createDamiXuanjiaoDeities() {
	damiXuanjiaoDeities := map[string]map[string]interface{}{
		"God": {
			"Religion": "Dami Xuanjiao",
			"Sex":      "Male",
			"Oversight": []string{
				"Everything",
			},
		},
	}

	for k, v := range damiXuanjiaoDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
