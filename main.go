package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mogensen/helm-changelog/pkg/git"
	"github.com/mogensen/helm-changelog/pkg/helm"
	"github.com/mogensen/helm-changelog/pkg/output"
	log "github.com/sirupsen/logrus"
)

func main() {

	currentDir, err := os.Getwd()
	CheckIfError(err)

	gitBaseDir, err := git.FindGitRepositoryRoot()
	if err != nil {
		log.Fatalf("Could not determine git root directory. helm-changelog depends largely on git history.")
	}

	fileList, err := helm.FindCharts(currentDir)
	CheckIfError(err)

	for _, chartFileFullPath := range fileList {
		log.Infof("Handling: %s\n", chartFileFullPath)

		fullChartDir := filepath.Dir(chartFileFullPath)
		chartFile := strings.TrimPrefix(chartFileFullPath, gitBaseDir+"/")
		relativeChartFile := strings.TrimPrefix(chartFileFullPath, currentDir+"/")
		relativeChartDir := filepath.Dir(relativeChartFile)

		allCommits, err := git.GetAllCommits(fullChartDir)
		CheckIfError(err)

		releases := createHelmReleases(chartFile, relativeChartDir, allCommits)

		changeLogFilePath := filepath.Join(fullChartDir, "Changelog.md")
		output.Markdown(changeLogFilePath, releases)

	}
}

func createHelmReleases(chartFile, chartDir string, commits []git.GitCommit) []*helm.Release {

	fullValuesFile := filepath.Join(filepath.Dir(chartFile), "values.yaml")
	relativeValuesFile := filepath.Join(chartDir, "values.yaml")

	res := []*helm.Release{}
	currentRelease := ""
	releaseCommits := []git.GitCommit{}

	log.Infof(" - Found commits for chart: %d\n", len(commits))

	for _, l := range commits {

		releaseCommits = append(releaseCommits, l)

		chartContent, err := git.GetFileContent(l.Commit, chartFile)
		if err != nil {
			log.Infof("Chart.yaml not found in: %s\n", l.Commit)

			continue
		}
		chart, err := helm.GetChart(strings.NewReader(chartContent))
		CheckIfError(err)

		if chart.Version != currentRelease {
			log.Infof(" - Found version: %s\n", chart.Version)

			r := &helm.Release{
				ReleaseDate: l.Author.Date,
				Chart:       chart,
				Commits:     releaseCommits,
			}
			res = append(res, r)
			currentRelease = chart.Version
			releaseCommits = []git.GitCommit{}
		}

	}

	// Check if we have any unreleased commits
	if len(releaseCommits) > 0 {
		chartContent, err := git.GetFileContent("HEAD", chartFile)
		if err == nil {
			chart, err := helm.GetChart(strings.NewReader(chartContent))
			CheckIfError(err)
			chart.Version = "Next Release"

			r := &helm.Release{
				ReleaseDate: nil,
				Chart:       chart,
				Commits:     releaseCommits,
			}
			res = append(res, r)
		}
	}

	// Diff values files across versions
	for v, release := range res {
		diff := ""
		if v > 0 {
			lastRelease := res[v-1]
			diff, _ = git.GetDiffBetweenCommits(lastRelease.Commits[len(lastRelease.Commits)-1].Commit, release.Commits[len(release.Commits)-1].Commit, relativeValuesFile)
		} else {
			diff, _ = git.GetFileContent(release.Commits[0].Commit, fullValuesFile)

		}
		release.ValueDiff = diff
	}

	return res
}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	log.Infof("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}
