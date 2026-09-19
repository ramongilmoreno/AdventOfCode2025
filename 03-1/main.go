package main

import (
  "bufio"
  "fmt"
  "log"
  "strconv"
  "strings"
  "os"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  total := 0
  for {
    // Read line and clean it
    line, err := reader.ReadString('\n')
    if err != nil || len(line) == 0  {
      break
    }
    line = strings.TrimSpace(line)

    // Make an array of values
    length := len(line)
    numbers := make([]int, length)
    for i := 0; i < length; i++ {
      number, err2 := strconv.Atoi(line[i:i + 1])
      if err2 != nil {
        log.Fatal("Failed to parse input", line, i)
      }
      numbers[i] = number
    }

    // Start by the first two and try all possible combinations
    value := (numbers[0] * 10) + numbers[1]
    for i := 0; i < length - 1; i++ {
      for j := i + 1; j < length; j++ {
        value = max(value, (numbers[i] * 10) + numbers[j])
      }
    }

    total += value
    // fmt.Println(line, value)
  }
  fmt.Println("Done", total)
}

