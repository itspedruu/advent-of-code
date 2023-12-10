const fs = require('fs');
const lines = fs.readFileSync('./input.txt', 'utf8').split('\r\n').slice(0, -1).map((line) => line.split(''));
const originalLines = JSON.parse(JSON.stringify(lines));

const N = lines.length;

const queue = [];

function getConnectingNeighbours(i, j) {
  const valid = {
    '|': [[i - 1, j], [i + 1, j]],
    '-': [[i, j - 1], [i, j + 1]],
    'L': [[i, j - 1], [i + 1, j]],
    'J': [[i, j + 1], [i + 1, j]],
    '7': [[i, j + 1], [i - 1, j]],
    'F': [[i, j - 1], [i - 1, j]]
  };

  return [[i - 1, j], [i + 1, j], [i, j - 1], [i, j + 1]]
    .filter(([x, y]) => {
      if (x < 0 || x >= N || y < 0 || y >= N) {
        return false;
      }

      if (!(lines[x][y] in valid)) {
        return false;
      }

      return valid[lines[x][y]].some(([x1, y1]) => x1 === x && y === y1);
    });
}

for (let i = 0; i < N; i++) {
  for (let j = 0; j < N; j++) {
    if (lines[i][j] === 'S') {
      const neighbours = getConnectingNeighbours(i, j);
      
      lines[i][j] = '1';
      queue.push([...neighbours, [[i, j], ...neighbours]]);
    }
  }
}

let max = 0;
let maxPath = [];

while (queue.length > 0) {
  const [left, right, path] = queue.shift();
  const steps = (path.length - 1) / 2;

  lines[left[0]][left[1]] = '1';
  lines[right[0]][right[1]] = '1';

  if (left[0] === right[0] && left[1] === right[1] && steps > max) {
    max = steps; 
    maxPath = path.slice(0, -1);
  }

  const leftNeighbours = getConnectingNeighbours(...left);
  const rightNeighbours = getConnectingNeighbours(...right);

  for (const leftNeighbour of leftNeighbours) {
    for (const rightNeighbour of rightNeighbours) {
      queue.push([leftNeighbour, rightNeighbour, [...path, leftNeighbour, rightNeighbour]]); 
    }
  }
}

// bool matrix
const grid = Array.from({ length: N }).map(() => Array.from({ length: N }).fill('0'));

for (const [i, j] of maxPath) {
  grid[i][j] = '1';
}

let area = 0;

for (let i = 0; i < N; i++) {
  let count = 0;

  for (let j = 0; j < N; j++) {
    if (grid[i][j] === '0' && count % 2 === 1) {
      area++;
    // S counts for my input because it has a vertical bar LS
    } else if (grid[i][j] === '1' && ['|', 'L', 'J', 'S'].includes(originalLines[i][j])) {
      count++;
    }
  }
}

console.log(`Area: ${area}`);
