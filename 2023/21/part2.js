const fs = require('fs');
const grid = fs.readFileSync('./input.txt', 'utf8')
  .split('\r\n')
  .slice(0, -1)
  .map((line) => line.split(''));

const WIDTH = grid[0].length;
const HEIGHT = grid.length;
const MAX_STEPS = 26501365;

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

function run(steps) {
  const dist = {
    [startingPosition.join('.')]: 0
  };

  const queue = [startingPosition];

  while (queue.length) {
    const [x, y] = queue.shift();
    const hash = `${x}.${y}`;

    const neighbours = [[0, 1], [0, -1], [1, 0], [-1, 0]]
      .map(([dx, dy]) => [x + dx, y + dy])
      .filter(([x, y]) => grid[x % HEIGHT < 0 ? x % HEIGHT + HEIGHT : x % HEIGHT][y % WIDTH < 0 ? y % WIDTH + WIDTH : y % WIDTH] === '.');

    for (const neighbour of neighbours) {
      const nHash = neighbour.join('.');

      if (!(nHash in dist)) {
        dist[nHash] = 1 + dist[hash];

        if (dist[nHash] !== steps) {
          queue.push(neighbour);
        }
      }
    }
  }

  return Object.values(dist).filter((value) => value % 2 === steps % 2).length;
}

// newton polynomial
const terms = Array.from({ length: 3 }).map((_, i) => run(startingPosition[0] + HEIGHT * i));

const a0 = terms[0];
const a1 = terms[1] - terms[0];
const a2 = (terms[2] - 2 * a1 - a0) / 2;

const finalTerm = Math.ceil(MAX_STEPS / HEIGHT);
const nx = a0 + a1 * (finalTerm - 1) + a2 * (finalTerm - 2) * (finalTerm - 1);

console.log(nx);
