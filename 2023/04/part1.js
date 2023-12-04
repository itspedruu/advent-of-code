const fs = require('fs');

const lines = fs.readFileSync('./input.txt', 'utf8').split('\r\n').slice(0, -1);

const sum = lines.reduce((acc, line) => {
  const [_, rest] = line.split(': ');
  const [winningNumbers, myNumbers] = rest.split(' | ').map((part) => part.trim().split(/\s+/));

  const matchingNumbers = myNumbers.filter((value) => winningNumbers.includes(value)).length;
  const points = matchingNumbers === 0 ? 0 : 2 ** (matchingNumbers - 1);

  return acc + points;
}, 0);

console.log(`Sum: ${sum}`);
