package git

import "time"

// gitformat defines a yaml formatting string for git
var gitformat string = `--pretty=format:
---
commit: "%H"
parent: "%P"
refs: "%D"
subject: |-
  %s

author: { "name": "%aN", "email": "%aE", "date": "%ad" }
commiter: { "name": "%cN", "email": "%cE", "date": "%cd" }
`

type GitPerson struct {
	Name  string     `yaml:"name"`
	Email string     `yaml:"email"`
	Date  *time.Time `yaml:"date"`
}

type GitCommit struct {
	Commit   string    `yaml:"commit"`
	Parent   string    `yaml:"parent"`
	Refs     string    `yaml:"refs"`
	Subject  string    `yaml:"subject"`
	Author   GitPerson `yaml:"author"`
	Commiter GitPerson `yaml:"commiter"`
}
