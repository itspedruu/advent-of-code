const fs = require('fs');

const patterns = fs.readFileSync('./input.txt', 'utf8')
  .trim()
  .split('\r\n\r\n')
  .map((pattern) => pattern.split('\r\n').map((line) => line.split('')));

let sum = 0;

for (const pattern of patterns) {
  let i = 0, j = 0;

  const width = pattern[0].length;
  const height = pattern.length;

  for (let n = 0; n < width - 1; n++) {
    for (const row of pattern) {
      i = n;
      j = i + 1; 

      while (i >= 0 && j < width && row[i] === row[j]) {
        i--;
        j++;
      }

      if (i >= 0 && j < width) {
        break;
      }
    }

    if (i < 0 || j >= width) {
      sum += n + 1;
      break;
    }
  }

  for (let n = 0; n < height - 1; n++) {
    i = n;
    j = i + 1;

    while (i >= 0 && j < height) {
      if (pattern[i].some((col, idx) => pattern[j][idx] !== col)) {
        break;
      }

      i--;
      j++;  
    }

    if (i < 0 || j >= height) {
      sum += (n + 1) * 100;
      break;
    }
  }
}

console.log(sum);
