const fs = require('fs');

const lines = fs.readFileSync('./input.txt', 'utf8').split('\r\n').slice(0, -1);

const WIDTH = lines[0].length;
const HEIGHT = lines.length;

let c = Array.from({ length: WIDTH }).fill(0);
let sum = 0;

for (let i = 0; i < HEIGHT; i++) {
  for (let j = 0; j < WIDTH; j++) {
    if (lines[i][j] === '.') {
      c[j]++;
    } else if (lines[i][j] === '#') {
      c[j] = 0;
    } else {
      sum += HEIGHT - i + c[j];
    }
  }
}

console.log(sum);
