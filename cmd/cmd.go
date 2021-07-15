package cmd

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var (
	montlyPlanning int

	rootCmd = &cobra.Command{
		Use:   "gofin",
		Short: "A monthly financial planning with suggested daily allowance",
		Long: `Gofin must receive a monthly spending cap as a down payment. 
The app will suggest the daily average so that this plan is not exceeded.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Prints GoFin version",
		Run: func(cmd *cobra.Command, args []string) {
			if info, ok := debug.ReadBuildInfo(); ok {
				sum := info.Main.Sum
				if sum == "" {
					sum = "none"
				}
				fmt.Printf("https://%s %s @ %s\n", info.Main.Path, info.Main.Version, sum)
			} else {
				fmt.Println("unknown")
			}
		},
	}

	rootCmd.Flags().IntVar(&montlyPlanning, "mp", 0, "Montly financial Planning")

	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
