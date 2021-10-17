package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "release-notary",
	Short: "No root command available please check below",
	Long:  "",
}

// Verbose is used to allow verbose/debug output for any given command
var Verbose bool

// DryRun is used for showing the changelog without actually publishing it
var DryRun bool

// RepoDir is used for setting the repository directory
var RepoDir string

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&DryRun, "dry-run", "d", false, "dry run")
	rootCmd.PersistentFlags().StringVar(&RepoDir, "repo-dir", ".", "directory where repository exists")
}

// Execute just executes the rootCmd for Cobra
func Execute() error {
	return rootCmd.Execute()
}
