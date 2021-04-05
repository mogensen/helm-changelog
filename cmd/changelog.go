package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/mogensen/helm-changelog/pkg/git"
	"github.com/mogensen/helm-changelog/pkg/helm"
	"github.com/mogensen/helm-changelog/pkg/output"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// changelogCmd represents the simulate command
var changelogCmd = &cobra.Command{
	Use:   "create",
	Short: "Create changelogs for Helm Charts, based on git history.",
	Long:  `Create changelogs for Helm Charts, based on git history.`,
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

			changeLogFilePath := filepath.Join(fullChartDir, "Changelog.md")
			output.Markdown(changeLogFilePath, releases)
		}
	},
}
