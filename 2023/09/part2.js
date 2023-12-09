const fs = require('fs');

const sequences = fs.readFileSync('./input.txt', 'utf8').split('\r\n').slice(0, -1).map((line) => line.split(' ').map(Number));

let sum = 0;

for (const sequence of sequences) {
  let firstValues = [sequence[0]];
  let lastDifference = sequence.slice();

  while (lastDifference.some((value) => value !== 0)) {
    lastDifference = Array
      .from({ length: lastDifference.length - 1 })
      .map((_, index) => lastDifference[index + 1] - lastDifference[index]);

    firstValues.push(lastDifference[0]);
  }

  sum += firstValues.reduceRight((acc, value) => value - acc, 0);
}

console.log(`Sum: ${sum}`);
