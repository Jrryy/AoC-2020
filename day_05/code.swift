func computeNumber(array: ArraySlice<Character>, char: Character) -> Int {
    var number = 0
    for c in array {
        number = number << 1
        if c == char { number += 1 }
    }
    return number
}

var max_1 = 0
var ids: Array<Int> = []

for input in CommandLine.arguments[1...] {
    let row = computeNumber(array: Array(input)[0...6], char: "B")
    let col = computeNumber(array: Array(input)[7...9], char: "R")
    let new = row*8 + col
    ids.append(new)
    if new > max_1 { max_1 = new }
}

print(max_1)

ids.sort()

for (index, id) in ids.enumerated() {
    if ids[index + 1] != id + 1 {
        print(id + 1)
        break
    }
}
