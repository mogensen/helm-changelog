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

author: { "email": "%aE", "date": "%ad" }
commiter: { "email": "%cE", "date": "%cd" }
`

// I removed the Name field from the GitPerson as it is not currently used and I bumped in a quote in a name that caused yaml panic: https://github.com/argoproj/argo-helm/commit/0956363ebb0d1449e86be457e2fa96fb77ddf6d4.patch
// If needed this field in the future we can do some complicated quote escaping: https://gist.github.com/textarcana/1306223?permalink_comment_id=3915918#gistcomment-3915918
type GitPerson struct {
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
