package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Operation struct {
	start     int
	operation rune
	length    int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	values := make([][]int, 0)
	operations := make([]Operation, 0)
	total := 0

	// Parse input
	for {
		// Read line and clean it
		input, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		line := []rune(input)

		if !strings.Contains(input, "+") && !strings.Contains(input, "+") {
			// Parse values
			values2 := make([]int, 0)
			for i := 0; i < len(line); i++ {
				item := line[i]
				if item >= '0' && item <= '9' {
					values2 = append(values2, int(item-'0'))
				} else if item == ' ' {
					values2 = append(values2, -1)
				}
			}
			values = append(values, values2)
		} else {
			// Parse operations
			var last_operation rune
			var last_index int
			for i := 0; i < len(line); i++ {
				item := line[i]
				if item == '+' || item == '*' {
					if i != 0 {
						operations = append(operations, Operation{start: last_index, operation: last_operation, length: i - last_index - 1})
					}
					last_operation = item
					last_index = i
				} else if item == '\r' || item == '\n' {
					operations = append(operations, Operation{start: last_index, operation: last_operation, length: i - last_index})
				}
			}
			break
		}
	}

	// Perform operations
	for i := 0; i < len(operations); i++ {
		acc := 0
		switch operations[i].operation {
		case '+':
			acc = 0
		case '*':
			acc = 1
		}

		for j := operations[i].length - 1; j >= 0; j-- {
			value := 0
			dec := 1
			for k := len(values) - 1; k >= 0; k-- {
				digit := values[k][operations[i].start+j]
				if digit != -1 {
					value += (values[k][operations[i].start+j] * dec)
					dec *= 10
				}
			}
			// fmt.Println("Operand", value)
			switch operations[i].operation {
			case '+':
				acc += value
			case '*':
				acc *= value
			}
		}
		// fmt.Println("Result", acc)

		total += acc
	}

	// fmt.Println("Done", values, operations, total)
	fmt.Println("Done", total)
}
