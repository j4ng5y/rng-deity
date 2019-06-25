package deitylib

import (
	"fmt"
	"math/rand"
	"time"
)

// Deities is a simple structure that houses an array of Diety structures under Entries
type Deities struct {
	Entries []Diety
}

// Diety is a struct that describes a diety
type Diety struct {
	ID        string
	Name      string
	Religion  string
	Sex       string
	Oversight []string
}

// DS is a pointer to Deities to be used elsewhere in the app
var DS = &Deities{}

// NewDiety returns a pointer to the updated instance of the Dieties structure
//
// Variables:
//     name (string):        The name of the diety
//     religion (string):    The religion the diety belongs to
//     sex (string):         The sex of the diety
//     oversight ([]string): An array of the gods responsibilities
//
// Returns:
//     A pointer to the Dieties structure containing the new information
func (Ds *Deities) NewDiety(id string, name string, religion string, sex string, oversight []string) *Deities {
	D := Diety{
		ID:        id,
		Name:      name,
		Religion:  religion,
		Sex:       sex,
		Oversight: oversight,
	}

	Ds.Entries = append(Ds.Entries, D)

	return Ds
}

// GetRandomDeity will print a random deity
//
// Variables:
//     None
//
// Returns:
//     A Random Deity (string)
func (Ds *Deities) GetRandomDeity() {
	var title string

	rand.Seed(time.Now().UnixNano())
	deity := Ds.Entries[rand.Intn(len(Ds.Entries))]

	if deity.Religion == "Buddhist" {
		title = "Buddah"
	} else if deity.Religion == "Jedi" {
		title = "Influencer"
	} else {
		if deity.Sex == "Male" {
			title = "God"
		} else if deity.Sex == "Female" {
			title = "Goddess"
		} else {
			title = "Deity"
		}
	}

	fmt.Printf("%s, the %s %s of %s\n", deity.Name, deity.Religion, title, deity.Oversight[rand.Intn(len(deity.Oversight))])
}

// BuildDeityStruct will create and fill in the Deities structure
//
// Variables:
//     None
//
// Returns:
//     None
func BuildDeityStruct() {
	createAmishDeities()
	createAsatruDeities()
	createBuddhistDeities()
	createBahaiDeities()
	createBrahamKumariDeities()
	createBranchDividianDeities()
	createJediDeities()
	createCelticDeities()
	createChenTaoDeities()
	createChineseDeities()
	createChristianDeities()
	createChristadelphianDeities()
	createConcernedChristianDeities()
	createFamilyOfGodDeities()
	createDamiXuanjiaoDeities()
	createHinduDeities()
	createDruidDeities()
	createDruzeDeities()
	createEbioniteDeities()
	createEckankarDeities()
	createEgyptianDeities()
	createGnosticDeities()
	createGreekDeities()
	createHareKrishnaDeities()
	createHeathenismDeities()
	createVodouDeities()
	createIslamDeities()
}
