const fs = require('fs');
const input = fs.readFileSync('./input.txt', { encoding: 'utf8' }).split('\r\n').slice(0, -1);

const MAX_CUBES = {
  red: 12,
  green: 13,
  blue: 14
}

let sum = 0;

for (const line of input) {
  const semicolonIndex = line.indexOf(':');
  const gameId = Number.parseInt(line.slice(5, semicolonIndex));
  const sets = line.slice(semicolonIndex + 2).split('; ');

  let valid = true;

  for (const set of sets) {
    const cubePairs = set.split(', ').map((str) => str.split(' '));

    for (const [rawNumber, color] of cubePairs) {
      const number = Number.parseInt(rawNumber);

      if (number > MAX_CUBES[color]) {
        valid = false;
        break;
      } 
    }

    if (!valid) {
      break;
    }
  }

  if (valid) {
    sum += gameId;
  }
}

console.log(`Sum: ${sum}`);
