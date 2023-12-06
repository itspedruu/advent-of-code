const fs = require('fs');
const lines = fs.readFileSync('./input.txt', 'utf8').split('\r\n');

const time = parseInt(lines[0].slice(5).trim().split(/\s+/).join(''))
const distance = parseInt(lines[1].slice(9).trim().split(/\s+/).join(''));

const roots = [
  (time - Math.sqrt(time ** 2 - 4 * distance)) / 2,
  (time + Math.sqrt(time ** 2 - 4 * distance)) / 2
];

const lessThanZero = [
  Math.floor(roots[0] + 1),
  Math.ceil(roots[1] - 1)
];

const total = lessThanZero[1] - lessThanZero[0] + 1;

console.log(`Total: ${total}`);
