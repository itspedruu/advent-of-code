const fs = require('fs');

const patterns = fs.readFileSync('./input.txt', 'utf8')
  .trim()
  .split('\r\n\r\n')
  .map((pattern) => pattern.split('\r\n').map((line) => line.split('')));

let sum = 0;

for (const pattern of patterns) {
  let i = 0, j = 0;
  let errors = 0;

  const width = pattern[0].length;
  const height = pattern.length;

  for (let n = 0; n < width - 1; n++) {
    errors = 0;

    for (const row of pattern) {
      i = n;
      j = i + 1; 

      while (i >= 0 && j < width) {
        if (row[i] !== row[j]) {
          errors++;
        }

        i--;
        j++;
      }
    }

    if (errors === 1) {
      sum += n + 1;
      break;
    }
  }

  for (let n = 0; n < height - 1; n++) {
    errors = 0;
    i = n;
    j = i + 1;

    while (i >= 0 && j < height) {
      errors += pattern[i].filter((col, idx) => pattern[j][idx] !== col).length;

      i--;
      j++; 
    }

    if (errors === 1) {
      sum += (n + 1) * 100;
      break;
    }
  }
}

console.log(sum);
