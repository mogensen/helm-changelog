package git

import (
	"bytes"
	"errors"
	"io"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.DebugLevel)
}

func FindGitRepositoryRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	log.Debugln(cmd)

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	log.Debugf("Result: %s", out)
	return strings.TrimSpace(string(out)), nil
}

func GetFileContent(hash, filePath string) (string, error) {
	cmd := exec.Command("git", "cat-file", "-p", hash+":"+filePath)
	log.Debugln(cmd)

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	log.Debugf("Result: %s", out)
	return string(out), nil
}

func GetAllCommits(chartPath string) ([]GitCommit, error) {
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
	log.Debugln(cmd)

	out, err := cmd.Output()
	if err != nil || len(out) == 0 {
		return []GitCommit{}, err
	}

	log.Debugf("Result: %s", out)

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
			log.Error(err)
			continue
		}

		log.Debugf("commit: %s %s", t.Commit, t.Subject)
		gitCommitList = append(gitCommitList, *t)
	}

	return gitCommitList, nil
}

func GetDiffBetweenCommits(start, end, diffPath string) (string, error) {
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
	log.Debugln(cmd)

	out, err := cmd.Output()
	if err != nil {
		return "err", err
	}

	log.Debugf("Result: %s", out)
	return string(out), nil
}
