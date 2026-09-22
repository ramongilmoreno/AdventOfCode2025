package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var board []string
	for {
		// Read line and clean it
		line, err := reader.ReadString('\n')
		if err != nil || len(line) == 0 {
			break
		}
		line = strings.TrimSpace(line)
		board = append(board, line)
	}
	// fmt.Println(board)
	width := len(board[0])
	height := len(board)
	occupied := func(x int, y int) bool {
		if x < 0 || x >= width || y < 0 || y >= height {
			return false
		} else {
			return board[y][x] == '@'
		}
	}
	v := func(x int, y int) int {
		if occupied(x, y) {
			return 1
		} else {
			return 0
		}
	}

	total := 0
	for {
		var newBoard []string
		subTotal := 0
		for i := 0; i < height; i++ {
			var newLine strings.Builder
			for j := 0; j < width; j++ {
				if occupied(j, i) {
					count := v(j - 1, i - 1) + v(j - 1, i) + v(j - 1, i + 1) + v(j, i - 1) + v(j, i + 1) + v(j + 1, i - 1) + v(j + 1, i) + v(j + 1, i + 1)
					if count < 4 {
						subTotal++
						newLine.WriteString(".")
					} else {
						newLine.WriteString("@")
					}
				} else {
					newLine.WriteString(".")
				}
			}
			newBoard = append(newBoard, newLine.String())
		}
		if subTotal == 0 {
			// Stop if no more movements
			break
		}
		// fmt.Println(newBoard)
		total += subTotal
		board = newBoard
	}
	fmt.Println("Done", total)
}
