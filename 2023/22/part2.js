const fs = require('fs');
const bricks = fs.readFileSync('./input.txt', 'utf8')
  .split('\r\n')
  .slice(0, -1)
  .map((line) => line.split('~').map((rawCoord) => rawCoord.split(',').map(Number)));

let maxX = 0, maxY = 0, maxZ = 0;

for (const [start, end] of bricks) {
  maxX = Math.max(start[0], end[0], maxX);
  maxY = Math.max(start[1], end[1], maxY);
  maxZ = Math.max(start[2], end[2], maxZ);
}

const grid = Array.from({ length: maxX + 1 })
  .map(() =>
    Array.from({ length: maxY + 1 }).map(() =>
      Array.from({ length: maxZ + 1 })
    )
  );

function setGridValues(start, end, value) {
  for (let x = start[0]; x <= end[0]; x++) {
    for (let y = start[1]; y <= end[1]; y++) {
      for (let z = start[2]; z <= end[2]; z++) {
        grid[x][y][z] = value;
      }
    }
  }  
}

for (let i = 0; i < bricks.length; i++) {
  setGridValues(...bricks[i], i);
}

let flag = false;

do {
  flag = false;

  for (let i = 0; i < bricks.length; i++) {
    const [[x1, y1, z1], [x2, y2, z2]] = bricks[i];

    let curZ = z1;
    let intersected = false;

    while (curZ > 1 && !intersected) {
      for (let x = x1; x <= x2 && !intersected; x++) {
        for (let y = y1; y <= y2; y++) {
          if (grid[x][y][curZ - 1] !== undefined) {
            intersected = true;
            break;  
          }
        }
      }

      if (!intersected) {
        curZ--;
      }
    }

    if (curZ !== z1) {
      flag = true;

      const newZ2 = z2 - (z1 - curZ);

      bricks[i][0][2] = curZ;
      bricks[i][1][2] = newZ2;

      setGridValues([x1, y1, newZ2 + 1], [x2, y2, z2], undefined);
      setGridValues(...bricks[i], i);
    }
  }
} while (flag);

const supportedBy = bricks.map(() => new Set());
const supports = bricks.map(() => new Set());

for (let i = 0; i < bricks.length; i++) {
  const [[x1, y1, z1], [x2, y2, _]] = bricks[i];

  if (z1 === 1) {
    supportedBy[i].add(i);
    continue;
  }

  for (let x = x1; x <= x2; x++) {
    for (let y = y1; y <= y2; y++) {
      if (grid[x][y][z1 - 1] !== undefined) {
        supportedBy[i].add(grid[x][y][z1 - 1]);
        supports[grid[x][y][z1 - 1]].add(i);
      }
    }
  }
}

const canNotBeRemoved = new Set();

for (let i = 0; i < bricks.length; i++) {
  for (const support of supports[i]) {
    if (supportedBy[support].size === 1) {
      canNotBeRemoved.add(i);
    }
  }
}

let sum = 0;

for (const i of canNotBeRemoved) {
  const queue = [i];
  const set = new Set();

  set.add(i);

  while (queue.length) {
    const j = queue.pop();

    for (const support of supports[j]) {
      if (set.has(support)) {
        continue;
      }

      let flag = true; // is going to fall

      for (const supportBy of supportedBy[support]) {
        if (!set.has(supportBy)) {
          flag = false;
          break;
        }
      }

      if (flag) {
        set.add(support);
        queue.push(support);
      }
    }
  }

  sum += set.size - 1;
}

console.log(sum);
