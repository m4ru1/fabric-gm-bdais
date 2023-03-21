#!/bin/bash

source_dirs=()
while IFS=$'\n' read -r source_dir; do
    source_dirs+=("$source_dir")
done < <(go list -f '{{.Dir}}' ./... | sed s,"${fabric_dir}".,,g | cut -f 1 -d / | sort -u)

OUTPUT="$(goimports -l "${source_dirs[@]}")"
echo $OUTPUT
