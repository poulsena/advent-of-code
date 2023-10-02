package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

func move(amount int, from int, to int, stacks [][]rune) [][]rune {
    for i := 0; i < amount; i++ {
        crate := stacks[from][0]
        stacks[from] = stacks[from][1:]
        stacks[to] = append([]rune{crate}, stacks[to]...)
    }

    return stacks
}

func main() {
    stacks := [][]rune{
        []rune("PDQRVBHF"),
        []rune("VWQZDL"),
        []rune("CPRGQZLH"),
        []rune("BVJFHDR"),
        []rune("CLWZ"),
        []rune("MVGTNPRJ"),
        []rune("SBMVLRJ"),
        []rune("JPD"),
        []rune("VWNCD"),
    }

    fileContent, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }

    lines := strings.Split(string(fileContent), "\n")

    for i := 0; i < len(lines) - 1; i++ {
        lineInfo := strings.Split(string(lines[i]), " ")

        amount, err := strconv.Atoi(lineInfo[1])
        if err != nil {
            panic(err)
        }

        from, err := strconv.Atoi(lineInfo[3])
        if err != nil {
            panic(err)
        }

        to, err := strconv.Atoi(lineInfo[5])
        if err != nil {
            panic(err)
        }

        stacks = move(amount, from - 1, to - 1, stacks)
    }

    for i := 0; i < len(stacks); i++ {
        fmt.Printf("%d: %s\n", i + 1, string(stacks[i][0]))
    }

}
