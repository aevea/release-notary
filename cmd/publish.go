package cmd

import (
	"strings"

	"github.com/commitsar-app/commitsar/pkg/history"
	"github.com/commitsar-app/release-notary/internal/releaser"
	"github.com/commitsar-app/release-notary/internal/text"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(publishCmd)
}

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publishes release notes",
	Long:  "",
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

		viper.AutomaticEnv()

		var releaseService *releaser.Releaser

		if viper.IsSet("GITHUB_TOKEN") {
			split := strings.Split(viper.GetString("GITHUB_REPOSITORY"), "/")

			options := releaser.Options{
				Token:    viper.GetString("GITHUB_TOKEN"),
				Owner:    split[0],
				Repo:     split[1],
				Provider: "github",
			}
			releaseService = releaser.CreateReleaser(options)
		}

		sections := text.SplitSections(parsedCommits)

		releaseNotes := text.ReleaseNotes(sections)

		err = releaseService.Release(releaseNotes)

		if err != nil {
			return err
		}

		return nil
	},
}
