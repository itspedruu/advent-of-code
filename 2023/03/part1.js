const fs = require('fs');

const lines = fs.readFileSync('./input.txt', { encoding: 'utf8'}).split('\r\n');

let N = lines.length, M = lines[0].length;

// hash -> (number, adjacents)
const numbers = {};

let numberStr = '';
let adjacents = [];

for (let i = 0; i < lines.length; i++) {
  for (let j = 0; j < lines[i].length; j++) {
    let charCode = lines[i].charCodeAt(j);

    if (charCode >= 48 && charCode <= 57) {
      numberStr += lines[i][j];
      adjacents.push(`${i}.${j}`);
    } else {
      const number = Number.parseInt(numberStr);

      for (const hash of adjacents) {
        numbers[hash] = [number, adjacents];
      }

      numberStr = '';
      adjacents = [];
    }
  }
}

function getValidNeighbourHashes(line, column) {
 return [
    [line - 1, column - 1],
    [line - 1, column],
    [line - 1, column + 1],
    [line, column - 1],
    [line, column + 1],
    [line + 1, column - 1],
    [line + 1, column],
    [line + 1, column + 1]
  ].filter(([n, m]) => n >= 0 && n < N && m >= 0 && m < M)
   .map(([n, m]) => `${n}.${m}`);
}

let visited = [];
let sum = 0;

for (let i = 0; i < lines.length; i++) {
  for (let j = 0; j < lines[i].length; j++) {
    let charCode = lines[i].charCodeAt(j);

    if ((charCode < 48 || charCode > 57) && charCode !== 46) {
      const neighbours = getValidNeighbourHashes(i, j);

      for (const hash of neighbours) {
        if (hash in numbers && !visited.includes(hash)) {
          sum += numbers[hash][0];
          visited.push(...numbers[hash][1]);
        }
      }
    }
  }
}

console.log(`Sum: ${sum}`);
