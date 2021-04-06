#!/bin/bash

MAKE_HELPTEXT=$(make -s help | sed $'s,\x1b\\[[0-9;]*[a-zA-Z],,g')
export MAKE_HELPTEXT

HELM_CHANGELOG_HELPTEXT=$(./bin/helm-changelog -h)
export HELM_CHANGELOG_HELPTEXT

cat README.tmpl.md | gomplate > README.md
