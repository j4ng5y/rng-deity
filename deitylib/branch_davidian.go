package deitylib

import (
	"github.com/google/uuid"
)

func createBranchDividianDeities() {
	branchDividianDeities := map[string]map[string]interface{}{
		"God": {
			"Religion": "Branch Dividian",
			"Sex":      "Male",
			"Oversight": []string{
				"Everything",
			},
		},
	}

	for k, v := range branchDividianDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
