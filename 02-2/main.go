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
	result := 0
	for {
		line, err := reader.ReadString(',')
		values := strings.Split(line, "-")
		from, err1 := strconv.Atoi(values[0])
		to, err2 := strconv.Atoi(strings.Replace(strings.TrimSpace(values[1]), ",", "", -1))
		if err1 != nil || err2 != nil {
			log.Fatal("Cannot parse integer", values)
			break
		}
		for i := from; i <= to; i++ {
			candidate := strconv.Itoa(i)
			len := len(candidate)

			// Test all possible sequences
			for j := 1; j <= (len / 2); j++ {
				master := candidate[0:j]
				k := j
				for ; k+j <= len; k += j {
					// fmt.Println("Testing", i, len, "master", master, k, "text", candidate[k:k + j])
					if master != candidate[k:k+j] {
						// Upon mismatch, go for next master
						break
					}
				}
				// fmt.Println("Out", "candidate", candidate, "master", master, "len", len, "k", k)
				if k == len {
					// fmt.Println("Adding", i, master)
					result += i

					// Do not continue testing sequences
					break
				}
			}
		}

		// Check if EOF was reached to break
		if err != nil {
			break
		}
	}
	fmt.Println("Done", result)
}
