package cmd

import (
	"log"

	"github.com/fallion/release-notary/internal/history"
	"github.com/fallion/release-notary/internal/text"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
)

func init() {
	rootCmd.AddCommand(logCmd)
}

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Prints commits between two tags",
	Long:  "In default prints the commits between 2 tags. Can be overriden to specify exact commits.",
	Run: func(cmd *cobra.Command, args []string) {
		debug := false
		if cmd.Flag("verbose").Value.String() == "true" {
			debug = true
		}

		repo, _ := git.PlainOpen(".")

		currentCommit := history.CurrentCommit(repo, debug)

		lastTag, _ := history.PreviousTag(repo, currentCommit, debug)

		commits, _ := history.CommitsBetween(repo, currentCommit, lastTag)

		var commitMessages []string

		for i := 0; i < len(commits); i++ {
			commitMessages = append(commitMessages,
				text.TrimMessage(history.CommitMessage(repo, commits[i])),
			)
		}

		log.Println(text.BuildHistory(commitMessages))
	},
}
