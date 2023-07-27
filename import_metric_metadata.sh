#!/bin/bash

mkdir metric-metadata

# get the latest branch, e.g. "v0.70.0"
branch=$(curl https://proxy.golang.org/github.com/open-telemetry/opentelemetry-collector-contrib/@latest | jq -r .Version)

cd /tmp
rm -rf opentelemetry-collector-contrib
git clone --depth 1 --branch $branch https://github.com/open-telemetry/opentelemetry-collector-contrib
cd -

find /tmp/opentelemetry-collector-contrib -name metadata.yaml -exec sh -c '
  filename="{}"
  parentdir="$(dirname "$filename")"
  parentdirname="$(basename "$parentdir")"
  cp "${filename}" "metric-metadata/${parentdirname}.yaml"
' \;

rm -rf /tmp/opentelemetry-collector-contrib
