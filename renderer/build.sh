#!/bin/bash

mkdir -p build/bin/public/images
cp -R ./frontend/public/. ./build/bin/public/

wails build
wails build -platform windows/amd64
