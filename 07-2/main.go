package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	EmptyField int = -iota
	StartField
	SplitField
	UnknownField
)

type Board [][]int

func fieldFromByte(input byte) int {
	switch input {
	case "."[0]:
		return EmptyField
	case "S"[0]:
		return StartField
	case "^"[0]:
		return SplitField
	default:
		log.Fatal("Uknown field value", input)
		return UnknownField
	}
}

func stringFromField(input int) string {
	switch input {
	case EmptyField:
		return "."
	case StartField:
		return "S"
	case SplitField:
		return "^"
	default:
		return strconv.Itoa(input)
	}
}

func printBoard(board Board) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Print(stringFromField(board[i][j]))
		}
		fmt.Println()
	}
}

func computeValues(x int, y int, board Board, valuesLine []int, valuesBoard Board) {
	width := len(board[0])
	value := func(dx int, dy int) int {
		if dx < 0 || dx >= width || dy < 0 {
			return UnknownField
		}
		return valuesBoard[dy][dx]
	}
	addValue := func(dx int, dy int, increment int) {
		if dx < 0 || dx >= width || dy < 0 {
			// Do nothing
		}
		valuesLine[dx] += increment
	}
	if board[y-1][x] == StartField {
		addValue(x, y, 1)
	} else {
		v := value(x, y-1)
		if v > 0 {
			v2 := board[y][x]
			if v2 == EmptyField {
				addValue(x, y, v)
			} else if v2 == SplitField {
				if x >= 1 && x < width {
					addValue(x-1, y, v)
					addValue(x+1, y, v)
				}
			}
		}
	}

	if board[y][x] == SplitField {
	} else if board[y][x] == SplitField {
		v := value(x, y-1)
		if v > 0 && x >= 1 && x < width {
			addValue(x-1, y, v)
			addValue(x+1, y, v)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	board := make([][]int, 0)
	for {
		// Read line and clean it
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		line2 := make([]int, len(line))
		for i := 0; i < len(line); i++ {
			line2[i] = fieldFromByte(line[i])
		}
		board = append(board, line2)
	}
	// printBoard(board)

	// Compute beam values
	board2 := make([][]int, 0)
	board2 = append(board2, make([]int, len(board[0])))
	for i := 1; i < len(board); i++ {
		line := board[i]
		// line2 is initialized to 0
		line2 := make([]int, len(line))
		for j := 0; j < len(line); j++ {
			computeValues(j, i, board, line2, board2)
		}
		for j := 0; j < len(line); j++ {
			if board[i][j] != EmptyField {
				line2[j] = board[i][j]
			}
		}
		board2 = append(board2, line2)
	}
	// printBoard(board2)

	// Compute result
	total := 0
	last := len(board) - 1
	for i := 0; i < len(board[last]); i++ {
		v := board2[last][i]
		if v >= 0 {
			total += v
		}
	}

	fmt.Println("Done", total)
}
