const fs = require('fs');
const grid = fs.readFileSync('./input.txt', 'utf8').split('\r\n').slice(0, -1).map((line) => line.split(''));

let H = grid.length;
let W = grid[0].length;

const empty = {
  rows: Array.from({ length: H }).fill(true),
  cols: Array.from({ length: W }).fill(true)
};

for (let r = 0; r < H; r++) {
  for (let c = 0; c < W; c++) {
    if (grid[r][c] === '#') {
      empty.rows[r] = false;
      empty.cols[c] = false;
    }
  }
}

const galaxies = [];

let coord = [0, 0];

for (let r = 0; r < H; r++) {
  for (let c = 0; c < W; c++) {
    if (grid[r][c] === '#') {
      galaxies.push(JSON.parse(JSON.stringify(coord)));
    }

    coord[1] += empty.cols[c] ? 2 : 1;
  }

  coord[0] += empty.rows[r] ? 2 : 1;
  coord[1] = 0;
}

let sum = 0;

for (let i = 0; i < galaxies.length - 1; i++) {
  for (let j = i + 1; j < galaxies.length; j++) {
    const nodeA = galaxies[i];
    const nodeB = galaxies[j];

    sum += Math.abs(nodeA[0] - nodeB[0]) + Math.abs(nodeA[1] - nodeB[1]);
  }
}

console.log(`Sum: ${sum}`);
