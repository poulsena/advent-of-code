package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

func main() {
    data, err := os.ReadFile("input.txt")
    if err != nil {
        fmt.Println("File reading error", err)
    }
    
    lines := strings.Split(string(data), "\n")
    
    elfIndex := 0
    sumCal := 0
    highestCal := 0
    for i := 0; i < len(lines); i++ {
        if len(lines[i]) == 0 {
            sumCal = 0
            continue
        }

        num, err := strconv.Atoi(lines[i])
        if err != nil {
            fmt.Println("Error converting string to int", err)
        }

        sumCal += num
        if sumCal > highestCal {
            highestCal = sumCal
            elfIndex = i
        }
    }
    fmt.Println("Highest Calorie: ", highestCal)
    fmt.Println("Elf Index: ", elfIndex)
}
