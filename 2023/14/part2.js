const fs = require('fs');

let lines = fs.readFileSync('./input.txt', 'utf8').split('\r\n').slice(0, -1).map((line) => line.split(''));

const WIDTH = lines[0].length;
const HEIGHT = lines.length;
const CYCLES = 1_000_000_000;

let c = Array.from({ length: WIDTH }).fill(0);
let r = Array.from({ length: HEIGHT }).fill(0);

let memo = {};
let cycle = 0;
let hash;

for (; cycle < CYCLES; cycle++) {
  hash = lines.map((l) => l.join('')).join('');

  if (memo[hash]) {
    break;
  }
  
  for (const direction of ['N', 'W', 'S', 'E']) {
    let newLine = JSON.parse(JSON.stringify(lines));

    c.fill(0);
    r.fill(0);

    for (let i = 0; i < HEIGHT; i++) {
      for (let j = 0; j < WIDTH; j++) {
        const realI = direction === 'N' ? i : HEIGHT - 1 - i;
        const realJ = direction === 'W' ? j : WIDTH - 1 - j;

        if (lines[realI][realJ] === '.') {
          if (['N', 'S'].includes(direction)) {
            c[j]++;
          } else {
            r[i]++;
          }
        } else if (lines[realI][realJ] === '#') {
          if (['N', 'S'].includes(direction)) {
            c[j] = 0;
          } else {
            r[i] = 0;
          }
        } else {
          newLine[realI][realJ] = '.';

          switch (direction) {
            case 'N':
              newLine[realI - c[j]][realJ] = 'O';
              break;
            case 'W':
              newLine[realI][realJ - r[i]] = 'O';
              break;
            case 'S':
              newLine[realI + c[j]][realJ] = 'O';
              break;
            case 'E':
              newLine[realI][realJ + r[i]] = 'O';
              break;
          }
        }
      }
    }

    lines = newLine;
  }

  memo[hash] = lines;
}

const loop = [hash];

do {
  hash = memo[hash].map((line) => line.join('')).join('');
  loop.push(hash);
} while (loop[0] !== hash);

loop.splice(-1, 1);

lines = memo[loop[(CYCLES - cycle - 1) % loop.length]];

const sum = lines.reduce(
  (acc, line, i) => acc + line.filter((symbol) => symbol === 'O').length * (HEIGHT - i),
  0
);

console.log(sum);
