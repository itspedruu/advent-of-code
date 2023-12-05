const fs = require('fs');
const lines = fs.readFileSync('./input.txt', 'utf8').split('\r\n');

// name -> (destStart, sourceStart, range)
const maps = {};
let currentMapName = '';

const initialSeeds = lines[0].slice(7).split(' ').map(Number);

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

function getMappedValue(source, mapName) {
  const parameter = maps[mapName].find(([_, start, range]) => source >= start && source < start + range);

  return parameter ? parameter[0] + (source - parameter[1]) : source;
}

function getSeedLocation(seed) {
  const soil = getMappedValue(seed, 'seed-to-soil');
  const fertilizer = getMappedValue(soil, 'soil-to-fertilizer');
  const water = getMappedValue(fertilizer, 'fertilizer-to-water');
  const light = getMappedValue(water, 'water-to-light');
  const temperature = getMappedValue(light, 'light-to-temperature');
  const humidity = getMappedValue(temperature, 'temperature-to-humidity');
  
  return getMappedValue(humidity, 'humidity-to-location');
}

const minLocation = Math.min(...initialSeeds.map(getSeedLocation));

console.log(`Location: ${minLocation}`);
