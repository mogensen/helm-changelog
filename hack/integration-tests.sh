#!/bin/bash

mkdir -p charts
mkdir -p examples

git clone git@github.com:elastic/helm-charts.git charts/elastic || true

(
    cd charts/elastic
    git pull
    ../../bin/helm-changelog create
)

for changelog in $(find . -name 'Changelog.md' | sed "s|./charts/elastic/||g"); do
    dirName=$(dirname $changelog)
    mv ./charts/elastic/$changelog examples/${dirName}.md
done
