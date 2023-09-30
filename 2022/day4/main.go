package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
)

func main() {
    fileContent, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    lines := strings.Split(string(fileContent), "\n")

    totalContained := 0
    totalOverlap := 0
    for i := 0; i < len(lines) - 1; i++ {
        sections := strings.Split(lines[i], ",")
        firstSection := strings.Split(sections[0], "-")
        secondSection := strings.Split(sections[1], "-")

        firstSectionStart, _ := strconv.Atoi(firstSection[0])
        firstSectionEnd, _ := strconv.Atoi(firstSection[1])
        secondSectionStart, _ := strconv.Atoi(secondSection[0])
        secondSectionEnd, _ := strconv.Atoi(secondSection[1])

        if(firstSectionStart <= secondSectionStart && firstSectionEnd >= secondSectionEnd) {
            totalContained++
        } else if(secondSectionStart <= firstSectionStart && secondSectionEnd >= firstSectionEnd) {
            totalContained++
        }
    

        if(firstSectionStart >= secondSectionStart && firstSectionStart <= secondSectionEnd) {
            totalOverlap++
        } else if(secondSectionStart >= firstSectionStart && secondSectionStart <= firstSectionEnd) {
            totalOverlap++
        }
        
    }
    fmt.Println(totalContained)
    fmt.Println(totalOverlap)
}
