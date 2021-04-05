#!/bin/bash

mkdir -p charts

git clone git@github.com:elastic/helm-charts.git charts/elastic || true

(
    cd charts/elastic
    git pull
    ../../bin/helm-changelog create
)
