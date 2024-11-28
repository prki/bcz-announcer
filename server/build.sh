#!/bin/bash

mkdir -p ./build/bin
cp ./ggob.csv /build/bin/ggob.csv
cp ./players.db ./build/bin/players.db

wails build

wails build -platform windows/amd64
