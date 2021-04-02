package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var url string

var rootCmd = &cobra.Command{
	Use:   "roku-cli",
	Short: "Command line tool for Roku External Control Protocol",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&url, "url", "", "url of the device (location returned by search)")
	rootCmd.MarkPersistentFlagRequired("url")
}

func printJSON(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(string(b))
}

// https://github.com/spf13/cobra/issues/1330
func disableInheritedPersistentRequiredFlag(cmd *cobra.Command, name string) {
	cmd.InheritedFlags().SetAnnotation(name, cobra.BashCompOneRequiredFlag, []string{"false"})
}
