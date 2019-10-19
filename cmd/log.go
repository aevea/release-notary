package cmd

import (
	"log"

	history "github.com/commitsar-app/git/pkg"
	"github.com/commitsar-app/release-notary/internal/text"
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
		debug := false
		if cmd.Flag("verbose").Value.String() == "true" {
			debug = true
		}

		repo, err := history.OpenGit(".", debug)

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

		commits, err := repo.CommitsBetween(currentCommit.Hash, lastTag)

		if err != nil {
			return err
		}

		var parsedCommits []text.Commit

		for _, commit := range commits {
			commitObject, err := repo.Commit(commit)

			if err != nil {
				return err
			}

			parsedCommits = append(parsedCommits,
				text.ParseCommitMessage(commitObject.Message),
			)
		}

		log.Println(text.BuildHistory(parsedCommits))

		return nil
	},
}
