const fs = require('fs');
const strings = fs.readFileSync('./input.txt', 'utf8').trim().split(',');
const sum = strings.reduce((sum, string) =>
    sum + string.split('').reduce((acc, symbol) => (acc + symbol.charCodeAt(0)) * 17 % 256, 0)
, 0);

console.log(sum);
