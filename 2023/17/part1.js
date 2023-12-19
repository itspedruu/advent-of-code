const fs = require('fs');
const grid = fs.readFileSync('./input.txt', 'utf8')
  .split('\r\n')
  .slice(0, -1)
  .map((line) => line.split('').map(Number));

const WIDTH = grid[0].length;
const HEIGHT = grid.length;

// (coord, dir, numSteps)
const heap = [[[0, 0], [0, 1], 0]];
const dist = {
  '0.0.0.1.0': 0
};

const getLeftIndex = (index) => index * 2 + 1;
const getRightIndex = (index) => index * 2 + 2;
const getParentIndex = (index) => Math.floor((index - 1) / 2);

function compare(i, j) {
  if (i < 0 || i >= heap.length || j < 0 || j >= heap.length) {
    return false;
  }

  const [[xI, yI], [dxI, dyI], numStepsI] = heap[i];
  const hashI = `${xI}.${yI}.${dxI}.${dyI}.${numStepsI}`;

  const [[xJ, yJ], [dxJ, dyJ], numStepsJ] = heap[j];
  const hashJ = `${xJ}.${yJ}.${dxJ}.${dyJ}.${numStepsJ}`;

  return dist[hashI] < dist[hashJ];
}

function swap(i, j) {
  const temp = heap[i];

  heap[i] = heap[j];
  heap[j] = temp;
}

function extractMin() {
  const min = heap[0];

  heap[0] = heap[heap.length - 1];
  heap.pop();

  let index = 0;

  while (getLeftIndex(index) < heap.length) {
    let smallerIndex = getLeftIndex(index);

    if (compare(getRightIndex(index), smallerIndex)) {
      smallerIndex = getRightIndex(index);
    }

    if (compare(index, smallerIndex)) {
      break;
    }

    swap(index, smallerIndex);
    index = smallerIndex;
  }

  return min;
}

function addValue(value) {
  heap.push(JSON.parse(JSON.stringify(value)));
  
  let index = heap.length - 1;

  while (compare(index, getParentIndex(index))) {
    swap(getParentIndex(index), index);
    index = getParentIndex(index);
  }
}

while (heap.length) {
  const [[x, y], [dx, dy], numSteps] = extractMin();
  const hash = `${x}.${y}.${dx}.${dy}.${numSteps}`;

  if (hash.startsWith(`${HEIGHT - 1}.${WIDTH - 1}`)) {
    continue;
  }

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
      addValue([[nx, ny], [ndx, ndy], nextNumSteps]);
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
