const fs = require('fs');

const games = fs.readFileSync('./input.txt', 'utf8')
  .split('\r\n')
  .slice(0, -1)
  .map((line) => [line.slice(0, 5), parseInt(line.slice(6))]);

const CARD_STRENGTH = ['A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'];

function getHandOrder(hand) {
  const cards = {};

  for (const card of hand) {
    if (card in cards) {
      cards[card]++;
    } else {
      cards[card] = 1;
    }
  }

  const uniqueCards = Object.keys(cards).sort((a, b) => cards[b] - cards[a]);

  switch (uniqueCards.length) {
    case 1:
      return 7; // five of a kind
    case 2:
      return cards[uniqueCards[0]] === 4
        ? 6 // four of a kind
        : 5; // full house
    case 3:
      return cards[uniqueCards[0]] === 3
        ? 4 // three of a kind
        : 3; // two of a kind
    case 4:
      return 2; // one pair
    case 5:
      return 1; // high card
  }
}

function getFirstNonMatchingDistance(handA, handB) {
  let i = 0;

  while (handA[i] === handB[i]) {
    i++;
  }

  return CARD_STRENGTH.indexOf(handA[i]) - CARD_STRENGTH.indexOf(handB[i]);
}

games.sort((a, b) => {
  const handOrderComparison = getHandOrder(a[0]) - getHandOrder(b[0]);
  
  return handOrderComparison === 0 ? getFirstNonMatchingDistance(b[0], a[0]) : handOrderComparison;
});

const sum = games.reduce((acc, game, index) => acc + game[1] * (index + 1), 0);

console.log(`Sum: ${sum}`);
