const fs = require('fs');
const grid = fs.readFileSync('./input.txt', 'utf8')
  .split('\r\n')
  .slice(0, -1);

const WIDTH = grid[0].length;
const HEIGHT = grid.length;

const start = [0, grid[0].indexOf('.')];
const end = [HEIGHT - 1, grid.at(-1).indexOf('.')];

const junctions = [start, end];

for (let x = 0; x < HEIGHT; x++) {
  for (let y = 0; y < WIDTH; y++) {
    if (grid[x][y] === '#') {
      continue;
    }

    const count = [[0, 1], [0, -1], [1, 0], [-1, 0]]
      .map(([dx, dy]) => [x + dx, y + dy])
      .filter(([x, y]) => x >= 0 && x < HEIGHT && y >= 0 && y < WIDTH && grid[x][y] !== '#')
      .length;

    if (count >= 3) {
      junctions.push([x, y]);
    }
  }
}

const edges = {};

for (const junction of junctions) {
  const queue = [[junction, 0]];
  const visited = new Set();
  visited.add(junction.join('.'));

  while (queue.length) {
    const [[x, y], steps] = queue.pop();

    if (steps > 0 && junctions.some(([jx, jy]) => jx === x && jy === y)) {
      (edges[junction.join('.')] ??= {})[`${x}.${y}`] = steps;
      continue;
    }

    const neighbours = [[0, 1], [0, -1], [1, 0], [-1, 0]]
      .map(([dx, dy]) => [x + dx, y + dy])
      .filter(([nx, ny]) => nx >= 0 && nx < HEIGHT && ny >= 0 && ny < WIDTH && grid[nx][ny] !== '#' && !visited.has(`${nx}.${ny}`));

    for (const neighbour of neighbours) {
      visited.add(neighbour.join('.'));
      queue.push([neighbour, steps + 1]);
    }
  }
}

const visited = new Set();

function dfs(node) {
  if (node[0] === end[0]) {
    return 0;
  }

  let max = 0;

  visited.add(node.join('.'));

  for (const nHash of Object.keys(edges[node.join('.')])) {
    const neighbour = nHash.split('.').map(Number); 

    if (!visited.has(nHash)) {
      const len = dfs(neighbour) + edges[node.join('.')][nHash];

      max = Math.max(max, len);
    }
  }

  // backtracking
  visited.delete(node.join('.'));

  return max;
}

console.log(dfs(start));
