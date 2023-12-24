const fs = require('fs');
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

const TEST_AREA = [200000000000000, 400000000000000];

let sum = 0;

for (let i = 0; i < paths.length - 1; i++) {
  for (let j = i + 1; j < paths.length; j++) {
    const [[x1, y1], [vx1, vy1]] = paths[i];
    const [[x2, y2], [vx2, vy2]] = paths[j];

    const t2 = ((y1 - y2) + vy1/vx1 * (x2 - x1)) / (vy2 - vy1 * vx2/vx1);
    const t1 = (x2 - x1 + vx2 * t2)/vx1;

    if (t1 < 0 || t2 < 0) {
      continue;
    } 

    const x = x2 + vx2 * t2;
    const y = y2 + vy2 * t2;

    if (x >= TEST_AREA[0] && x <= TEST_AREA[1] && y >= TEST_AREA[0] && y <= TEST_AREA[1]) {
      sum++;
    } 
  }
}

console.log(sum);
