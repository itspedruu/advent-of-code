const fs = require('fs');

const lines = fs.readFileSync('./input.txt', 'utf8').split('\r\n').slice(0, -1);

const instructions = lines[0].split('').map((char) => char === 'R' ? 1 : 0);
const adjs = {};

for (const line of lines.slice(2)) {
  const origin = line.slice(0, 3);
  const paths = line.slice(7, -1).split(', ');

  adjs[origin] = paths;
}

let curNodes = Object.keys(adjs).filter((node) => node.endsWith('A'));
let steps = curNodes.map(() => 0);

for (let index = 0; index < curNodes.length; index++) {
  let i = 0;

  do {
      curNodes[index] = adjs[curNodes[index]][instructions[i]];
      steps[index]++;
      i = (i + 1) % instructions.length;
  } while (!curNodes[index].endsWith('Z'))
}

function euclides_gcd(a, b) {
  let old_r = a, r = b, old_s = 1, s = 0, old_t = 0, t = 1;

  while (r !== 0) {
    const q = Math.floor(old_r / r);

    let temp_r = r, temp_s = s, temp_t = t;

    r = old_r - q * r, old_r = temp_r;
    s = old_s - q * s, old_s = temp_s;
    t = old_t - q * t, old_t = temp_t;
  }

  return old_r;
}

const totalSteps = steps.reduce((acc, cur) => acc * cur / euclides_gcd(acc, cur), 1);

console.log(`Total steps: ${totalSteps}`);
