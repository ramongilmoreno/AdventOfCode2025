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

    max_digits := 12
    selection := make([]int, max_digits)
    // Chosen index of each position will be updated in the loop to find each
    // digit, so 0 initialization happens out of it
    chosen_index := 0
    for i := 0; i < max_digits; i++ {
      // Iterate over the available numbers at the front, leaving space for the
      // remaining values at the end
      for j := chosen_index; j < (length - max_digits + i + 1); j++ {
        if numbers[j] > numbers[chosen_index] {
          chosen_index = j
        }
      }
      selection[i] = numbers[chosen_index]
      chosen_index++
    }

    // fmt.Println(line, selection)
    value := 0
    for i := 0; i < max_digits; i++ {
      value = (value * 10) + selection[i]
    }
    total += value
  }
  fmt.Println("Done", total)
}

