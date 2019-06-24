package cmd

import (
	"log"
	"os"

	"github.com/j4ng5y/rng-deity/deitylib"
	"github.com/spf13/cobra"
)

var rngDeity = &cobra.Command{
	Use:   "rng-deity",
	Short: "A little app that outputs a random deity from a number of religions.",
	Long:  "",

	Run: func(ccmd *cobra.Command, args []string) {
		deitylib.BuildDeityStruct()
		deitylib.DS.GetRandomDeity()
	},
}

// Execute runs the application
func Execute() {
	err := rngDeity.Execute()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
