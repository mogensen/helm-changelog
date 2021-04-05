package output

import (
	"github.com/mogensen/helm-changelog/pkg/git"
	"github.com/mogensen/helm-changelog/pkg/helm"
)

func reverseReleases(a []*helm.Release) []*helm.Release {
	reversed := []*helm.Release{}
	for i := range a {
		n := a[len(a)-1-i]
		reversed = append(reversed, n)
	}
	return reversed
}

func reverseCommits(a []git.GitCommit) []git.GitCommit {
	reversed := []git.GitCommit{}
	for i := range a {
		n := a[len(a)-1-i]
		reversed = append(reversed, n)
	}
	return reversed
}
