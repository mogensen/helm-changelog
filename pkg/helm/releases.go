package helm

import (
	"path/filepath"
	"strings"

	"github.com/mogensen/helm-changelog/pkg/git"
	"github.com/sirupsen/logrus"
)

func CreateHelmReleases(log *logrus.Logger, chartFile, chartDir string, g git.Git, commits []git.GitCommit) []*Release {

	res := []*Release{}
	currentRelease := ""
	releaseCommits := []git.GitCommit{}

	log.Infof(" - Found commits for chart: %d\n", len(commits))

	for _, l := range commits {

		releaseCommits = append(releaseCommits, l)

		chartContent, err := g.GetFileContent(l.Commit, chartFile)
		if err != nil {
			log.Infof("Chart.yaml not found in: %s\n", l.Commit)
			continue
		}

		chart, err := GetChart(strings.NewReader(chartContent))
		if err != nil {
			log.Warnf("Ignoring Chart.yaml file that cannot be parsed: %s", err)
			continue
		}

		if chart.Version != currentRelease {
			log.Infof(" - Found version: %s\n", chart.Version)

			r := &Release{
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
		chartContent, err := g.GetFileContent("HEAD", chartFile)
		if err == nil {
			chart, err := GetChart(strings.NewReader(chartContent))
			if err != nil {
				log.Warnf("Ignoring Chart.yaml file that cannot be parsed: %s", err)
			} else {
				chart.Version = "Next Release"
				res = append(res, &Release{
					Chart:   chart,
					Commits: releaseCommits,
				})
			}
		}
	}

	// Diff values files across versions
	createValueDiffs(res, g, chartFile, chartDir)

	return res
}

func createValueDiffs(res []*Release, g git.Git, chartFile, chartDir string) {

	fullValuesFile := filepath.Join(filepath.Dir(chartFile), "values.yaml")
	relativeValuesFile := filepath.Join(chartDir, "values.yaml")
	// Diff values files across versions
	for v, release := range res {
		diff := ""
		if v > 0 {
			lastRelease := res[v-1]
			diff, _ = g.GetDiffBetweenCommits(lastRelease.Commits[len(lastRelease.Commits)-1].Commit, release.Commits[len(release.Commits)-1].Commit, relativeValuesFile)
		} else {
			diff, _ = g.GetFileContent(release.Commits[0].Commit, fullValuesFile)
		}
		release.ValueDiff = diff
	}
}
