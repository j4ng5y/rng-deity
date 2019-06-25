package deitylib

import (
	"github.com/google/uuid"
)

func createVodouDeities() {
	vodouDeities := map[string]map[string]interface{}{
		"Bondye": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Everything",
			},
		},
		"Adjassou-Linguetor": {
			"Religion": "Vodou",
			"Sex":      "Female",
			"Oversight": []string{
				"Spring Water",
			},
		},
		"Adjinakou": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Elephants",
			},
		},
		"Adya Houn'to": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Drums",
			},
		},
		"Agassou": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Dohomean Traditions",
			},
		},
		"Agwe": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Fish",
				"Aquatic Plants",
			},
		},
		"Aido Quedo": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Fertility",
				"Snakes",
			},
		},
		"Ayida-Weddo": {
			"Religion": "Vodou",
			"Sex":      "Female",
			"Oversight": []string{
				"Rainbows",
				"Snakes",
			},
		},
		"Ayizan": {
			"Religion": "Vodou",
			"Sex":      "Female",
			"Oversight": []string{
				"the Marketplace",
			},
		},
		"Azaka Medeh": {
			"Religion": "Vodou",
			"Sex":      "Female",
			"Oversight": []string{
				"the Harvest",
			},
		},
		"Azaka-Tonnerre": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Thunder",
				"Agriculture",
				"Farmers",
			},
		},
		"Becalou": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Evil",
			},
		},
		"Badessy": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"the Sky",
			},
		},
		"Baron La Croix": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"the Dead",
				"Sexuality",
			},
		},
		"Baron Samedi": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"the Dead",
			},
		},
		"Boli Shah": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Family",
			},
		},
		"Bossou Ashadeh": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Dahomey",
			},
		},
		"Boum'ba Maza": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Family",
			},
		},
		"Bugid Y Aiba": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"War",
			},
		},
		"Captian Debas": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Family",
			},
		},
		"Clermeil": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Flowing Waters",
			},
		},
		"Conga": {
			"Religion": "Vodou",
			"Sex":      "Female",
			"Oversight": []string{
				"Vodou",
			},
		},
		"Congo": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Vodou",
			},
		},
		"Dambalia": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Creation",
			},
		},
		"Dan Petro": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Farmers",
			},
		},
		"Dan Wedo": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"the King of France",
			},
		},
		"Diable Tonnere": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Thunder",
			},
		},
		"Diejuste": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Vodou",
			},
		},
		"Dinclinsin": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Vodou",
			},
		},
		"Eleggua": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Tricks",
			},
		},
		"Erzulie Dantor": {
			"Religion": "Vodou",
			"Sex":      "Female",
			"Oversight": []string{
				"Wealth",
				"Vengeance",
				"Protection",
			},
		},
		"Erzulie Freda Dahomey": {
			"Religion": "Vodou",
			"Sex":      "Female",
			"Oversight": []string{
				"Beauty",
				"Dancing",
				"Flowers",
				"Jewels",
				"Love",
				"Luxury",
			},
		},
		"Gran Ibo": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Wisdom",
				"Patience",
			},
		},
		"Gran Maitre": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Creation",
			},
		},
		"Gran Bois": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Creation",
			},
		},
		"Kalfu": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"the Night",
				"the Moon",
			},
		},
		"Lemba": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Rock",
				"Cannabelism",
			},
		},
		"L'inglesou": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Rock",
				"Ravines",
			},
		},
		"Loco": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Trees",
				"Plants",
				"Healers",
			},
		},
		"Mademoiselle Charlotte": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Caucasian Women",
			},
		},
		"Mait' Carrefour": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Magicians",
				"Crossroads",
			},
		},
		"Maitresse Delai": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Tambourine Players",
			},
		},
		"Maitresse Hounon'gon": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Fire",
			},
		},
		"Maman Brigitte": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Death",
			},
		},
		"Marassa": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Twins",
			},
		},
		"Marassa Jumeaux": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Twins",
			},
		},
		"Marinette": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Vodou",
			},
		},
		"Mambo": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Storms",
			},
		},
		"Mounanchou": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Vodou",
			},
		},
		"Nago Shango": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Vodou",
			},
		},
		"Obatala": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Creation",
			},
		},
		"Ogoun": {
			"Religion": "Vodou",
			"Sex":      "Female",
			"Oversight": []string{
				"Fire",
				"Iron",
				"Politics",
				"Thunder",
				"War",
			},
		},
		"Oshun": {
			"Religion": "Vodou",
			"Sex":      "Female",
			"Oversight": []string{
				"Love",
			},
		},
		"Oya": {
			"Religion": "Vodou",
			"Sex":      "Female",
			"Oversight": []string{
				"Warriors",
			},
		},
		"Papa Legba": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Ghosts",
				"The Underworld",
			},
		},
		"Pie": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Floods",
				"Soldiers",
			},
		},
		"Simbi": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Water Snakes",
			},
		},
		"Sobo": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Thunder",
			},
		},
		"Sousson-Pannan": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Evil",
				"Plagues",
				"Disease",
			},
		},
		"Ti Jean Quinto": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Police",
			},
		},
		"Ti Malice": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Trickery",
			},
		},
		"Ti-Jean Petro": {
			"Religion": "Vodou",
			"Sex":      "Male",
			"Oversight": []string{
				"Snakes",
			},
		},
		"Yemalla": {
			"Religion": "Vodou",
			"Sex":      "Female",
			"Oversight": []string{
				"Creation",
				"Mothers",
			},
		},
	}

	for k, v := range vodouDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
