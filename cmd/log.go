package cmd

import (
	history "github.com/aevea/git/v3"
	"github.com/aevea/quoad"
	"github.com/aevea/release-notary/internal/text"
	"github.com/apex/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(logCmd)
}

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Prints commits between two tags",
	Long:  "In default prints the commits between 2 tags. Can be overriden to specify exact commits.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if Verbose {
			log.SetLevel(log.DebugLevel)
		}

		repo, err := history.OpenGit(".")

		if err != nil {
			return err
		}

		currentCommit, err := repo.CurrentCommit()

		if err != nil {
			return err
		}

		lastTag, err := repo.PreviousTag(currentCommit.Hash)

		if err != nil {
			return err
		}

		commits, err := repo.CommitsBetween(currentCommit.Hash, lastTag.Hash)

		if err != nil {
			return err
		}

		var parsedCommits []quoad.Commit

		for _, commit := range commits {
			commitObject, err := repo.Commit(commit)

			if err != nil {
				return err
			}

			parsedCommits = append(parsedCommits,
				quoad.ParseCommitMessage(commitObject.Message),
			)
		}

		log.Info(text.BuildHistory(parsedCommits))

		return nil
	},
}
