const fs = require('fs');
const { init } = require('z3-solver');

const paths = fs.readFileSync('./input.txt', 'utf8')
  .split('\r\n')
  .slice(0, -1)
  .map((line) => {
    const [coords, velocities] = line.split(' @ ');

    return [
      coords.split(', ').map(Number),
      velocities.split(', ').map(Number)
    ];
  });

async function run() {
  const { Context } = await init();
  const { Solver, Real, Eq } = Context('main');
  const solver = new Solver();

  const x = Real.const('x');
  const y = Real.const('y');
  const z = Real.const('z');
  const vx = Real.const('vx');
  const vy = Real.const('vy');
  const vz = Real.const('vz');

  for (let i = 0; i < paths.length; i++) {
    const [[xh, yh, zh], [vxh, vyh, vzh]] = paths[i];
    const t = Real.const(`t${i}`);

    solver.add(t.gt(0));
    solver.add(Eq(t.mul(vxh).add(xh), t.mul(vx).add(x)))
    solver.add(Eq(t.mul(vyh).add(yh), t.mul(vy).add(y)))
    solver.add(Eq(t.mul(vzh).add(zh), t.mul(vz).add(z)))
  };

  const res = await solver.check();

  if (res !== 'sat') {
    throw new Error('Unsat');
  }

  const model = solver.model();

  const xVal = model.get(x).value();
  const yVal = model.get(y).value();
  const zVal = model.get(z).value();

  console.log(xVal.numerator + yVal.numerator + zVal.numerator);
}

run();
