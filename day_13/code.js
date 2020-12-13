let fs = require('fs')

function inv(N, m) {
    N %= m;
    for (let i = BigInt(1); i < m; i++) {
        if (N*i%m == 1) {
            return i;
        }
    }
}

fs.readFile('input.txt', 'utf8', (_, data) => {
    let dataArray = data.toString().split('\n');
    let initialTime = Number(dataArray[0]);
    let buses = dataArray[1].split(',');
    let minTime = initialTime;
    let minBus = initialTime;
    let N = BigInt(1);
    let ri = [], mi = [];
    buses.forEach((value, index) => {
        if (value != 'x') {
            let busNumber = Number(value);
            if ((timeLeft = busNumber - (initialTime%busNumber)) < minTime) {
                minBus = busNumber;
                minTime = timeLeft;
            }
            ri.push(BigInt(index));
            mi.push(BigInt(busNumber));
            N *= BigInt(busNumber);
        }
    });
    console.log(minTime*minBus);

    let Ni = [];

    mi.forEach(value => Ni.push(N/value));

    let result = BigInt(0);

    for (let i = 0; i < mi.length; i++) result += ri[i]*inv(Ni[i], mi[i])*Ni[i];

    console.log(N - (result%N));
});