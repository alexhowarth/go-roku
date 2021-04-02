package cmd

import (
	"log"

	"github.com/alexhowarth/go-roku"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info for device",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := roku.NewClient(url)
		if err != nil {
			log.Fatalf("Unable to create client: %v", err)
		}
		info, err := client.DeviceInfo()
		if err != nil {
			log.Fatalf("Unable to get device info: %v", err)
		}
		printJSON(info)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
