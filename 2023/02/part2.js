const fs = require('fs');
const input = fs.readFileSync('./input.txt', { encoding: 'utf8' }).split('\r\n').slice(0, -1);

let sum = 0;

for (const line of input) {
  const semicolonIndex = line.indexOf(':');
  const sets = line.slice(semicolonIndex + 2).split('; ');
  
  const max = {
    red: 1,
    green: 1,
    blue: 1
  }

  for (const set of sets) {
    const cubePairs = set.split(', ').map((str) => str.split(' '));

    for (const [rawNumber, color] of cubePairs) {
      const number = Number.parseInt(rawNumber);

      if (number > max[color]) {
        max[color] = number;
      } 
    }
  }

  sum += max.red * max.green * max.blue;
}

console.log(`Sum: ${sum}`);
