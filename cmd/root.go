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

// Verbose is used to allow verbose/debug output for any given command
var Verbose bool

// DryRun is used for showing the changelog without actually publishing it
var DryRun bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&DryRun, "dry-run", "d", false, "dry run")
}

// Execute just executes the rootCmd for Cobra
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
