package git

import (
	"bytes"
	"errors"
	"io"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Git struct {
	Log *logrus.Logger
}

func (g *Git) FindGitRepositoryRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	g.Log.Debugln(cmd)

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	g.Log.Debugf("Result: %s", out)
	return strings.TrimSpace(string(out)), nil
}

func (g *Git) GetFileContent(hash, filePath string) (string, error) {
	cmd := exec.Command("git", "cat-file", "-p", hash+":"+filePath)
	g.Log.Debugln(cmd)

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	g.Log.Debugf("Result: %s", out)
	return string(out), nil
}

func (g *Git) GetAllCommits(chartPath string) ([]GitCommit, error) {
	cmd := exec.Command(
		"git",
		"log",
		"--date=iso-strict",
		"--reverse",
		gitformat,
		"--",
		chartPath,
		":(exclude)"+chartPath+"/Changelog.md",
	)
	g.Log.Debugln(cmd)

	out, err := cmd.Output()
	if err != nil || len(out) == 0 {
		return []GitCommit{}, err
	}

	g.Log.Debugf("Result: %s", out)

	gitCommitList := []GitCommit{}
	dec := yaml.NewDecoder(bytes.NewReader(out))

	for {
		// create new GitCommit here
		t := new(GitCommit)
		// pass a reference to GitCommit reference
		err := dec.Decode(&t)
		if t == nil {
			continue
		}

		// break the loop in case of EOF
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			g.Log.Error(err)
			continue
		}

		g.Log.Debugf("commit: %s %s", t.Commit, t.Subject)
		gitCommitList = append(gitCommitList, *t)
	}

	return gitCommitList, nil
}

func (g *Git) GetDiffBetweenCommits(start, end, diffPath string) (string, error) {
	if start == end {
		return "", nil
	}
	cmd := exec.Command(
		"git",
		"--no-pager",
		"diff",
		start+"..."+end,
		"--",
		diffPath,
	)
	g.Log.Debugln(cmd)

	out, err := cmd.Output()
	if err != nil {
		return "err", err
	}

	g.Log.Debugf("Result: %s", out)
	return string(out), nil
}
