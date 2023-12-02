const fs = require('fs');

const input = fs.readFileSync('./input.txt', { encoding: 'utf8' }).split('\r\n').slice(0, -1);

const chars = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9'];
const words = ['zero', 'one', 'two', 'three', 'four', 'five', 'six', 'seven', 'eight', 'nine'];

let sum = 0;

for (const line of input) {
  let i = 0, j = line.length;
  let iFlag = false, jFlag = false;

  let firstNumber;
  let lastNumber;

  while (!iFlag || !jFlag) {
    const str = line.slice(i, j);

    for (let digit = 0; digit < 10; digit++) {
      if (!iFlag && (str.startsWith(chars[digit]) || str.startsWith(words[digit]))) {
        iFlag = true;
        firstNumber = digit;
      }

      if (!jFlag && (str.endsWith(chars[digit]) || str.endsWith(words[digit]))) {
        jFlag = true;
        lastNumber = digit;
      }
    }

    if (!iFlag) {
      i++;
    }

    if (!jFlag) {
      j--;
    }
  }

  sum += firstNumber * 10 + lastNumber;
}

console.log(`Sum: ${sum}`);
