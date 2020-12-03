package main

import (
	"bufio"
	"fmt"
	"os"
)

func calculateTrees(treeMap []string, right int, down int) (count int) {
	column := 0
	skip := down

	for _, row := range treeMap {
		if skip == down {
			if row[column] == '#' {count++}
			column = (column + right)%31
			skip = 0
		}
		skip++
	}

	return count
}

func main() {
	file, _ := os.Open("input.txt")
	reader := bufio.NewScanner(file)

	var treeMap []string
	for reader.Scan() {
		treeMap = append(treeMap, reader.Text())
	}

	part1 := calculateTrees(treeMap, 3,1)
	part2 := part1*
		calculateTrees(treeMap, 1,1)*
		calculateTrees(treeMap, 5,1)*
		calculateTrees(treeMap, 7,1)*
		calculateTrees(treeMap, 1,2)

	fmt.Println(part1)
	fmt.Println(part2)
}