package deitylib

import (
	"github.com/google/uuid"
)

func createChineseDeities() {
	chineseDeities := map[string]map[string]interface{}{
		"Shang Tian": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Highest Heaven",
			},
		},
		"Huang Tian": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Yellow Heaven",
			},
		},
		"Hao Tian": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Vast Heaven",
			},
		},
		"Min Tian": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Compassionate Heaven",
			},
		},
		"Cang Tian": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Deep-Green Heaven",
			},
		},
		"Yudi": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Highest Heaven",
				"Purety",
				"Creation",
			},
		},
		"Doumu": {
			"Religion": "Chinese",
			"Sex":      "Female",
			"Oversight": []string{
				"The Great Chariot (The Big Dipper)",
			},
		},
		"Pangu": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"The Cosmos",
				"The Earth",
				"The Sky",
				"Yin",
				"Yang",
			},
		},
		"Xiwangmu": {
			"Religion": "Chinese",
			"Sex":      "Female",
			"Oversight": []string{
				"The Kunlun Mountain",
				"Shamanic Inspiration",
				"Death",
				"Immortality",
			},
		},
		"Yangwang": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"The Underworld",
			},
		},
		"Yinyanggong": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"The Union of Yin and Yang",
			},
		},
		"Huangdi": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"The Essence of Earth",
			},
		},
		"Cangdi": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"The Essence of Wood",
				"Fertility",
				"Spring",
			},
		},
		"Bixia": {
			"Religion": "Chinese",
			"Sex":      "Female",
			"Oversight": []string{
				"Fertility",
			},
		},
		"Heidi": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"The Essence of Water",
				"Winter",
			},
		},
		"Chidi": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"The Essence of Fire",
				"Summer",
			},
		},
		"Baidi": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"The Essence of Metal",
				"Autum",
			},
		},
		"Longshen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Dragons",
			},
		},
		"Baoshen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Hail",
			},
		},
		"Bala Chongshen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Insects",
			},
		},
		"Doushen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Smallpox",
			},
		},
		"Fengshen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"The Wind",
			},
		},
		"Haishen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"The Sea",
			},
		},
		"Heshen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Rivers",
			},
		},
		"Gushen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Vallies",
			},
		},
		"Huoshen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Fire",
			},
		},
		"Hushen": {
			"Religion": "Chinese",
			"Sex":      "Female",
			"Oversight": []string{
				"Lakes",
				"Foxes",
			},
		},
		"Jinshen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Gold",
			},
		},
		"Jingshen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Watersprings",
			},
		},
		"Leishen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Thunder",
			},
		},
		"Mushen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Woodlands",
			},
		},
		"Shanshen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Mountains",
			},
		},
		"Shuishen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Water",
			},
		},
		"Tudishen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"The Earth",
			},
		},
		"Wenshen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Plagues",
			},
		},
		"Xiangshuishen": {
			"Religion": "Chinese",
			"Sex":      "Female",
			"Oversight": []string{
				"The Xiang River",
			},
		},
		"Xueshen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Snow",
			},
		},
		"Yushen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Rain",
				"Jail",
				"Purgatory",
			},
		},
		"Xihe": {
			"Religion": "Chinese",
			"Sex":      "Female",
			"Oversight": []string{
				"The Sun",
			},
		},
		"Yueshen": {
			"Religion": "Chinese",
			"Sex":      "Female",
			"Oversight": []string{
				"The Moon",
			},
		},
		"Wendi": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Culture",
			},
		},
		"Wudi": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Military",
			},
		},
		"Baoshengdadi": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"the Protection of Life",
			},
		},
		"Canshen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Silkworms",
			},
		},
		"Caishen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Wealth",
			},
		},
		"Cangjie": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"the Chinese Characters",
			},
		},
		"Chenghuangshen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Moats",
				"Walls",
			},
		},
		"Cheshen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Vehicles",
			},
		},
		"Eriangshen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Engineering",
			},
		},
		"Guanyin": {
			"Religion": "Chinese",
			"Sex":      "Female",
			"Oversight": []string{
				"Compassion ",
			},
		},
		"Jigong": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Help",
			},
		},
		"Jiushen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Wine",
			},
		},
		"Luban": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Carpentry",
			},
		},
		"Lushen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Roads",
			},
		},
		"Ping'anshen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Peace",
			},
		},
		"Taoshen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Pottery",
			},
		},
		"Tuershen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Leveret",
			},
		},
		"Xishen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Joy",
			},
		},
		"Yaoshen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Medicine",
			},
		},
		"Zaoshen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"the Hearth",
			},
		},
		"Fuxing": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Happiness",
			},
		},
		"Luxing": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Success in Life",
				"Firmness",
			},
		},
		"Shouxing": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Longevity",
			},
		},
		"Huashen": {
			"Religion": "Chinese",
			"Sex":      "Female",
			"Oversight": []string{
				"Flowers",
			},
		},
		"Mashen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Horses",
			},
		},
		"Niushen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Cattle",
			},
		},
		"Langshen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Wolves",
			},
		},
		"Shushen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Trees",
			},
		},
		"Wugushen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"the Five Cereals",
			},
		},
		"Yuanshen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Monkeys",
			},
		},
		"Shimashen": {
			"Religion": "Chinese",
			"Sex":      "Male",
			"Oversight": []string{
				"Sesame",
			},
		},
	}

	for k, v := range chineseDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
