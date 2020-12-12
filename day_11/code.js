let fs = require('fs');

fs.readFile('input.txt', (_, data) => {
    let lines = data.toString().split('\n');
    lines.forEach((line, i, newLines) => {
        newLines[i] = line.split('');
        newLines[i].forEach((seat, j, seats) => {
            if (seat === 'L') seats[j] = '#';
        });
    });
    lines.pop();
    let lines2 = JSON.parse(JSON.stringify(lines));

    let sizeX = lines.length;
    let sizeY = lines[0].length;
    let changed = true;
    let changed2 = true;
    let times = 0;
    let padding = [[-1, -1], [0, -1], [-1, 0], [1, -1], [-1, 1], [1, 0], [0, 1], [1, 1]]

    while (changed || changed2) {
        // console.log(lines.map(value => value.join('')).join('\n'));
        // console.log('\n');
        changed = false;
        changed2 = false;
        times += 1;
        let newLines = JSON.parse(JSON.stringify(lines));
        let newLines2 = JSON.parse(JSON.stringify(lines2));
        for (let i = 0; i < sizeX; i += 1) {
            for (let j = 0; j < sizeY; j += 1) {
                let seat = lines[i][j];
                let seat2 = lines2[i][j];
                if (seat != '.') {
                    let occupied = 0;
                    padding.forEach(pad => {
                        let nI = i + pad[0];
                        let nJ = j + pad[1];
                        if (nI >= 0 && nI <= sizeX - 1 && nJ >= 0 && nJ <= sizeY - 1 && lines[nI][nJ] == '#') { occupied += 1; }
                    });
                    if (seat == 'L' && occupied == 0) {
                        newLines[i][j] = '#';
                        changed = true;
                    }
                    if (seat == '#' && occupied >= 4) {
                        newLines[i][j] = 'L';
                        changed = true;
                    }
                }
                if (seat2 != '.') {
                    let occupied = 0;
                    padding.forEach(pad => {
                        let nI = i;
                        let nJ = j;
                        do {
                            nI += pad[0];
                            nJ += pad[1];
                        } while (nI >= 0 && nI <= sizeX - 1 && nJ >= 0 && nJ <= sizeY - 1 && lines2[nI][nJ] == '.')
                        if (nI >= 0 && nI <= sizeX - 1 && nJ >= 0 && nJ <= sizeY - 1 && lines2[nI][nJ] == '#') { occupied += 1; }
                    });
                    if (seat2 == 'L' && occupied == 0) {
                        newLines2[i][j] = '#';
                        changed2 = true;
                    }
                    if (seat2 == '#' && occupied >= 5) {
                        newLines2[i][j] = 'L';
                        changed2 = true;
                    }
                }
            }
        }
        lines = newLines;
        lines2 = newLines2;
    }

    let count = 0;
    let count2 = 0;

    lines.forEach(row => {
        row.forEach(seat => { if (seat == '#') { count += 1; } });
    });
    lines2.forEach(row => {
        row.forEach(seat => { if (seat == '#') { count2 += 1; } });
    });

    console.log(count);
    console.log(count2);
});
