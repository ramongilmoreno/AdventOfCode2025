package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	values := make([][]int, 0)
	operations := make([]string, 0)
	total := 0
	for {
		// Read line and clean it
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		fields := strings.Split(line, " ")
		values2 := make([]int, 0)
		for i := 0; i < len(fields); i++ {
			if len(fields[i]) != 0 {
				switch fields[i] {
				case "+":
					operations = append(operations, "+")
				case "*":
					operations = append(operations, "*")
				default:
					number, err := strconv.Atoi(fields[i])
					if err != nil {
						log.Fatal("Failed to parse input value #1", line, err)
					}
					values2 = append(values2, number)
				}

			}
		}
		if len(operations) > 0 {
			break
		} else {
			values = append(values, values2)
		}
	}

	for i := 0; i < len(operations); i++ {
		acc := 0
		switch operations[i] {
		case "+":
			acc = 0
		case "*":
			acc = 1
		}
		for j := 0; j < len(values); j++ {
			switch operations[i] {
			case "+":
				acc += values[j][i]
			case "*":
				acc *= values[j][i]
			}
		}
		total += acc
	}

	// fmt.Println("Done", values, operations, total)
	fmt.Println("Done", total)
}
