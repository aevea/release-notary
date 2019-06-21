package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "release-notary",
	Short: "Commits up to last tag",
	Long:  "",
	Run: (func(cmd *cobra.Command, args []string) {
		log.Println("hi")
	}),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
