const fs = require('fs');
const grid = fs.readFileSync('./input.txt', 'utf8').split('\r\n').slice(0, -1);

const WIDTH = grid[0].length;
const HEIGHT = grid.length;

function run(startingCoord, startingDirection) {
  const tiles = {};
  const beams = [[startingCoord, startingDirection]];

  while (beams.length) {
    for (let i = 0; i < beams.length; i++) {
      const beam = beams[i];
      const hash = beam[0].join('.');
      const directionHash = beam[1].join('.');
      const outOfBounds = beam[0][0] < 0 || beam[0][0] >= HEIGHT || beam[0][1] < 0 || beam[0][1] >= WIDTH;

      if (tiles[hash]?.includes?.(directionHash) || outOfBounds) {
        beams.splice(i, 1);
        i--;
        continue;
      }

      (tiles[hash] ??= []).push(directionHash);

      const tile = grid[beam[0][0]][beam[0][1]];

      switch (tile) {
        case '\\':
          beam[1].reverse(); // reverse direction
          break;
        case '/':
          beam[1] = beam[1].reverse().map((value) => -value); // reverse dir and mult by -1
          break;
        case '|': {
          if (beam[1][1] === 0) {
            break;
          }

          const newBeams = [
            [[beam[0][0] + 1, beam[0][1]], [1, 0]],
            [[beam[0][0] - 1, beam[0][1]], [-1, 0]]
          ];

          beams.splice(i, 1, ...newBeams);
          i++;
          continue;
        }
        case '-': {
          if (beam[1][0] === 0) {
            break;
          }

          const newBeams = [
            [[beam[0][0], beam[0][1] + 1], [0, 1]],
            [[beam[0][0], beam[0][1] - 1], [0, -1]]
          ];

          beams.splice(i, 1, ...newBeams);
          i++;
          continue;
        }
      }

      beam[0][0] += beam[1][0];
      beam[0][1] += beam[1][1];
    }
  }

  return Object.keys(tiles).length;
}

let max = 0;
const startingDirections = [[1, 0], [0, -1], [-1, 0], [0, 1]]; // t, r, b, l

for (let i = 0; i < startingDirections.length; i++) {
  let coord, direction;

  if (i === 0) {
    coord = [0, 0];
    direction = [0, 1];
  } else if (i === 1) {
    coord = [0, WIDTH - 1];
    direction = [1, 0];
  } else if (i === 2) {
    coord = [HEIGHT - 1, 0];
    direction = [0, 1];
  } else {
    coord = [0, 0];
    direction = [1, 0];
  }

  while (coord[0] >= 0 && coord[0] < HEIGHT && coord[1] >= 0 && coord[1] < WIDTH) {
    const sum = run(
      JSON.parse(JSON.stringify(coord)),
      JSON.parse(JSON.stringify(startingDirections[i]))
    );

    if (sum > max) {
      max = sum;
    }

    coord[0] += direction[0];
    coord[1] += direction[1];
  }

}

console.log(max);
