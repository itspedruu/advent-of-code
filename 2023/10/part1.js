const fs = require('fs');
const lines = fs.readFileSync('./input.txt', 'utf8').split('\r\n').slice(0, -1).map((line) => line.split(''));

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
      queue.push([...neighbours, 1]);
    }
  }
}

let max = 0;

while (queue.length > 0) {
  const [left, right, steps] = queue.shift();

  lines[left[0]][left[1]] = '1';
  lines[right[0]][right[1]] = '1';

  if (left[0] === right[0] && left[1] === right[1] && steps > max) {
    max = steps;
  }

  const leftNeighbours = getConnectingNeighbours(...left);
  const rightNeighbours = getConnectingNeighbours(...right);

  for (const leftNeighbour of leftNeighbours) {
    for (const rightNeighbour of rightNeighbours) {
      queue.push([leftNeighbour, rightNeighbour, steps + 1]); 
    }
  }
}

console.log(`Max: ${max}`);
