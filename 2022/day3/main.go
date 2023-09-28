package main

import (
    "fmt"
    "os"
    "strings"
    "unicode"
)

func main() {
    data, err := os.ReadFile("input.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }

    lines := strings.Split(string(data), "\n")
    sumPriorities := 0
    commonFound := false
    for i := 0; i < len(lines) - 1; i++ {
        compartmentOne := lines[i][:len(lines[i])/2]
        compartmentTwo := lines[i][len(lines[i])/2:]

        for _, charOne := range compartmentOne {
            for _, charTwo := range compartmentTwo {
                if charOne == charTwo {
                    if unicode.IsUpper(charOne) {
                        sumPriorities += int(charOne) - 38
                    } else {
                        sumPriorities += int(charOne) - 96
                    }
                    commonFound = true
                    break
                }
            }
            if commonFound {
                commonFound = false
                break
            }
        }
    }
    fmt.Println(sumPriorities)

    sumPriorities = 0
    commonFound = false
    for i := 0; i < len(lines) - 1; i += 3 {
        for _, charOne := range lines[i] {
            for _, charTwo := range lines[i+1] {
                for _, charThree := range lines[i+2] {
                    if charOne == charTwo && charTwo == charThree {
                        if unicode.IsUpper(charOne){
                            sumPriorities += int(charOne) - 38
                        } else {
                            sumPriorities += int(charOne) - 96
                        }
                        commonFound = true
                        break
                    }
                }
                if commonFound {
                    break
                }
            }
            if commonFound {
                commonFound = false
                break
            }
        }
    }
    fmt.Println(sumPriorities)
}
