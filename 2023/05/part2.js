const fs = require('fs');
const lines = fs.readFileSync('./input.txt', 'utf8').split('\r\n');

// name -> (destStart, sourceStart, range)
const maps = {};
let currentMapName = '';

const initialNumbers = lines[0].slice(7).split(' ').map(Number);
const seeds = Array.from({ length: initialNumbers.length / 2 })
  .fill()
  .map((_, index) => initialNumbers.slice(index * 2, index * 2 + 2))
  .map(([start, range]) => [start, start + range - 1]);

for (const line of lines.slice(2)) {
  if (line === '') {
    continue;
  }

  if (line[0] >= '0' && line[0] <= '9') {
    maps[currentMapName].push(line.split(' ').map(Number));
  } else {
    currentMapName = line.split(' ')[0];
    
    maps[currentMapName] = [];
  }
}

function getIntersection(a, b) {
  let min = a[0] < b[0] ? a : b;
  let max = a[0] < b[0] ? b : a;

  return min[1] < max[0]
    ? null
    : [max[0], (min[1] < max[1] ? min[1] : max[1])];
}

function getMappedValue(ranges, mapName) {
  let currentRanges = ranges.slice();
  const newRanges = [];

  for (const [destStart, sourceStart, range] of maps[mapName]) {
    const sourceEnd = sourceStart + range - 1;
    const forRanges = currentRanges.slice();

    for (const [start, end] of forRanges) {
      const intersection = getIntersection([sourceStart, sourceEnd], [start, end]);

      if (intersection) {
        newRanges.push([destStart + (intersection[0] - sourceStart), destStart + (intersection[1] - sourceStart)].sort());
        
        currentRanges = currentRanges.filter((range) => range[0] !== start || range[1] !== end);

        if (intersection[0] - start > 0) {
          currentRanges.push([start, intersection[0] - 1]);
        }

        if (end - intersection[1] > 0) {
          currentRanges.push([intersection[1] + 1, end]);
        }
      }
    }
  }

  return [...newRanges, ...currentRanges];
}

function getSeedLocationRange(seed) {
  const soil = getMappedValue(seed, 'seed-to-soil');
  const fertilizer = getMappedValue(soil, 'soil-to-fertilizer');
  const water = getMappedValue(fertilizer, 'fertilizer-to-water');
  const light = getMappedValue(water, 'water-to-light');
  const temperature = getMappedValue(light, 'light-to-temperature');
  const humidity = getMappedValue(temperature, 'temperature-to-humidity');
  
  return getMappedValue(humidity, 'humidity-to-location');
}

const location = Math.min(...getSeedLocationRange(seeds).reduce((acc, value) => [...acc, ...value], []));

console.log(`Location: ${location}`);
