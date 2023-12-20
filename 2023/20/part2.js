const fs = require('fs');
const lines = fs.readFileSync('./input.txt', 'utf8').split('\r\n').slice(0, -1);

const modules = {}; // [label]: { type, emits }

for (const line of lines) {
  const [def, rawEmits] = line.split(' -> ');
  const emits = rawEmits.split(', ');
  let label;

  if (def === 'broadcaster') {
    label = def;
    modules.broadcaster = { emits };
  } else {
    const type = def[0] === '%' ? 'FF' : 'C';
    label = def.slice(1);

    modules[label] = { type, emits };
  }
}

const context = {};

for (const [label, curModule] of Object.entries(modules)) {
  if (label === 'broadcaster') {
    continue;
  }

  for (const emit of curModule.emits) {
    if (modules[emit]?.type === 'C' && !(label in (context[emit] ??= {}))) {
      context[emit][label] = false;
    }
  }
}

const cycles = {};
const gate = Object.entries(modules)
  .find(([_, element]) => element.emits.includes('rx'))[0];

function run() {
  // (label, low - false; high - true)
  const queue = [['broadcaster', false, 'button']];

  while (queue.length) {
    let [label, pulse, inputLabel] = queue.shift();
    const curModule = modules[label];

    if (!curModule) {
      continue;
    }

    if (curModule.type === 'FF') {
      if (pulse) {
        continue;
      }

      const previousValue = (context[label] ??= false);

      context[label] = !previousValue;
      pulse = !previousValue;
    } else if (curModule.type === 'C') {
      context[label][inputLabel] = pulse;

      const allHigh = Object.values(context[label]).every((value) => value);
      
      pulse = !allHigh;
    }

    for (const [label, value] of Object.entries(context[gate])) {
      if (value && !(label in cycles)) {
        cycles[label] = i;
      }
    }

    for (const emit of curModule.emits) {
      queue.push([emit, pulse, label]);
    }
  }
}

let i = 0;

while (Object.keys(cycles).length !== 4) {
  i++;
  run();
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

const min = Object.values(cycles).reduce((acc, cur) => acc * cur / euclides_gcd(acc, cur), 1);

console.log(min);
