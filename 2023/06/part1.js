const fs = require('fs');
const lines = fs.readFileSync('./input.txt', 'utf8').split('\r\n');

const times = lines[0].slice(5).trim().split(/\s+/).map(Number);
const distances = lines[1].slice(9).trim().split(/\s+/).map(Number);

let total = 1;

for (let i = 0; i < times.length; i++) {
  const time = times[i];
  const distance = distances[i];

  const roots = [
    (time - Math.sqrt(time ** 2 - 4 * distance)) / 2,
    (time + Math.sqrt(time ** 2 - 4 * distance)) / 2
  ];

  const lessThanZero = [
    Math.floor(roots[0] + 1),
    Math.ceil(roots[1] - 1)
  ];

  total *= lessThanZero[1] - lessThanZero[0] + 1;
}

console.log(`Total: ${total}`);
