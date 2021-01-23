package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	history "github.com/aevea/git/v2"
	"github.com/aevea/integrations"
	"github.com/aevea/quoad"
	"github.com/aevea/release-notary/internal/releaser"
	"github.com/aevea/release-notary/internal/slack"
	"github.com/aevea/release-notary/internal/text"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var complexOutput bool

func init() {
	publishCmd.PersistentFlags().BoolVar(&complexOutput, "complex", false, "use complex output")
	rootCmd.AddCommand(publishCmd)
}

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publishes release notes",
	Long:  "",
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

		commits, err := getCommits(repo)

		if err != nil {
			return err
		}

		var parsedCommits []quoad.Commit

		for _, commit := range commits {
			commitObject, err := repo.Commit(commit)

			if err != nil {
				return err
			}

			parsedCommit := quoad.ParseCommitMessage(commitObject.Message)

			parsedCommit.Hash = quoad.Hash(commitObject.Hash)

			parsedCommits = append(parsedCommits,
				parsedCommit,
			)
		}

		sections := text.SplitSections(parsedCommits)

		releaseNotes := text.ReleaseNotes{
			Complex: complexOutput,
		}

		content := releaseNotes.Generate(sections, DryRun)

		if err != nil {
			return err
		}

		if DryRun {
			fmt.Println("my job here is done...")
			return nil
		}

		viper.AutomaticEnv()

		var releaseService *releaser.Releaser

		if viper.IsSet("GITHUB_TOKEN") {
			split := strings.Split(viper.GetString("GITHUB_REPOSITORY"), "/")
			ref := viper.GetString("GITHUB_REF")

			options := releaser.Options{
				Token:    viper.GetString("GITHUB_TOKEN"),
				Owner:    split[0],
				Repo:     split[1],
				TagName:  integrations.GetCurrentRef(),
				Provider: "github",
			}

			const tagRef = "/refs/tags/"

			if strings.HasPrefix(ref, tagRef) {
				options.TagName = strings.TrimPrefix(ref, tagRef)
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
				TagName:   integrations.GetCurrentRef(),
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

			err = slack.Publish(sections, integrations.GetGitRemote())

			if err != nil {
				return err
			}
		}

		return nil
	},
}
