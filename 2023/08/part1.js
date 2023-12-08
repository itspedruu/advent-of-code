const fs = require('fs');

const lines = fs.readFileSync('./input.txt', 'utf8').split('\r\n').slice(0, -1);

const instructions = lines[0].split('').map((char) => char === 'R' ? 1 : 0);
const adjs = {};

for (const line of lines.slice(2)) {
  const origin = line.slice(0, 3);
  const paths = line.slice(7, -1).split(', ');

  adjs[origin] = paths;
}

let i = 0;
let steps = 1;
let curNode = 'AAA';

while (adjs[curNode][instructions[i]] !== 'ZZZ') {
  curNode = adjs[curNode][instructions[i]];
  i++;
  steps++;

  if (i === instructions.length) {
    i = 0;
  }
}

console.log(`Steps: ${steps}`);
