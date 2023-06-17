package cmd

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/mogensen/helm-changelog/pkg/git"
	"github.com/mogensen/helm-changelog/pkg/helm"
	"github.com/mogensen/helm-changelog/pkg/output"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var changelogFilename string
var releaseTemplatePath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "helm-changelog",
	Short: "Create changelogs for Helm Charts, based on git history",
	Run: func(cmd *cobra.Command, args []string) {

		log := logrus.StandardLogger()

		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		g := git.Git{Log: log}

		gitBaseDir, err := g.FindGitRepositoryRoot()
		if err != nil {
			log.Fatalf("Could not determine git root directory. helm-changelog depends largely on git history.")
		}

		fileList, err := helm.FindCharts(currentDir)
		if err != nil {
			log.Fatal(err)
		}

		for _, chartFileFullPath := range fileList {
			log.Infof("Handling: %s\n", chartFileFullPath)

			fullChartDir := filepath.Dir(chartFileFullPath)
			chartFile := strings.TrimPrefix(chartFileFullPath, gitBaseDir+"/")
			relativeChartFile := strings.TrimPrefix(chartFileFullPath, currentDir+"/")
			relativeChartDir := filepath.Dir(relativeChartFile)

			allCommits, err := g.GetAllCommits(fullChartDir)
			if err != nil {
				log.Fatal(err)
			}

			releases := helm.CreateHelmReleases(log, chartFile, relativeChartDir, g, allCommits)

			changeLogFilePath := filepath.Join(fullChartDir, changelogFilename)
			output.Markdown(log, changeLogFilePath, releaseTemplatePath, releases)
		}
	},
}

// Execute sets all flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	var v string

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if err := setUpLogs(os.Stdout, v); err != nil {
			return err
		}
		return nil
	}

	rootCmd.PersistentFlags().StringVarP(&changelogFilename, "filename", "f", "Changelog.md", "Filename for changelog")
	rootCmd.PersistentFlags().StringVarP(&v, "verbosity", "v", logrus.WarnLevel.String(), "Log level (debug, info, warn, error, fatal, panic)")
	rootCmd.PersistentFlags().StringVarP(&releaseTemplatePath, "release-template", "r", "", "Path to a Go template used for each release")

	cobra.CheckErr(rootCmd.Execute())
}

// setUpLogs set the log output ans the log level
func setUpLogs(out io.Writer, level string) error {
	logrus.SetOutput(out)
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	logrus.SetLevel(lvl)
	return nil
}
