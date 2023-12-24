const fs = require('fs');
const grid = fs.readFileSync('./input.txt', 'utf8')
  .split('\r\n')
  .slice(0, -1);

const WIDTH = grid[0].length;
const HEIGHT = grid.length;

const start = [0, grid[0].indexOf('.')];
const end = [HEIGHT - 1, grid.at(-1).indexOf('.')];

const visited = new Set();
visited.add(start.join('.'));

const queue = [[start, visited, 0]];
const nextDirections = {
  '>': [[0, 1]],
  '<': [[0, -1]],
  '^': [[-1, 0]],
  'v': [[1, 0]],
  '.': [[0, 1], [0, -1], [1, 0], [-1, 0]]
};

let max = 0;

while (queue.length) {
  const [[x, y], visited, len] = queue.pop(); 

  if (x === end[0] && y === end[1]) {
    max = Math.max(len, max);
    continue;
  }

  const neighbours = nextDirections[grid[x][y]]
    .map(([dx, dy]) => [x + dx, y + dy])
    .filter(([x, y]) => x >= 0 && x < HEIGHT && y >= 0 && y < WIDTH && grid[x][y] !== '#' && !visited.has(`${x}.${y}`));

  for (const neighbour of neighbours) {
    visited.add(neighbour.join('.'));

    queue.push([neighbour, new Set(visited), len + 1]);
  }
}

console.log(max);
