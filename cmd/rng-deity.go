package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/j4ng5y/rng-deity/deitylib"
	"github.com/spf13/cobra"
)

var (
	version = "2018.06.24"

	rngDeity = &cobra.Command{
		Use:   "rng-deity",
		Short: "A little app that outputs a random deity from a number of religions.",
		Long:  "",

		Run: func(ccmd *cobra.Command, args []string) {
			deitylib.BuildDeityStruct()
			deitylib.DS.GetRandomDeity()
		},
	}

	verFlag = rngDeity.Flags().BoolP("version", "v", true, "Print the application version")
)

// Execute runs the application
func Execute() {
	if *verFlag {
		fmt.Printf("RNG-Deity: version %s", version)
		os.Exit(0)
	}
	err := rngDeity.Execute()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
