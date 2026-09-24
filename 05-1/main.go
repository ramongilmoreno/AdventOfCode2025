package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type rng struct {
	first int
	last  int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	intervals := make([]rng, 0)
	values := make([]int, 0)
	in_values := false
	total := 0
	for {
		// Read line and clean it
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			in_values = true
			continue
		}

		if !in_values {
			values := strings.Split(line, "-")
			number1, errIn1 := strconv.Atoi(values[0])
			if errIn1 != nil {
				log.Fatal("Failed to parse input value #1", line, errIn1)
			}
			number2, errIn2 := strconv.Atoi(values[1])
			if errIn2 != nil {
				log.Fatal("Failed to parse input value #2", line, errIn2)
			}
			intervals = append(intervals, rng{first: number1, last: number2})
		} else {
			number3, errIn3 := strconv.Atoi(line)
			if errIn3 != nil {
				log.Fatal("Failed to parse input value #3", line, errIn3)
			}
			values = append(values, number3)
		}
	}

	// fmt.Println("Intervals", intervals)
	// fmt.Println("Values", values)

	for _, value := range values {
		for _, r := range intervals {
			if value >= r.first && value <= r.last {
				// Found fresh interval
				total++
				break
			}
		}
	}

	fmt.Println("Done", total)
}
