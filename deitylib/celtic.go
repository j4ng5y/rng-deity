package deitylib

import (
	"github.com/google/uuid"
)

func createCelticDeities() {
	celticDeities := map[string]map[string]interface{}{
		"Abandinus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Cambridgeshire",
			},
		},
		"Abellio": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Apple Trees",
			},
		},
		"Aereda": {
			"Religion": "Gallic Celt",
			"Sex":      "Male",
			"Oversight": []string{
				"Serpents",
			},
		},
		"Alisanos": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Mountains",
			},
		},
		"Alus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Agriculture",
			},
		},
		"Ambisagrus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Thunder",
				"Lightning",
				"The Sky",
				"Wind",
				"Rain",
				"Hail",
			},
		},
		"Andeis": {
			"Religion": "Gallic Celt",
			"Sex":      "Male",
			"Oversight": []string{
				"The Pyrenees",
			},
		},
		"Ankou": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Death",
			},
		},
		"Apollo Cunomaglus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Hunting",
			},
		},
		"Apollo Grannus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Healing",
			},
		},
		"Arixus": {
			"Religion": "Gallic Celt",
			"Sex":      "Male",
			"Oversight": []string{
				"The Pyrenees",
			},
		},
		"Arpeninus": {
			"Religion": "Gallic Celt",
			"Sex":      "Male",
			"Oversight": []string{
				"The Pyrenees",
			},
		},
		"Artahe": {
			"Religion": "Gallic Celt",
			"Sex":      "Male",
			"Oversight": []string{
				"Protection",
			},
		},
		"Atepomarus": {
			"Religion": "Gallic Celt",
			"Sex":      "Male",
			"Oversight": []string{
				"Horses",
			},
		},
		"Bedaius": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"The Sea",
			},
		},
		"Belatucadros": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"War",
			},
		},
		"Belenus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Healing",
			},
		},
		"Bergimus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Heights",
				"Mountains",
			},
		},
		"Borvo": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Minerals",
				"Healing",
			},
		},
		"Brasennus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Nothing",
			},
		},
		"Caletos": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Herds",
			},
		},
		"Caturix": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"The Helvetii people",
			},
		},
		"Cernunnos": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Animals",
				"The Underworld",
			},
		},
		"Cissonius": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Trade",
			},
		},
		"Mars Cnabetius": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"War",
			},
		},
		"Condatis": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Confluences of Rivers",
			},
		},
		"Cusianus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Jupiter",
			},
		},
		"Deus Latis": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Unknown",
			},
		},
		"Deus Decavavius": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Rivers",
			},
		},
		"Deus Orevalus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Unknown",
			},
		},
		"Dis Pater": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"The Underworld",
			},
		},
		"Divano": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"War",
			},
		},
		"Dorminus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Hot Springs",
			},
		},
		"Intarabus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"The Treveri people",
			},
		},
		"Erditse": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Unknown",
			},
		},
		"Esus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Vegetation",
			},
		},
		"Glanis": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Healing Springs",
			},
		},
		"Gobannus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Smithing",
			},
		},
		"Lalonus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Meadows",
			},
		},
		"Lhamnagalla Sqnnagalla": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Cambridgeshire",
			},
		},
		"Jupiter Felvennis": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"The Sky",
			},
		},
		"Leno": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"The city of Lerins",
			},
		},
		"Leucetios": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Thunder",
			},
		},
		"Maponos": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Youth",
			},
		},
		"Matunus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Bears",
			},
		},
		"Moccus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Protection for Boars and Pigs",
			},
		},
		"Moritasgus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Healing",
			},
		},
		"Mullo": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Assisting Apollo",
			},
		},
		"Nemausus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"The temple of Nimes",
			},
		},
		"Niskus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"The Sea",
			},
		},
		"Nodens": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Healing",
				"The Sea",
				"Hunting",
				"Dogs",
			},
		},
		"Ogmios": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Eloquence",
			},
		},
		"Paronnus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Brixia",
			},
		},
		"Rudiobus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"War",
			},
		},
		"Smertrios": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"War",
			},
		},
		"Sucellus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Agriculture",
				"Wine",
			},
		},
		"Taranis": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Thunder",
			},
		},
		"Toutatis": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Tribal Protection",
			},
		},
		"Tridamos": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Abundance",
				"Bovine Triplication",
			},
		},
		"Ucuetis": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Blacksmithing",
			},
		},
		"Vasio": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"The city of Vaison-la-Romaine",
			},
		},
		"Vellaunus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Death",
				"The Underworld",
			},
		},
		"Vernostonus": {
			"Religion": "Brythonic Celt",
			"Sex":      "Male",
			"Oversight": []string{
				"Funerals",
				"War",
			},
		},
		"Vindonnus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Belenus",
			},
		},
		"Vinotonus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Vines",
				"Wine",
			},
		},
		"Viridios": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Plants",
			},
		},
		"Virotutis": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Assisting Apollo",
			},
		},
		"Visucius": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Trade",
			},
		},
		"Vitucadrus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Energy",
			},
		},
		"Vosegus": {
			"Religion": "Celtic",
			"Sex":      "Male",
			"Oversight": []string{
				"Vosges Mountains",
			},
		},
		"Acionna": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Rivers",
			},
		},
		"Adsullata": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"The River Savubalabada",
			},
		},
		"Aericula": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Goddesses",
			},
		},
		"Aeron": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"War",
				"Slaughter",
				"Rivers",
			},
		},
		"Alantedoba": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Agriculture",
			},
		},
		"Alaterviae": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Mothers",
			},
		},
		"Ammaca": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Grandmothers",
			},
		},
		"Ancamna": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"The Moselle River",
			},
		},
		"Ancasta": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"The Itchen River",
			},
		},
		"Andarta": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"War",
			},
		},
		"Andraste": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Victory",
			},
		},
		"Anesiaminehae": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Mothers",
			},
		},
		"Annea Clivana": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Spirit Protection",
			},
		},
		"Arduinna": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"The Ardennes Forest",
			},
		},
		"Arnemetia": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Water",
			},
		},
		"Artio": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Bears",
			},
		},
		"Axona": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"The River Aisne",
			},
		},
		"Baeserta": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"The Pyrenees",
			},
		},
		"Belisama": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Rivers",
				"Lakes",
				"Fire",
				"Crafts",
				"Light",
			},
		},
		"Bergusia": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Prosperity",
			},
		},
		"Bormana": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Minerals",
				"Spring Water",
			},
		},
		"Bricta": {
			"Religion": "Gallic Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Water",
			},
		},
		"Cailleach": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Creation",
				"Destruction",
			},
		},
		"Caimineae": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Mothers",
			},
		},
		"Cantrusteihiae": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Mothers",
			},
		},
		"Carlin": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Creation",
				"Destruction",
			},
		},
		"Carpundia": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Rivers",
			},
		},
		"Cathubodua": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"War",
			},
		},
		"Caticatona": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Water",
			},
		},
		"Clota": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"The River Clyde",
			},
		},
		"Clutoida": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Rivers",
			},
		},
		"Coinchend": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"War",
			},
		},
		"Coventina": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Wells",
				"Springs",
			},
		},
		"Domara": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Fertility",
			},
		},
		"Damona": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Healing",
				"Fertility",
				"Incubation",
			},
		},
		"Dea Latis": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Bogs",
				"Pools",
			},
		},
		"Dea Matrona": {
			"Religion": "Gallic Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"The River Marne",
			},
		},
		"Dea Mediotautehae": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Mothers",
			},
		},
		"Dea Meduna": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Healing",
			},
		},
		"Dea Sequana": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"The River Seine",
			},
		},
		"Deae Vediantiae": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Mothers",
			},
		},
		"Dervonnae": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Mothers",
			},
		},
		"Dibona": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Fountains",
			},
		},
		"Divona": {
			"Religion": "Gallic Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"The Spring of Burdigala",
			},
		},
		"Dominae": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Mothers",
			},
		},
		"Epona": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Protection",
				"Horses",
			},
		},
		"Erecure": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"The Earth",
			},
		},
		"Feminae": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Mothers",
			},
		},
		"Gobroig": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Wells",
				"Springs",
			},
		},
		"Habetrot": {
			"Religion": "Brythonic Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Spinning",
				"Healing",
			},
		},
		"Henwen": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Sows",
			},
		},
		"Erecura": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Death",
				"Fertility",
			},
		},
		"Histria": {
			"Religion": "Gallic Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Land",
			},
		},
		"Icaunus": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Rivers",
			},
		},
		"Icovellauna": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Springs",
			},
		},
		"Imona": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Wells",
			},
		},
		"Inciona": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"The Treveran People",
			},
		},
		"Lerina": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"The City of Lerins",
			},
		},
		"Litavis": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"The Earth",
			},
		},
		"Maiabus": {
			"Religion": "Gallic Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Growth",
			},
		},
		"Matres Britannae": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Mothers",
			},
		},
		"Matres Eburnicae": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Mothers",
			},
		},
		"Matres Mogontiones": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Mothers",
			},
		},
		"Matres Nemetiales": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Mothers",
			},
		},
		"Matres Ollototae": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Mothers",
			},
		},
		"Matris Augustis": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Mothers",
			},
		},
		"Matronae Aufaniea": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Mothers",
			},
		},
		"Matronae Dervonnae": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Mothers",
			},
		},
		"Matronae Ollogabiae": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Mothers",
			},
		},
		"Matronae Senonae": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Mothers",
			},
		},
		"Matronae Seleviae": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Mothers",
			},
		},
		"Motronae Vediantiae": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Mothers",
			},
		},
		"Maximia": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Fountians",
			},
		},
		"Melusine": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Mermaids",
			},
		},
		"Nantosuelta": {
			"Religion": "Gallic Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Nature",
				"Growth",
				"The Earth",
				"Fire",
				"Fertility",
			},
		},
		"Natae": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Unknown",
			},
		},
		"Niskai": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Water Sprites",
			},
		},
		"Ricagambeda": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Furrows",
			},
		},
		"Ritona": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"The Treveri People",
			},
		},
		"Rocloisiabo": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Mothers",
			},
		},
		"Rosmerta": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Fertility",
				"Abundance",
			},
		},
		"Seixomniai Leuciticai": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Child Birth",
				"Women",
			},
		},
		"Senua": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Springs",
			},
		},
		"Sequana": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"The River Seine",
			},
		},
		"Sueta": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Hot Springs",
			},
		},
		"Suleviae": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Healing Waters",
			},
		},
		"Sulis": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Healing Waters",
			},
		},
		"Tamesis": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"The River Thames",
			},
		},
		"Veica Noriceia": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"War",
			},
		},
		"Verbeia": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"The River Wharfe",
			},
		},
		"Vesunna": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"The city of Vesona",
			},
		},
		"Vibes": {
			"Religion": "Celtic",
			"Sex":      "Female",
			"Oversight": []string{
				"Unknown",
			},
		},
		"Amaethon": {
			"Religion": "Welsh Celt",
			"Sex":      "Male",
			"Oversight": []string{
				"Agriculture",
			},
		},
		"Beli Mawr": {
			"Religion": "Welsh Celt",
			"Sex":      "Male",
			"Oversight": []string{
				"Ancestors",
			},
		},
		"Gofannon": {
			"Religion": "Welsh Celt",
			"Sex":      "Male",
			"Oversight": []string{
				"Smithing",
			},
		},
		"Agrona": {
			"Religion": "Welsh Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"The River Ayr",
				"War",
				"Slaughter",
			},
		},
		"Arianrhod": {
			"Religion": "Welsh Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"The Moon",
				"The Stars",
			},
		},
		"Blodeuwedd": {
			"Religion": "Welsh Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Spring",
			},
		},
		"Branwen": {
			"Religion": "Welsh Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Love",
				"Beauty",
			},
		},
		"Ceridwen": {
			"Religion": "Welsh Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Rebirth",
				"Transformation",
			},
		},
		"Creiddylad": {
			"Religion": "Welsh Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Flowers",
				"Love",
			},
		},
		"Creirwy": {
			"Religion": "Welsh Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Beauty",
			},
		},
		"Don": {
			"Religion": "Welsh Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Mothers",
			},
		},
		"Modron": {
			"Religion": "Welsh Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"The River Marne",
			},
		},
		"Rhiannon": {
			"Religion": "Welsh Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Horses",
				"The Moon",
				"The Otherworld",
			},
		},
		"Builg": {
			"Religion": "Gaelic Celt",
			"Sex":      "Male",
			"Oversight": []string{
				"Fir Bholg",
			},
		},
		"Luchtaine": {
			"Religion": "Gaelic Celt",
			"Sex":      "Male",
			"Oversight": []string{
				"Carpentry",
			},
		},
		"Nelt": {
			"Religion": "Gaelic Celt",
			"Sex":      "Male",
			"Oversight": []string{
				"War",
			},
		},
		"Airmed": {
			"Religion": "Gaelic Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Healing",
				"Herbalism",
			},
		},
		"Anu": {
			"Religion": "Gaelic Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"The Earth",
				"Fertility",
			},
		},
		"Bec": {
			"Religion": "Gaelic Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Magical Wells",
			},
		},
		"Boann": {
			"Religion": "Gaelic Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"The River Boyne",
			},
		},
		"Brigid": {
			"Religion": "Gaelic Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Poetry",
				"Healing",
				"Smithing",
				"Spring",
			},
		},
		"Morrigan": {
			"Religion": "Gaelic Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"War",
				"Fate",
				"Doom",
				"Death",
				"Victory",
			},
		},
		"Besencia": {
			"Religion": "Lustanian Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Protection of Community and Household",
			},
		},
		"Erbina": {
			"Religion": "Lustanian Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Wild Animals",
				"Hunting",
				"Domestic Security",
			},
		},
		"Laneana": {
			"Religion": "Lustanian Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Spring",
				"Floods",
			},
		},
		"Nabia": {
			"Religion": "Lustanian Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Versitility",
			},
		},
		"Reva": {
			"Religion": "Lustanian Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Flowing Water",
			},
		},
		"Sedatus": {
			"Religion": "Germanian Celt",
			"Sex":      "Male",
			"Oversight": []string{
				"Security",
			},
		},
		"Apadeva": {
			"Religion": "Germanian Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Water",
			},
		},
		"Cissonia": {
			"Religion": "Germanian Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Trade",
			},
		},
		"Matres Mopates": {
			"Religion": "Germanian Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Mothers",
			},
		},
		"Matres Treverae": {
			"Religion": "Germanian Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Mothers",
			},
		},
		"Nehelennia": {
			"Religion": "Germanian Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Sailors",
			},
		},
		"Boria": {
			"Religion": "Illyrian Celt",
			"Sex":      "Male",
			"Oversight": []string{
				"The North Wind",
			},
		},
		"Eia": {
			"Religion": "Illyrian Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Healing",
			},
		},
		"Matres Nutrices": {
			"Religion": "Illyrian Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Mothers",
			},
		},
		"Trita": {
			"Religion": "Illyrian Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Health",
			},
		},
		"Tava": {
			"Religion": "Pictish Celt",
			"Sex":      "Female",
			"Oversight": []string{
				"Rivers",
			},
		},
	}

	for k, v := range celticDeities {
		id := uuid.New().String()
		DS = DS.NewDiety(id, k, v["Religion"].(string), v["Sex"].(string), v["Oversight"].([]string))
	}
}
