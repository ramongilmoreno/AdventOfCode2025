#!/bin/bash

[ ! -f go.mod ] && go mod init AdventOfCode/`basename "${PWD}"`r
echo "Example" && cat ../03-1/example.txt | go run .
echo "Input" && cat ../03-1/input.txt | go run .
