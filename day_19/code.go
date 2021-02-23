package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "sync"
)

type rule struct {
    first  []int
    second []int
    letter uint8
}

var count int

func strToIntegers(str string) []int {
    splitStr := strings.Split(str, " ")
    vsm := make([]int, len(splitStr))
    for i, v := range splitStr {
        vsm[i], _ = strconv.Atoi(v)
    }
    return vsm
}

func checkString(s string, rules []rule, stack []int) bool {
    if len(stack) > len(s) {
        return false
    } else if len(stack) == 0 || len(s) == 0 {
        return len(stack) == len(s)
    }

    r := rules[stack[0]]
    newStack := stack[1:]
    if r.letter != 0 {
        if r.letter == s[0] {
            return checkString(s[1:], rules, newStack)
        }
    } else {
        if checkString(s, rules, append(r.first, newStack...)) {
            return true
        }
        if r.second != nil && checkString(s, rules, append(r.second, newStack...)) {
            return true
        }
    }
    return false
}

var mutex = &sync.Mutex{}

func main() {
    rules := make([]rule, 129)

    //file, _ := os.Open("input.txt")
    file, _ := os.Open("input_2.txt")
    reader := bufio.NewScanner(file)

    for reader.Scan() {
        line := reader.Text()
        if line == "" { break }
        lineParts := strings.Split(line, ": ")
        curNumber, _ := strconv.Atoi(lineParts[0])
        curRule := lineParts[1]
        if curRule == "\"a\"" || curRule == "\"b\"" {
            rules[curNumber] = rule{
                letter: curRule[1],
            }
        } else {
            curRules := strings.Split(curRule, " | ")
            if len(curRules) == 1 {
                rules[curNumber] = rule{
                    first: strToIntegers(curRules[0]),
                }
            } else {
                rules[curNumber] = rule{
                    first:  strToIntegers(curRules[0]),
                    second: strToIntegers(curRules[1]),
                }
            }
        }
    }

    count = 0
    var wg sync.WaitGroup

    for reader.Scan() {
        line := reader.Text()
        wg.Add(1)
        go func(s string, r []rule) {
           if valid := checkString(s, r, rules[0].first); valid {
               mutex.Lock()
               count++
               mutex.Unlock()
           }
           defer wg.Done()
        }(line, rules)
    }

    wg.Wait()
    fmt.Println(count)
}