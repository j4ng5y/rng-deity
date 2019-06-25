package deitylib

import (
	"github.com/google/uuid"
)

func createEgyptianDeities() {
	egyptianDeities := map[string]map[string]interface{}{
		"Aker": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"the Earth",
				"the Horizon",
			},
		},
		"Amun": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"Creation",
				"the City of Thebes",
			},
		},
		"Anhur": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"War",
				"Hunting",
			},
		},
		"Aten": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"the Sun",
			},
		},
		"Atum": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"Creation",
				"the Sun",
			},
		},
		"Bennu": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"Creation",
				"the Sun",
			},
		},
		"Geb": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"the Earth",
			},
		},
		"Hapi": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"Floods",
			},
		},
		"Horus": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"the Sky",
				"the Sun",
				"Kings",
				"Protection",
				"Healing",
			},
		},
		"Khepri": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"the Sun",
				"Creation",
				"Mornings",
			},
		},
		"Khnum": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"Rams",
				"Life",
				"Floods",
				"the City of Elephantine",
			},
		},
		"Khonsu": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"the Moon",
			},
		},
		"Maahes": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"Lions",
			},
		},
		"Montu": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"War",
				"the Sun",
			},
		},
		"Nefertum": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"Lotus Blooms",
			},
		},
		"Nemty": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"Falcons",
				"Ferrymen",
			},
		},
		"Neper": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"Grain",
			},
		},
		"Osiris": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"Death",
				"Resurection",
				"the Underworld",
				"Vegetation Enlivenment",
				"Deceased Souls",
			},
		},
		"Ptah": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"Creation",
				"Craftsmen",
				"the City of Memphis",
			},
		},
		"Ra": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"the Sun",
				"Creation",
				"the Afterlife",
				"the City of Heliopolis",
			},
		},
		"Set": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"Violence",
				"Chaos",
				"Strength",
				"the Desert",
			},
		},
		"Shu": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"Wind",
				"Air",
			},
		},
		"Sobek": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"Crocodiles",
			},
		},
		"Sopdu": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"the Sky",
				"the East",
			},
		},
		"Thoth": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"the Moon",
				"Writing",
				"Scribes",
				"the City of Hermopolis",
			},
		},
		"Wadj-wer": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"the Mediterranean Sea",
				"Lakes",
			},
		},
		"Anumbis": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"Embalming",
				"the Protection of the Dead",
			},
		},
		"Heh": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"Infinity",
			},
		},
		"Kek": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"Darkness",
				"Chaos",
			},
		},
		"Nu": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"Disorder",
			},
		},
		"Tatenen": {
			"Religion": "Egyptian",
			"Sex":      "Male",
			"Oversight": []string{
				"the Earth",
			},
		},
		"Amunet": {
			"Religion": "Egyptian",
			"Sex":      "Female",
			"Oversight": []string{
				"Creation",
				"the City of Thebes",
			},
		},
		"Anuket": {
			"Religion": "Egyptian",
			"Sex":      "Female",
			"Oversight": []string{
				"the South",
			},
		},
		"Bastet": {
			"Religion": "Egyptian",
			"Sex":      "Female",
			"Oversight": []string{
				"the City of Bubastis",
				"the Protection from Evil",
			},
		},
		"Bat": {
			"Religion": "Egyptian",
			"Sex":      "Female",
			"Oversight": []string{
				"Cows",
			},
		},
		"Hathor": {
			"Religion": "Egyptian",
			"Sex":      "Female",
			"Oversight": []string{
				"the Sky",
				"the Sun",
				"Sexuality",
				"Motherhood",
			},
		},
		"Heqet": {
			"Religion": "Egyptian",
			"Sex":      "Female",
			"Oversight": []string{
				"Frogs",
				"Childbirth",
			},
		},
		"Hesat": {
			"Religion": "Egyptian",
			"Sex":      "Female",
			"Oversight": []string{
				"Cows",
			},
		},
		"Imentet": {
			"Religion": "Egyptian",
			"Sex":      "Female",
			"Oversight": []string{
				"the Afterlife",
			},
		},
		"Isis": {
			"Religion": "Egyptian",
			"Sex":      "Female",
			"Oversight": []string{
				"Funerals",
				"Motherhood",
				"Protection",
				"Magic",
			},
		},
		"Maat": {
			"Religion": "Egyptian",
			"Sex":      "Female",
			"Oversight": []string{
				"Truth",
				"Justice",
				"Order",
			},
		},
		"Menhit": {
			"Religion": "Egyptian",
			"Sex":      "Female",
			"Oversight": []string{
				"Lionesses",
			},
		},
		"Mut": {
			"Religion": "Egyptian",
			"Sex":      "Female",
			"Oversight": []string{
				"Consorts",
			},
		},
		"Neith": {
			"Religion": "Egyptian",
			"Sex":      "Female",
			"Oversight": []string{
				"Creation",
				"Hunting",
				"the City of Sais",
			},
		},
		"Nekhbet": {
			"Religion": "Egyptian",
			"Sex":      "Female",
			"Oversight": []string{
				"Vultures",
				"Upper Egypt",
			},
		},
		"Nephthys": {
			"Religion": "Egyptian",
			"Sex":      "Female",
			"Oversight": []string{
				"Consorts",
			},
		},
		"Nepit": {
			"Religion": "Egyptian",
			"Sex":      "Female",
			"Oversight": []string{
				"Grain",
			},
		},
		"Nut": {
			"Religion": "Egyptian",
			"Sex":      "Female",
			"Oversight": []string{
				"the Sky",
			},
		},
		"Pakhet": {
			"Religion": "Egyptian",
			"Sex":      "Female",
			"Oversight": []string{
				"Lionesses",
			},
		},
		"Renenutet": {
			"Religion": "Egyptian",
			"Sex":      "Female",
			"Oversight": []string{
				"Agriculture",
			},
		},
		"Satet": {
			"Religion": "Egyptian",
			"Sex":      "Female",
			"Oversight": []string{
				"the South",
			},
		},
		"Sekhmet": {
			"Religion": "Egyptian",
			"Sex":      "Female",
			"Oversight": []string{
				"Destruction",
				"Violence",
				"the Protection from Disease",
			},
		},
		"Tefnut": {
			"Religion": "Egyptian",
			"Sex":      "Female",
			"Oversight": []string{
				"Moisture",
			},
		},
		"Wadjet": {
			"Religion": "Egyptian",
			"Sex":      "Female",
			"Oversight": []string{
				"Cobras",
				"Lower Egypt",
			},
		},
		"Wosret": {
			"Religion": "Egyptian",
			"Sex":      "Female",
			"Oversight": []string{
				"the City of Thebes",
			},
		},
		"Anput": {
			"Religion": "Egyptian",
			"Sex":      "Female",
			"Oversight": []string{
				"Embalming",
				"the Protection of the Dead",
			},
		},
		"Kauket": {
			"Religion": "Egyptian",
			"Sex":      "Female",
			"Oversight": []string{
				"Darkness",
				"Chaos",
			},
		},
		"Re": {
			"Religion": "Egyptian",
			"Sex":      "Female",
			"Oversight": []string{
				"the Sun",
				"Creation",
				"the Afterlife",
				"the City of Heliopolis",
			},
		},
	}

	for k, v := range egyptianDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
