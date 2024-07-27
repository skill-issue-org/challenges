#!/bin/bash

rm -rf bin
mkdir -p bin

src="dist"

for filepath in "$src"/*/*; do
    filename=$(basename "$filepath")
    dirname=$(dirname "$filepath")
    dirname=$(basename "$dirname")

    out="bin/${dirname}"
    if [[ "$filepath" == *"windows"* ]]; then
        out+=".exe"
    fi

    cp --preserve=all "$filepath" "$out"
done