const fs = require('fs');
const grid = fs.readFileSync('./input.txt', 'utf8')
  .split('\r\n')
  .slice(0, -1)
  .map((line) => line.split(''));

const WIDTH = grid[0].length;
const HEIGHT = grid.length;
const MAX_STEPS = 64;

let startingPosition = [];

for (let i = 0; i < HEIGHT; i++) {
  for (let j = 0; j < WIDTH; j++) {
    if (grid[i][j] === 'S') {
      startingPosition = [i, j];
      grid[i][j] = '.';
      break;
    }
  }
}

const dist = {
  [startingPosition.join('.')]: 0
};

const queue = [startingPosition];

while (queue.length) {
  const [x, y] = queue.shift();
  const hash = `${x}.${y}`;

  const neighbours = [[0, 1], [0, -1], [1, 0], [-1, 0]]
    .map(([dx, dy]) => [x + dx, y + dy])
    .filter(([x, y]) => x >= 0 && x < HEIGHT && y >= 0 && y < WIDTH && grid[x][y] === '.');

  for (const neighbour of neighbours) {
    const nHash = neighbour.join('.');

    if (!(nHash in dist)) {
      dist[nHash] = 1 + dist[hash];
      
      if (dist[nHash] !== MAX_STEPS) {
        queue.push(neighbour);
      }
    }
  }
}

dist[startingPosition.join('.')] = 2;

const sum = Object.values(dist).filter((value) => value % 2 === 0).length;

console.log(sum);
