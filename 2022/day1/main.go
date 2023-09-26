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
    topThreeElves := make(map[int]int)
    topThreeElves[0] = 0;
    topThreeElves[1] = 0;
    topThreeElves[2] = 0;
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

        if(sumCal > topThreeElves[0]) {
            topThreeElves[2] = topThreeElves[1]
            topThreeElves[1] = topThreeElves[0]
            topThreeElves[0] = sumCal
        }else if(sumCal > topThreeElves[1]) {
            topThreeElves[2] = topThreeElves[1]
            topThreeElves[1] = sumCal
        }else if(sumCal > topThreeElves[2]) {
            topThreeElves[2] = sumCal
        }
    }
    fmt.Println("Highest Calorie: ", highestCal)
    fmt.Println("Elf Index: ", elfIndex)

    fmt.Println("Top Three Elves: ", topThreeElves)
    topThreeSum := topThreeElves[0] + topThreeElves[1] + topThreeElves[2]
    fmt.Println("Top Three Sum: ", topThreeSum)
}
