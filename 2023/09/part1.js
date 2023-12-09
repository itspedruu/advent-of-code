const fs = require('fs');

const sequences = fs.readFileSync('./input.txt', 'utf8').split('\r\n').slice(0, -1).map((line) => line.split(' ').map(Number));

let sum = 0;

for (const sequence of sequences) {
  let lastValues = [sequence.at(-1)];
  let lastDifference = sequence.slice();

  while (lastDifference.some((value) => value !== 0)) {
    lastDifference = Array
      .from({ length: lastDifference.length - 1 })
      .map((_, index) => lastDifference[index + 1] - lastDifference[index]);

    lastValues.push(lastDifference.at(-1));
  }

  sum += lastValues.reduce((acc, value) => acc + value, 0);
}

console.log(`Sum: ${sum}`);
