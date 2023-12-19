const fs = require('fs');
const [workflows, rawInputs] = fs.readFileSync('./input.txt', 'utf8')
  .trim()
  .split('\r\n\r\n')
  .map((block) => block.split('\r\n'));

const functions = {};

for (const workflow of workflows) {
  const [label, rawSteps] = workflow.slice(0, -1).split('{');
  const steps = [];

  for (const rawStep of rawSteps.split(',')) {
    if (rawStep.includes(':')) {
      const [statement, jump] = rawStep.split(':');

      steps.push({
        variable: statement[0],
        operation: statement[1],
        value: parseInt(statement.slice(2)),
        jump
      });
    } else {
      steps.push({ jump: rawStep });
    }
  }

  functions[label] = steps;
}

const inputs = [];

for (const rawInput of rawInputs) {
  const input = {};

  for (const assign of rawInput.slice(1, -1).split(',')) {
    const variable = assign[0];
    const value = parseInt(assign.slice(2));

    input[variable] = value; 
  }

  inputs.push(input);
}

let sum = 0;

for (const input of inputs) {
  let label = 'in';

  while (!['A', 'R'].includes(label)) {
    const steps = functions[label];

    for (const step of steps) {
      const jumps = !step.operation
        || step.operation === '>' && input[step.variable] > step.value
        || step.operation === '<' && input[step.variable] < step.value;

      if (jumps) {
        label = step.jump;
        break;
      }
    }
  }

  if (label === 'A') {
    sum += Object.values(input).reduce((acc, value) => acc + value, 0);
  }
}

console.log(sum);
