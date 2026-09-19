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
			if len%2 == 0 {
				left := candidate[0:(len / 2)]
				right := candidate[(len / 2):]
				if left == right {
					result += i
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
