# helm-changelog

[![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2Fmogensen%2Fhelm-changelog%2Fbadge%3Fref%3Dmain&style=flat)](https://actions-badge.atrox.dev/mogensen/helm-changelog/goto?ref=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/mogensen/helm-changelog)](https://goreportcard.com/report/github.com/mogensen/helm-changelog)

Create changelogs for Helm Charts, based on git history

## Installation

All relevant commands are added to the `Makefile`

```bash
$ make help

Usage:
  make <target>

Targets:
  help                  Display this help
  test-unit             Run unit-tests
  test-integration      Run integration-tests
  build                 Build binary
  verify                tests and builds
  image                 build docker image
  clean                 clean up created files
  all                   Runs test, build and docker
  test-coverage         Generate test coverage report
  lint                  Generate static analysis report
  update-docs           Upgrade automatic documentations
```

### Run Standalone

First build the standalone binary:

```bash
$ make build
go build -o ./bin/helm-changelog .
```

This results in a binary in `./bin/helm-changelog`.

```bash
# See all cli options
$ ./bin/helm-changelog --help
Create changelogs for Helm Charts, based on git history

Usage:
  helm-changelog [flags]

Flags:
  -h, --help               help for helm-changelog
  -v, --verbosity string   Log level (debug, info, warn, error, fatal, panic (default "warning")

# Run helm-changelog creator with default params
$ ./bin/helm-changelog
```

### Run in Docker

The helm-changelog can alternatively be run in docker.

This is done in a multi-stage build, and does not require a working go development environment.

```bash
# Building the application and docker image
$ make image

# Run the resulting docker image
$ docker run -it --rm -v $(pwd):/data mogensen/helm-changelog:latest
```

The helm-changelog app is running as a non-root user.
This is a security best-practice. If user `1000` does not have write access to the chart dir, you may need to run the container as an other user.


```bash
# Run docker image as current user
$ docker run --user $UID it --rm -v $(pwd):/data mogensen/helm-changelog:latest
```

---

## Testing

### Running Unit Tests

Unit tests are implemented with Go's standard test framework.

All tests are located in their own test-packages, enforcing that the tests only test the 
public interface of the go packages.

```bash
# Run unit tests and code coverage, including test for race conditions
$ make test-coverage
```
