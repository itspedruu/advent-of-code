const fs = require('fs');
const lines = fs.readFileSync('./input.txt', 'utf8').split('\r\n').slice(0, -1);

const modules = {}; // [label]: { type, emits }

for (const line of lines) {
  const [def, rawEmits] = line.split(' -> ');
  const emits = rawEmits.split(', ');

  if (def === 'broadcaster') {
    modules.broadcaster = { emits };
  } else {
    const type = def[0] === '%' ? 'FF' : 'C';
    const label = def.slice(1);

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

function run() {
  // (label, low - false; high - true)
  const queue = [['broadcaster', false, 'button']];
  const pulses = [0, 0];

  while (queue.length) {
    let [label, pulse, inputLabel] = queue.shift();
    const curModule = modules[label];

    pulses[Number(pulse)]++;

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

    for (const emit of curModule.emits) {
      queue.push([emit, pulse, label]);
    }
  }

  return pulses;
}

const pulses = [0, 0];

for (let i = 0; i < 1000; i++) {
  const result = run();

  pulses[0] += result[0];
  pulses[1] += result[1];
}

console.log(pulses[0] * pulses[1]);
