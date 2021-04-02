package cmd

import (
	"log"

	"github.com/alexhowarth/go-roku"
	"github.com/spf13/cobra"
)

var appid, contentid, mediatype string

var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Launch app on device",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := roku.NewClient(url)
		if err != nil {
			log.Fatalf("Unable to create client: %v", err)
		}
		err = client.Launch(appid, contentid, mediatype)
		if err != nil {
			log.Fatalf("Unable to launch %#v", err)
		}
	},
}

func init() {
	launchCmd.Flags().StringVar(&appid, "appid", "", "app to launch")
	launchCmd.Flags().StringVar(&contentid, "contentid", "", "content to launch")
	launchCmd.Flags().StringVar(&mediatype, "mediatype", "", "media type launch")
	launchCmd.MarkFlagRequired("appid")
	rootCmd.AddCommand(launchCmd)
}
