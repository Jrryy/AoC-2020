let fs = require('fs');

fs.readFile('input.txt', (_, data) => {
    let instructions = data.toString().split('\n');
    let visited = [];
    let swappable = [];

    let counter = 0;
    let pointer = 0;

    while (!visited.includes(pointer)) {
        let instruction = instructions[pointer].slice(0, 3);
        let quantity = Number(instructions[pointer].slice(4));

        visited.push(pointer);
        
        switch (instruction) {
            case 'acc':
                pointer += 1;
                counter += quantity;
                break;
            case 'jmp':
                swappable.push(pointer);
                pointer += quantity;
                break;
            default:
                if (quantity != 0 && !visited.includes(pointer + quantity)) swappable.push(pointer);
                pointer += 1;
                break;
        }
    }

    console.log(counter);

    swappable.some(toSwap => {
        visited = [];

        counter = 0;
        pointer = 0;

        let original = instructions[toSwap].slice(0, 3)

        if (original == 'nop') instructions[toSwap] = instructions[toSwap].replace('nop', 'jmp');
        else instructions[toSwap] = instructions[toSwap].replace('jmp', 'nop');

        while (!visited.includes(pointer) && pointer < instructions.length){
            let instruction = instructions[pointer].slice(0, 3)
            let quantity = Number(instructions[pointer].slice(4))
            
            visited.push(pointer);
    
            switch (instruction) {
                case 'acc':
                    pointer += 1;
                    counter += quantity;
                    break;
                case 'jmp':
                    pointer += quantity;
                    break;
                default:
                    pointer += 1;
                    break;
            }
        }

        if (pointer == instructions.length) {
            console.log(counter);
            return true;
        }
        if (original == 'nop') instructions[toSwap] = instructions[toSwap].replace('jmp', 'nop');
        else instructions[toSwap] = instructions[toSwap].replace('nop', 'jmp');
        return false
    });
});