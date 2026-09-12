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
  position := 50
  modulus := 100
  zeroes := 0
  for {
    line, err := reader.ReadString('\n')
    line = strings.TrimSpace(line)
    if err != nil || len(line) == 0  {
      break
    }
    orientation := line[0:1]
    amount, err2  := strconv.Atoi(line[1:])
    if err2 != nil {
      log.Fatal(err2)
    }
    // fmt.Println(orientation, amount)
    switch orientation {
    case "L":
      amount = -amount
    case "R":
    default:
      log.Fatal("Unknown direction", orientation)
    }

    // Make amount positive in the 0-modulus range
    amount = amount % modulus
    amount += modulus

    // Add to the current position, also keeping within the 0-modulus range
    position += amount
    position = position % modulus
    // fmt.Println(position)

    if position == 0 {
      zeroes++
    }

  }
  fmt.Println("Done", zeroes)
}

