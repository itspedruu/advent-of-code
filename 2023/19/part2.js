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

const queue = [['in', { x: [1, 4000], m: [1, 4000], a: [1, 4000], s: [1, 4000] }]];
const approvedContexts = [];

while (queue.length) {
  const [label, context] = queue.pop();

  if (label === 'A') {
    approvedContexts.push(context);
    continue;
  } else if (label === 'R') {
    continue;
  }

  const steps = functions[label];

  let nextContext = JSON.parse(JSON.stringify(context));

  for (const step of steps) {
    const newContext = JSON.parse(JSON.stringify(nextContext));

    if (step.operation) {
      if (step.operation === '>') {
        newContext[step.variable][0] = Math.max(step.value + 1, newContext[step.variable][0]);         
        nextContext[step.variable][1] = Math.min(step.value, newContext[step.variable][1]);
      } else {
        newContext[step.variable][1] = Math.min(step.value - 1, newContext[step.variable][1]);
        nextContext[step.variable][0] = Math.max(step.value, newContext[step.variable][0]);         
      }
    }

    queue.push([step.jump, newContext]);
  }
}

const sum = approvedContexts.reduce((acc, context) =>
  acc + Object.values(context).reduce((prev, [min, max]) => prev * (max - min + 1), 1)
, 0);

console.log(sum);
