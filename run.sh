#!/bin/bash

# clean output
rm main.pdf

# generate syntax highlighted snippets
rm -rf gen
python minted.py

# replace german umlaute
sed -i 's/ß/{\\ss}/g' content/*

# compile latex
tectonic main.tex
