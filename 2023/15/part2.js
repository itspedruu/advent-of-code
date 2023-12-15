const fs = require('fs');
const strings = fs.readFileSync('./input.txt', 'utf8').trim().split(',');

const boxes = Array.from({ length: 256 }).map(() => []);

for (const string of strings) {
  const isDashOperation = string.at(-1) === '-';

  const label = string.slice(0, isDashOperation ? -1 : -2);
  const operation = isDashOperation ? '-' : '=';
  const lens = isDashOperation ? null : parseInt(string.at(-1));
  const boxId = label.split('').reduce((acc, symbol) => (acc + symbol.charCodeAt(0)) * 17 % 256, 0);
  const index = boxes[boxId].findIndex((box) => box[0] === label);

  if (operation === '=') {
    if (index === -1) {
      boxes[boxId].push([label, lens]);
    } else {
      boxes[boxId][index][1] = lens;
    }
  } else if (index > -1) {
    boxes[boxId].splice(index, 1);
  }
}


const sum = boxes.reduce((sum, box, i) =>
  sum + box.reduce((acc, item, slot) => acc + (i + 1) * item[1] * (slot + 1), 0)
, 0);

console.log(sum);
