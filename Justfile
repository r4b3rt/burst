#!/usr/bin/env just --justfile

run:
   go run .

# call go mod tidy
tidy:
    go mod tidy
