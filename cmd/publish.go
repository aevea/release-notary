package cmd

import (
	"errors"
	"fmt"
	"strings"

	history "github.com/outillage/git/pkg"
	"github.com/outillage/release-notary/internal/releaser"
	"github.com/outillage/release-notary/internal/slack"
	"github.com/outillage/release-notary/internal/text"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var simpleOutput bool

func init() {
	publishCmd.PersistentFlags().BoolVar(&simpleOutput, "simple", false, "use simplified output")
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

		dryRun := false
		if cmd.Flag("dry-run").Value.String() == "true" {
			dryRun = true
		}

		repo, err := history.OpenGit(".", debug)

		if err != nil {
			return err
		}

		commits, err := getCommits(repo)

		if err != nil {
			return err
		}

		var parsedCommits []text.Commit

		for _, commit := range commits {
			commitObject, err := repo.Commit(commit)

			if err != nil {
				return err
			}

			parsedCommit := text.ParseCommitMessage(commitObject.Message)

			parsedCommit.Hash = text.Hash(commitObject.Hash)

			parsedCommits = append(parsedCommits,
				parsedCommit,
			)
		}

		sections := text.SplitSections(parsedCommits)

		releaseNotes := text.ReleaseNotes{
			Simple: simpleOutput,
		}

		content := releaseNotes.Generate(sections, dryRun)

		if err != nil {
			return err
		}

		if dryRun {
			fmt.Println("my job here is done...")
			return nil
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
			service, err := releaser.CreateReleaser(options)
			if err != nil {
				return err
			}

			releaseService = service
		}

		if viper.IsSet("GITLAB_TOKEN") {
			if !viper.IsSet("CI_COMMIT_TAG") {
				fmt.Print("Release Notary is not running on a tag or CI_COMMIT_TAG is not set")
				return nil
			}

			options := releaser.Options{
				Token:     viper.GetString("GITLAB_TOKEN"),
				APIURL:    viper.GetString("CI_API_V4_URL"),
				TagName:   viper.GetString("CI_COMMIT_TAG"),
				ProjectID: viper.GetInt("CI_PROJECT_ID"),
				Provider:  "gitlab",
			}

			service, err := releaser.CreateReleaser(options)
			if err != nil {
				return err
			}

			releaseService = service
		}

		if releaseService == nil {
			return errors.New("Missing release service, please consult documentation on required env vars")
		}

		err = releaseService.Release(content)

		if err != nil {
			return err
		}

		if viper.IsSet("SLACK_WEBHOOK") {
			slack := &slack.Slack{
				WebHookURL: viper.GetString("SLACK_WEBHOOK"),
			}

			err = slack.Publish(sections)

			if err != nil {
				return err
			}
		}

		return nil
	},
}
