const fs = require('fs');

const lines = fs.readFileSync('./input.txt', 'utf8').split('\r\n').slice(0, -1);
const cards = Object.keys(lines).reduce((acc, key) => ({...acc, [key]: 1 }), {});

for (const line of lines) {
  const [_, cardId, rest] = line.match(/Card\s+(\d+): (.*)/);
  const [winningNumbers, myNumbers] = rest.split(' | ').map((part) => part.trim().split(/\s+/));

  const matchingNumbers = myNumbers.filter((value) => winningNumbers.includes(value)).length;
  const cardIdNumber = Number.parseInt(cardId);
  const toAdd = cards[(cardIdNumber - 1).toString()];

  for (let nextCardId = cardIdNumber; nextCardId < cardIdNumber + matchingNumbers; nextCardId++) {
    cards[nextCardId.toString()] += toAdd;
  }
}

const sum = Object.values(cards).reduce((acc, value) => acc + value, 0);

console.log(`Sum: ${sum}`);
