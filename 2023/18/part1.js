const fs = require('fs');
const instructions = fs.readFileSync('./input.txt', 'utf8')
  .split('\r\n')
  .slice(0, -1)
  .map((line) => {
    const [direction, amount] = line.split(' ');

    return [direction, parseInt(amount)];
  });

const coord = [0, 0];
const nodes = [];

let perimeter = 0;

for (const [direction, amount] of instructions) {
  perimeter += amount;

  switch (direction) {
    case 'R': {
      coord[0] += amount;
      break;
    }
    case 'L': {
      coord[0] -= amount;
      break;
    }
    case 'U': {
      coord[1] += amount;
      break;
    }
    case 'D': {
      coord[1] -= amount;
      break;
    }
  }

  nodes.push([...coord]);
}

// Shoelace formula
const area = Math.abs(nodes.reduce((acc, [xi, yi], i) => acc + (xi * nodes[(i + 1) % nodes.length][1] - yi * nodes[(i + 1) % nodes.length][0]), 0)) / 2;

// Pick's theorem, b = perimeter
const interiorPoints = area - perimeter / 2 + 1;

console.log(perimeter + interiorPoints);
