#!/bin/bash

[ ! -f go.mod ] && go mod init AdventOfCode/`basename "${PWD}"`
echo "Example" && cat ../04-1/example.txt | go run .
echo "Input" && cat ../04-1/input.txt | go run .
