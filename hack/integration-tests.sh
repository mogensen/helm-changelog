#!/bin/bash

mkdir -p charts
mkdir -p examples

git clone git@github.com:prometheus-community/helm-charts.git charts/prometheus-community || true
(
    cd charts/prometheus-community/
    git pull
    ../../bin/helm-changelog -v info
)

rm examples/*.md

for changelog in $(find charts -name 'Changelog.md'); do
    echo $changelog
    dirName=$(dirname $changelog)
    cp $changelog examples/$(basename ${dirName}).md
done
