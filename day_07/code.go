package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type bag struct {
    contains map[string]int
}

const GOLD string = "shiny gold"
var graph = make(map[string]bag)

func FindGold(current bag) bool {
    if _, found := current.contains[GOLD]; found { return true }
    for key := range current.contains {
        if FindGold(graph[key]) { return true }
    }
    return false
}

func HowManyBags(current bag) int {
    total := 0
    for color, quantity := range current.contains {
        total += quantity + quantity*HowManyBags(graph[color])
    }
    return total
}

func main() {
    file, _ := os.Open("input.txt")
    reader := bufio.NewScanner(file)

    for reader.Scan() {
        line := reader.Text()
        separated := strings.Split(line, ", ")

        children := separated[1:]

        bagData := strings.Split(separated[0], " bags ")
        bagName := bagData[0]
        bagQuantity := bagData[1]

        if strings.Contains(bagQuantity, "no other") {
            graph[bagName] = bag{contains: nil}
            continue
        }

        childTree := make(map[string]int)
        separated = strings.Split(bagQuantity, " ")
        childTree[fmt.Sprintf("%s %s", separated[2], separated[3])], _ = strconv.Atoi(separated[1])

        for _, child := range children {
            separated = strings.Split(child, " ")
            childTree[fmt.Sprintf("%s %s", separated[1], separated[2])], _ = strconv.Atoi(separated[0])
        }

        graph[bagName] = bag{contains: childTree}
    }

    count := 0

    for _, bagData := range graph {
        if FindGold(bagData) { count++ }
    }

    fmt.Println(count)

    fmt.Println(HowManyBags(graph[GOLD]))
}