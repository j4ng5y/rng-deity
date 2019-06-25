package cmd

import (
	"log"
	"os"

	"github.com/j4ng5y/rng-deity/deitylib"
	"github.com/spf13/cobra"
)

var (
	rngDeity = &cobra.Command{
		Use:     "rng-deity",
		Short:   "A little app that outputs a random deity from a number of religions.",
		Long:    "",
		Version: "2018.06.25",

		Run: func(ccmd *cobra.Command, args []string) {
			deitylib.BuildDeityStruct()
			deitylib.DS.GetRandomDeity()
		},
	}
)

func init() {
	rngDeity.AddCommand(newReligion)
}

// Execute runs the application
func Execute() {
	err := rngDeity.Execute()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
