package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Field int

const (
	EmptyField Field = iota
	StartField
	SplitField
	BeamField
	UnknownField
)

func fieldFromByte(input byte) Field {
	switch input {
	case "."[0]:
		return EmptyField
	case "S"[0]:
		return StartField
	case "|"[0]:
		return BeamField
	case "^"[0]:
		return SplitField
	default:
		log.Fatal("Uknown field value", input)
		return UnknownField
	}
}

func stringFromField(input Field) string {
	switch input {
	case EmptyField:
		return "."
	case StartField:
		return "S"
	case BeamField:
		return "|"
	case SplitField:
		return "^"
	default:
		log.Fatal("Uknown field", input)
		return "?"
	}
}

func printBoard(board [][]Field) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Print(stringFromField(board[i][j]))
		}
		fmt.Println()
	}
}

func computeBeam(x int, y int, cell Field, board [][]Field) Field {
	width := len(board[0])
	value := func(dx int, dy int) Field {
		if dx < 0 || dx >= width {
			return UnknownField
		}
		return board[dy][dx]
	}
	if cell == EmptyField {
		left := value(x-1, y)
		right := value(x+1, y)
		top := value(x, y-1)
		if top == BeamField || top == StartField {
			return BeamField
		} else if left == SplitField {
			if value(x-1, y-1) == BeamField {
				return BeamField
			}
		} else if right == SplitField {
			if value(x+1, y-1) == BeamField {
				return BeamField
			}
		}
		return EmptyField
	} else {
		return cell
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	board := make([][]Field, 0)
	for {
		// Read line and clean it
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		line2 := make([]Field, len(line))
		for i := 0; i < len(line); i++ {
			line2[i] = fieldFromByte(line[i])
		}
		board = append(board, line2)
	}
	// printBoard(board)

	// Compute beam
	board2 := make([][]Field, 0)
	board2 = append(board2, board[0])
	for i := 1; i < len(board); i++ {
		line := board[i]
		board3 := append(board2, line)
		line2 := make([]Field, len(line))
		for j := 0; j < len(line); j++ {
			line2[j] = computeBeam(j, i, line[j], board3)
		}
		board2 = append(board2, line2)
	}
	// printBoard(board2)

	// Calculate splits
	total := 0
	for i := 1; i < len(board2); i++ {
		for j := 0; j < len(board2[i]); j++ {
			if board2[i][j] == SplitField && board2[i-1][j] == BeamField {
				total++
			}
		}
	}

	fmt.Println("Done", total)
}
