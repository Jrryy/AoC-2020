var adapters = [0]

for input in CommandLine.arguments[1...] {
    adapters.append(Int(input)!)
}

adapters.sort()
adapters.append(adapters.last! + 3)

var ones = 0
var threes = 0
var total = 1
var sequence = [0, 0, 1, 1]
var current = 2

for i in 1...(adapters.count - 1) {
    if adapters[i] == adapters[i - 1] + 1 {
        ones += 1
        current += 1
        if current == sequence.count {
            sequence.append(sequence[current - 1]*2 - sequence[current - 4])
        }
    }
    else {
        threes += 1
        total *= sequence[current]
        current = 2
    }
}

print(ones*threes)
print(total)
