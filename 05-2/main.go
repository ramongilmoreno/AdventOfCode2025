package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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
		}
	}
	// fmt.Println("Intervals", intervals)

	slices.SortFunc(intervals,
		func(a, b rng) int {
			r := a.first - b.first
			if r == 0 {
				r = a.last - b.last
			}
			return r
		})

	// fmt.Println("Sorted", intervals)

	// Merge intervals
	merged := make([]rng, 0)
	var last *rng
	for _, r := range intervals {
		if last == nil {
			last = &r
		} else {
			if r.first >= last.first && r.first <= last.last {
				last = &rng{first: last.first, last: max(last.last, r.last)}
			} else {
				merged = append(merged, *last)
				last = &r
			}
		}
	}
	if last != nil {
		merged = append(merged, *last)
	}
	// fmt.Println("Merged", merged)

	// Count ranges sizes
	for _, r := range merged {
		total += (r.last - r.first + 1)
	}

	fmt.Println("Done", total)
}
