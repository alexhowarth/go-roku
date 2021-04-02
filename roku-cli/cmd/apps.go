package cmd

import (
	"log"

	"github.com/alexhowarth/go-roku"
	"github.com/spf13/cobra"
)

var all, active bool

var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "Get apps installed on device",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := roku.NewClient(url)
		if err != nil {
			log.Fatalf("Unable to create client: %v", err)
		}
		if active {
			active, err := client.ActiveApp()
			if err != nil {
				log.Fatalf("Unable to get Active App: %v", err)
			}
			printJSON(active)
		} else {
			apps, err := client.Apps()
			if err != nil {
				log.Fatalf("Unable to get Apps: %v", err)
			}
			printJSON(apps)
		}
	},
}

func init() {
	appsCmd.Flags().BoolVar(&all, "all", false, "list all apps")
	appsCmd.Flags().BoolVar(&active, "active", false, "list active app")
	appsCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(appsCmd)
}
