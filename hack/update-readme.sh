#!/bin/bash

MAKE_HELPTEXT=$(make -s help | sed -r "s/\x1B\[([0-9]{1,3}(;[0-9]{1,2})?)?[mGK]//g")
export MAKE_HELPTEXT

HELM_CHANGELOG_HELPTEXT=$(./bin/helm-changelog -h)
export HELM_CHANGELOG_HELPTEXT

cat README.tmpl.md | gomplate > README.md
