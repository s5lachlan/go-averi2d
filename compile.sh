#!/bin/bash
echo "Compiling Linux build"
go build -ldflags "-s -w" -o build/go-averi2d.x86_64
echo "Compiling Windows build"
CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o build/go-averi2d.exe

