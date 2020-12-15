package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)

func writeAllMemoryAddresses(mask string, address string, built string, value uint64, memory map[string]uint64) {
    if len(mask) == 0 {
        memory[built] = value
    } else {
        if mask[0] == 'X' {
            writeAllMemoryAddresses(mask[1:], address[1:], built + "0", value, memory)
            writeAllMemoryAddresses(mask[1:], address[1:], built + "1", value, memory)
        } else if mask[0] == '0'{
            writeAllMemoryAddresses(mask[1:], address[1:], built + string(address[0]), value, memory)
        } else {
            writeAllMemoryAddresses(mask[1:], address[1:], built + "1", value, memory)
        }
    }
}

func main() {
    file, _ := os.Open("input.txt")
    reader := bufio.NewScanner(file)
    memory := make(map[string]uint64)
    mask := ""
    r, _ := regexp.Compile(`^mem\[(\d+)] = (\d+)$`)
    // PART 1
    for reader.Scan() {
        line := reader.Text()
        if strings.HasPrefix(line, "mask") {
            mask = line[7:]
        } else {
            groups := r.FindStringSubmatch(line)
            address := groups[1]
            number := groups[2]
            intNumber, _ := strconv.Atoi(number)
            bitsNumber := fmt.Sprintf("%036b", intNumber)
            newNumber := ""
            for i := 0; i < 36; i++ {
                if mask[i] != 'X' { newNumber += string(mask[i]) } else { newNumber += string(bitsNumber[i]) }
            }
            // fmt.Printf("%s ~\n%s =\n%s\n\n", mask, bitsNumber, newNumber)
            newNumberInt, _ := strconv.ParseInt(newNumber, 2, 37)
            memory[address] = uint64(newNumberInt)
        }
    }

    sum1 := uint64(0)

    for _, value := range memory {
        sum1 += value
    }

    // PART 2
    _, _ = file.Seek(0, 0)
    reader = bufio.NewScanner(file)
    memory = make(map[string]uint64)

    for reader.Scan() {
        line := reader.Text()
        if strings.HasPrefix(line, "mask") {
            mask = line[7:]
        } else {
            groups := r.FindStringSubmatch(line)
            address := groups[1]
            number := groups[2]
            intAddress, _ := strconv.Atoi(address)
            bitsAddress := fmt.Sprintf("%036b", intAddress)
            intNumber, _ := strconv.Atoi(number)
            writeAllMemoryAddresses(mask, bitsAddress, "", uint64(intNumber), memory)
        }
    }

    sum2 := uint64(0)

    for _, value := range memory {
        sum2 += value
    }

    fmt.Println(sum1)
    fmt.Println(sum2)
}