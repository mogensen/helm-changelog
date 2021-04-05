#!/bin/bash

mkdir -p charts

git clone git@github.com:elastic/helm-charts.git charts/elastic

(
    cd charts/elastic
    ../../bin/helm-changelog
)
