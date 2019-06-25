package cmd

import (
	"fmt"
	"os"
	"log"
	"strings"
	"text/template"
	"github.com/spf13/cobra"
)

type newR struct {
	Name string
	LowerName string
}

var (
	religionName string

	newReligion = &cobra.Command{
	Use: "new-religion",
	Short: "Create a new religion file",
	Long: "",

	Run: newReligionCmd,
	}
)

func init() {
	newReligion.Flags().StringVarP(&religionName, "religion-name", "n", "", "The name of the religion to create")
}

func newReligionCmd(ccmd *cobra.Command, args []string) {
	RELIGION := newR{
		Name: religionName,
		LowerName: strings.ToLower(religionName),
	}
	tmpl, err := template.ParseFiles("./templates/religion_template.gotmpl")
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	f, err := os.Create(fmt.Sprintf("./deitylib/%s.go", RELIGION.LowerName))
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	defer f.Close()

	err = tmpl.Execute(f, RELIGION)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}