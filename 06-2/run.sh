#!/bin/bash

[ ! -f go.mod ] && go mod init AdventOfCode/`basename "${PWD}"`
echo "Example" && cat ../06-1/example.txt | go run .
echo "Input" && cat ../06-1/input.txt | go run .
