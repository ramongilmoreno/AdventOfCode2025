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

    start := position
    adjusted_amount:= amount + start
    if amount < 0 {
      if adjusted_amount <= 0 {
        if start != 0 {
          zeroes++
        }
        zeroes += -(adjusted_amount / modulus)
        position = ((adjusted_amount % modulus) + modulus ) % modulus
      } else {
        position = adjusted_amount
      }
    } else {
      zeroes += (adjusted_amount / modulus)
      position = adjusted_amount % modulus
    }

    // fmt.Println(start, amount, position, zeroes)
  }
  fmt.Println("Done", zeroes)
}

