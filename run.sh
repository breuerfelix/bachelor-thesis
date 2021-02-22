#!/bin/bash

rm main.pdf
rm -rf gen
python minted.py
tectonic main.tex
