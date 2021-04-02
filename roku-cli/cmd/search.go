package cmd

import (
	"log"

	"github.com/alexhowarth/go-roku"
	"github.com/spf13/cobra"
)

var timeout int

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for devices on the network",
	Run: func(cmd *cobra.Command, args []string) {
		devices, err := roku.Search(timeout)
		if err != nil {
			log.Fatalf("Unable to find devices %#v", err)
		}
		printJSON(devices)
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		// disable inherited persistent required 'url' flag
		disableInheritedPersistentRequiredFlag(cmd, "url")
	},
}

func init() {
	searchCmd.Flags().IntVar(&timeout, "timeout", 2, "timeout in seconds")
	rootCmd.AddCommand(searchCmd)
}
