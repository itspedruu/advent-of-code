const fs = require('fs');

const groups = fs.readFileSync('./input.txt', 'utf8').split('\r\n').slice(0, -1).map((line) => {
  const [springs, rawOrder] = line.split(' ');

  return [springs, rawOrder.split(',').map(Number)];
});

function extensions(candidate, group) {
  const nextSpring = group[0][candidate.length];

  return nextSpring === '?' ? ['.', '#'] : [nextSpring];
}

function valid(candidate, group) {
  let continuous = 0;
  let j = -1;

  for (const symbol of candidate) {
    if (symbol === '#') {
      if (continuous === 0) {
        j++;
      }

      if (j === group[1].length) {
        return false;
      }

      if (continuous === group[1][j]) {
        return false;
      }

      continuous++;
    } else if (continuous > 0) {
      if (continuous !== group[1][j]) {
        return false;
      }

      continuous = 0;
    }
  }

  return j === group[1].length - 1 && (continuous === 0 || continuous === group[1][j]);
}

function solve(candidate, group) {
  if (candidate.length === group[0].length) {
    return valid(candidate, group) ? 1 : 0;
  }

  let sum = 0;

  for (const extension of extensions(candidate, group)) {
    candidate.push(extension);
    sum += solve(candidate, group);
    candidate.pop();
  }

  return sum;
}

let sum = 0;

for (const group of groups) {
  sum += solve([], group);
}

console.log(sum);
