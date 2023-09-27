package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    data, err := os.ReadFile("input.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }

    lines := strings.Split(string(data), "\n")
    lines = lines[:len(lines)-1]

    sumPoints := 0
        for i := 0; i < len(lines); i++ {
            instructions := strings.Split(lines[i], " ")

        instPoint := 0
        gamePoint := 0
        switch instructions[1] {
        case "X":
            instPoint = 1
        case "Y":
            instPoint = 2
        case "Z":
            instPoint = 3
        }
        // Opponent chooses rock
        if instructions[0] == "A" {
            // You choose rock
            if instructions[1] == "X" {
                gamePoint = 3
            }/* You choose paper */ else if instructions[1] == "Y" {
                gamePoint = 6
            }/* You choose scissors */ else {
                gamePoint = 0
            }
        }/* Opponent chooses paper */ else if instructions[0] == "B" {
            // You choose rock
            if instructions[1] == "X" {
                gamePoint = 0
            }/* You choose paper */ else if instructions[1] == "Y" {
                gamePoint = 3
            }/* You choose scissors */ else {
                gamePoint = 6
            }
        }/* Opponent chooses scissors */ else {
            // You choose rock
            if instructions[1] == "X" {
                gamePoint = 6
            }/* You choose paper */ else if instructions[1] == "Y" {
                gamePoint = 0
            }/* You choose scissors */ else {
                gamePoint = 3
            }
        }
        sumPoints += gamePoint + instPoint
    }

    fmt.Println("Sum of points part one: ", sumPoints)

    sumPoints = 0
    for i := 0; i < len(lines); i++ {
        instructions := strings.Split(lines[i], " ")

        instPoint := 0
        gamePoint := 0
        switch instructions[0] {
        // Oponent chooses rock
        case "A":
            switch instructions[1] {
                // You need to lose -> choose scissors
            case "X":
                instPoint = 3
                gamePoint = 0
                // You need to tie -> choose rock
            case "Y":
                instPoint = 1
                gamePoint = 3
                // You need to win -> choose paper
            case "Z":
                instPoint = 2
                gamePoint = 6
            }
        // Oponent chooses paper
        case "B":
            switch instructions[1] {
                // You need to lose -> choose rock
            case "X":
                instPoint = 1
                gamePoint = 0
                // You need to tie -> choose paper
            case "Y":
                instPoint = 2
                gamePoint = 3
                // You need to win -> choose scissors
            case "Z":
                instPoint = 3
                gamePoint = 6
            }
        // Oponent chooses scissors
        case "C":
            switch instructions[1] {
                // You need to lose -> choose paper
            case "X":
                instPoint = 2
                gamePoint = 0
                // You need to tie -> choose scissors
            case "Y":
                instPoint = 3
                gamePoint = 3
                // You need to win -> choose rock
            case "Z":
                instPoint = 1
                gamePoint = 6
            }
        }
        sumPoints += gamePoint + instPoint
    }

    fmt.Println("Sum of points part two: ", sumPoints)
}
