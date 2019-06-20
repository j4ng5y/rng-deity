package cmd

import (
	"log"
	"os"

	"github.com/j4ng5y/random-deity/deitylib"
	"github.com/spf13/cobra"
)

var rngDeity = &cobra.Command{
	Use:   "rng-deity",
	Short: "",
	Long:  "",

	Run: func(ccmd *cobra.Command, args []string) {
		deitylib.BuildDeityStruct()
		deitylib.DS.GetRandomDeity()
	},
}

func Execute() {
	err := rngDeity.Execute()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
