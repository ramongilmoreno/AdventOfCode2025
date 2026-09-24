#!/bin/bash

[ ! -f go.mod ] && go mod init AdventOfCode/`basename "${PWD}"`
echo "Example" && cat example.txt | go run .
echo "Input" && cat input.txt | go run .
