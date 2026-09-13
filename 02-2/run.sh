#!/bin/bash

[ ! -f go.mod ] && go mod init AdventOfCode/`basename "${PWD}"`
echo "Example" && cat ../02-1/example.txt | go run .
echo "Input" && cat ../02-1/input.txt | go run .
