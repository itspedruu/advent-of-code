// todo: optimize with different data structure for queue
const fs = require('fs');
const grid = fs.readFileSync('./input.txt', 'utf8')
  .split('\r\n')
  .slice(0, -1)
  .map((line) => line.split('').map(Number));

const WIDTH = grid[0].length;
const HEIGHT = grid.length;

// (coord, dir, numSteps)
const queue = [[[0, 0], [0, 1], 0]];
const dist = {
  '0.0.0.1.0': 0
};

function minIndex() {
  let index = 0;
  let minDist = Infinity;

  for (let i = 0; i < queue.length; i++) {
    const hash = [queue[i][0].join('.'), queue[i][1].join('.'), queue[i][2]].join('.');

    if (dist[hash] < minDist) {
      minDist = dist[hash];
      index = i;
    }
  }

  return index;
}

while (queue.length) {
  const index = minIndex();
  const [[x, y], [dx, dy], numSteps] = queue.splice(index, 1)[0];
  const hash = `${x}.${y}.${dx}.${dy}.${numSteps}`;

  const neighboursDirections = [[dx, dy], [dy, dx], [-dy, -dx]];

  for (let i = 0; i < neighboursDirections.length; i++) {
    const [ndx, ndy] = neighboursDirections[i];
    const [nx, ny] = [x + ndx, y + ndy];

    if (nx < 0 || nx >= HEIGHT || ny < 0 || ny >= WIDTH || (i === 0 && numSteps === 3)) {
      continue;
    }

    const nextNumSteps = i === 0 ? numSteps + 1 : 1;
    const nHash = `${nx}.${ny}.${ndx}.${ndy}.${nextNumSteps}`;

    if (!(nHash in dist)) {
      queue.push([[nx, ny], [ndx, ndy], nextNumSteps]);
      dist[nHash] = Infinity;
    }

    if (dist[hash] + grid[nx][ny] < dist[nHash]) {
      dist[nHash] = dist[hash] + grid[nx][ny]; 
    }
  }
}

const finalDistances = Object
  .keys(dist)
  .filter((key) => key.startsWith(`${HEIGHT - 1}.${WIDTH - 1}`))
  .map((key) => dist[key]);

console.log(Math.min(...finalDistances));
