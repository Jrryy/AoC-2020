package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type move struct {
    x int
    y int
    z int
}

type move4D struct {
    x int
    y int
    z int
    w int
}

var MOVES = [26]move{
    {0, 0, 1},
    {0, 0, -1},
    {0, 1, 0},
    {0, 1, 1},
    {0, 1, -1},
    {0, -1, 0},
    {0, -1, 1},
    {0, -1, -1},
    {1, 0, 0},
    {1, 0, 1},
    {1, 0, -1},
    {1, 1, 0},
    {1, 1, 1},
    {1, 1, -1},
    {1, -1, 0},
    {1, -1, 1},
    {1, -1, -1},
    {-1, 0, 0},
    {-1, 0, 1},
    {-1, 0, -1},
    {-1, 1, 0},
    {-1, 1, 1},
    {-1, 1, -1},
    {-1, -1, 0},
    {-1, -1, 1},
    {-1, -1, -1},
}

var MOVES4D = [80]move4D{
    {-1, -1, -1, -1},
    {-1, -1, -1, 0},
    {-1, -1, -1, 1},
    {-1, -1, 0, -1},
    {-1, -1, 0, 0},
    {-1, -1, 0, 1},
    {-1, -1, 1, -1},
    {-1, -1, 1, 0},
    {-1, -1, 1, 1},
    {-1, 0, -1, -1},
    {-1, 0, -1, 0},
    {-1, 0, -1, 1},
    {-1, 0, 0, -1},
    {-1, 0, 0, 0},
    {-1, 0, 0, 1},
    {-1, 0, 1, -1},
    {-1, 0, 1, 0},
    {-1, 0, 1, 1},
    {-1, 1, -1, -1},
    {-1, 1, -1, 0},
    {-1, 1, -1, 1},
    {-1, 1, 0, -1},
    {-1, 1, 0, 0},
    {-1, 1, 0, 1},
    {-1, 1, 1, -1},
    {-1, 1, 1, 0},
    {-1, 1, 1, 1},
    {0, -1, -1, -1},
    {0, -1, -1, 0},
    {0, -1, -1, 1},
    {0, -1, 0, -1},
    {0, -1, 0, 0},
    {0, -1, 0, 1},
    {0, -1, 1, -1},
    {0, -1, 1, 0},
    {0, -1, 1, 1},
    {0, 0, -1, -1},
    {0, 0, -1, 0},
    {0, 0, -1, 1},
    {0, 0, 0, -1},
    {0, 0, 0, 1},
    {0, 0, 1, -1},
    {0, 0, 1, 0},
    {0, 0, 1, 1},
    {0, 1, -1, -1},
    {0, 1, -1, 0},
    {0, 1, -1, 1},
    {0, 1, 0, -1},
    {0, 1, 0, 0},
    {0, 1, 0, 1},
    {0, 1, 1, -1},
    {0, 1, 1, 0},
    {0, 1, 1, 1},
    {1, -1, -1, -1},
    {1, -1, -1, 0},
    {1, -1, -1, 1},
    {1, -1, 0, -1},
    {1, -1, 0, 0},
    {1, -1, 0, 1},
    {1, -1, 1, -1},
    {1, -1, 1, 0},
    {1, -1, 1, 1},
    {1, 0, -1, -1},
    {1, 0, -1, 0},
    {1, 0, -1, 1},
    {1, 0, 0, -1},
    {1, 0, 0, 0},
    {1, 0, 0, 1},
    {1, 0, 1, -1},
    {1, 0, 1, 0},
    {1, 0, 1, 1},
    {1, 1, -1, -1},
    {1, 1, -1, 0},
    {1, 1, -1, 1},
    {1, 1, 0, -1},
    {1, 1, 0, 0},
    {1, 1, 0, 1},
    {1, 1, 1, -1},
    {1, 1, 1, 0},
    {1, 1, 1, 1},
}

func main() {
    // Create empty cube and hypercube
    cubes := make([][][]rune, 15)
    copyCubes := make([][][]rune, 15)
    cubes4D := make([][][][]rune, 15)
    copyCubes4D := make([][][][]rune, 15)
    for i := range cubes {
        cubes[i] = make([][]rune, 22)
        copyCubes[i] = make([][]rune, 22)
        cubes4D[i] = make([][][]rune, 15)
        copyCubes4D[i] = make([][][]rune, 15)
        for j := range cubes[i] {
            cubes[i][j] = []rune(strings.Repeat(".", 22))
            copyCubes[i][j] = []rune(strings.Repeat(".", 22))
        }
        for j := range cubes4D[i] {
            cubes4D[i][j] = make([][]rune, 22)
            copyCubes4D[i][j] = make([][]rune, 22)
            for k := range cubes4D[i][j] {
                cubes4D[i][j][k] = []rune(strings.Repeat(".", 22))
                copyCubes4D[i][j][k] = []rune(strings.Repeat(".", 22))
            }
        }
    }

    // Read the line and place the plane in the center of the cube and the hypercube
    file, _ := os.Open("input.txt")
    reader := bufio.NewScanner(file)
    i := 7
    for reader.Scan() {
        line := reader.Text()
        cubes[7][i] = []rune(strings.Repeat(".", 7) + line + strings.Repeat(".", 7))
        cubes4D[7][7][i] = []rune(strings.Repeat(".", 7) + line + strings.Repeat(".", 7))
        i++
    }

    // Do the iterations for both the cube and the hypercube
    for turn := 0; turn < 6; turn++ {
        // First the cube
        // fmt.Println("Turn ", turn)
        for i = 1; i < 14; i++ {
            // fmt.Println("z = ", i)
            for j := 1; j < 21; j++ {
                for k := 1; k < 21; k++ {
                    cube := cubes[i][j][k]
                    actives := 0
                    for _, m := range MOVES {
                        if cubes[i + m.x][j + m.y][k + m.z] == '#' {
                            actives++
                        }
                    }
                    if cube == '#' && actives != 2 && actives != 3 {
                        copyCubes[i][j][k] = '.'
                    } else if cube == '.' && actives == 3 {
                        copyCubes[i][j][k] = '#'
                    } else {
                        copyCubes[i][j][k] = cube
                    }
                    // if i == 5 && j == 11 && k == 13 {
                    //     fmt.Println(actives)
                    // }
                }
                // fmt.Println(string(copyCubes[i][j]))
            }
        }
        cubes = make([][][]rune, 15)
        for i := range cubes {
            cubes[i] = make([][]rune, 22)
            for j := range cubes[i] {
                cubes[i][j] = make([]rune, 22)
                copy(cubes[i][j],  copyCubes[i][j])
            }
        }

        // Then the hypercube
        for i = 1; i < 14; i++ {
            for j := 1; j < 14; j++ {
                for k := 1; k < 21; k++ {
                    for w := 1; w < 21; w++ {
                        cube := cubes4D[i][j][k][w]
                        actives := 0
                        for _, m := range MOVES4D {
                            if cubes4D[i+m.x][j+m.y][k+m.z][w + m.w] == '#' {
                                actives++
                            }
                        }
                        if cube == '#' && actives != 2 && actives != 3 {
                            copyCubes4D[i][j][k][w] = '.'
                        } else if cube == '.' && actives == 3 {
                            copyCubes4D[i][j][k][w] = '#'
                        } else {
                            copyCubes4D[i][j][k][w] = cube
                        }
                    }
                }
            }
        }
        cubes4D = make([][][][]rune, 15)
        for i := range cubes4D {
            cubes4D[i] = make([][][]rune, 15)
            for j := range cubes4D[i] {
                cubes4D[i][j] = make([][]rune, 22)
                for k := range cubes4D[i][j] {
                    cubes4D[i][j][k] = make([]rune, 22)
                    copy(cubes4D[i][j][k],  copyCubes4D[i][j][k])
                }
            }
        }
    }

    count := 0
    for _, plane := range cubes {
        for _, line := range plane {
            count += strings.Count(string(line), "#")
        }
    }

    fmt.Println(count)

    count = 0
    for _, cube := range cubes4D {
        for _, plane := range cube {
            for _, line := range plane {
                count += strings.Count(string(line), "#")
            }
        }
    }

    fmt.Println(count)
}
