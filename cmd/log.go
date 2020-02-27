package cmd

import (
	"io/ioutil"
	"log"
	"os"

	history "github.com/outillage/git/v2"
	"github.com/outillage/quoad"
	"github.com/outillage/release-notary/internal/text"
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
		debugLogger := log.Logger{}
		debugLogger.SetPrefix("[DEBUG] ")
		debugLogger.SetOutput(os.Stdout)

		if !Verbose {
			debugLogger.SetOutput(ioutil.Discard)
			debugLogger.SetPrefix("")
		}

		repo, err := history.OpenGit(".", &debugLogger)

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

		log.Println(text.BuildHistory(parsedCommits))

		return nil
	},
}
