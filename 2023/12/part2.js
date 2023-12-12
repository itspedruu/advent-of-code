const fs = require('fs');

const groups = fs.readFileSync('./input.txt', 'utf8').split('\r\n').slice(0, -1).map((line) => {
  const [springs, rawOrder] = line.split(' ');

  // needs extra dot to jump to finish after reading last combination
  return [Array.from({ length: 5 }).fill(springs).join('?') + '.', Array.from({ length: 5 }).fill(rawOrder).join(',').split(',').map(Number)];
});

let memo = {};

function solve(group, springIdx, groupIdx) {
  const key = `${springIdx}.${groupIdx}`;

  if (groupIdx === group[1].length) {
    return springIdx < group[0].length && group[0].slice(springIdx).includes('#')
      ? 0
      : 1;
  }

  if (springIdx === group[0].length) {
    return 0;
  }

  // took a solid 10m to realize memo[key] could be zero
  // changed it to key in memo
  if (key in memo) {
    return memo[key];
  }

  const run = group[1][groupIdx];

  switch (group[0][springIdx]) {
    case '.':
      memo[key] = solve(group, springIdx + 1, groupIdx);
      break;
    case '#': {
      if (group[0].slice(springIdx, springIdx + run).includes('.') || group[0][springIdx + run] === '#') {
        memo[key] = 0;
      } else {
        memo[key] = solve(group, springIdx + run + 1, groupIdx + 1);
      }
      break;
    }
    case '?': {
      // move to the right, emulating combinations
      if (group[0].slice(springIdx, springIdx + run).includes('.') || group[0][springIdx + run] === '#') {
        memo[key] = solve(group, springIdx + 1, groupIdx);
      // last combination
      } else {
        memo[key] = solve(group, springIdx + 1, groupIdx)
          + solve(group, springIdx + run + 1, groupIdx + 1);
      }
      break;
    }
  }

  return memo[key];
}

let sum = 0;

for (const group of groups) {
  sum += solve(group, 0, 0);
  memo = {};
}

console.log(sum);
