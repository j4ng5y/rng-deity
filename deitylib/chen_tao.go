package deitylib

import (
	"github.com/google/uuid"
)

func createChenTaoDeities() {
	chenTaoDeities := map[string]map[string]interface{}{
		"Yuanshi Tianzun": {
			"Religion": "Chen Tao",
			"Sex":      "Male",
			"Oversight": []string{
				"Origin",
			},
		},
		"Lingbao Tianzun": {
			"Religion": "Chen Tao",
			"Sex":      "Male",
			"Oversight": []string{
				"Good",
				"Evil",
			},
		},
		"Daode Tianzun": {
			"Religion": "Chen Tao",
			"Sex":      "Male",
			"Oversight": []string{
				"Man",
			},
		},
	}

	for k, v := range chenTaoDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
