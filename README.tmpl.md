# helm-changelog

[![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2Fmogensen%2Fhelm-changelog%2Fbadge%3Fref%3Dmain&style=flat)](https://actions-badge.atrox.dev/mogensen/helm-changelog/goto?ref=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/mogensen/helm-changelog)](https://goreportcard.com/report/github.com/mogensen/helm-changelog)

Create changelogs for Helm Charts, based on git history.

The application depends on the assumption that the helm chart is released on the first commit where the version number is bumped in the `Chart.yaml` file.

All sub-sequent commits are grouped as commits for the next release, until the version number is bumped again.

## Features

The changelog contains the following things:

- Commits are grouped by releases
- Each release displays 
  * Supported Helm version
  * Release date
  * App Version for the chart
  * Supported Kubernetes version
- Changes in default helm values

## Examples

This repository contains a set of example changelogs created for the github.com/elastic/helm-charts charts.

 - [apm-server](https://github.com/mogensen/helm-changelog/tree/main/examples/apm-server.md)
 - [elasticsearch](https://github.com/mogensen/helm-changelog/tree/main/examples/elasticsearch.md)
 - [filebeat](https://github.com/mogensen/helm-changelog/tree/main/examples/filebeat.md)
 - [kibana](https://github.com/mogensen/helm-changelog/tree/main/examples/kibana.md)
 - [logstash](https://github.com/mogensen/helm-changelog/tree/main/examples/logstash.md)
 - [metricbeat](https://github.com/mogensen/helm-changelog/tree/main/examples/metricbeat.md)

## Installation

All relevant commands are added to the `Makefile`

```bash
$ make help
{{ .Env.MAKE_HELPTEXT }}
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
{{ .Env.HELM_CHANGELOG_HELPTEXT }}

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
This is a security best-practice. If user `1000` does not have write access to the chart directory, you may need to run the container as an other user.


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
