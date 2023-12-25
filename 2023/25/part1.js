// not sufficient time to optimize stoer-wagner due to being xmas :D
const fs = require('fs');
const lines = fs.readFileSync('./input.txt', 'utf8').trim().split('\r\n');
const graph = {};

for (const line of lines) {
  const [from, part] = line.split(': ');

  for (const to of part.split(' ')) {
    (graph[from] ??= {})[to] = 1;
    (graph[to] ??= {})[from] = 1;
  }
}

const N = Object.keys(graph).length;

const startNode = Object.keys(graph)[0];

function shrinkGraph(g, [s, t]) {
  g[s + t] = {};

  for (const adj of [...Object.keys(g[s]), ...Object.keys(g[t])]) {
    if (adj === s || adj === t) {
      continue;
    }

    const weight = (g[s][adj] ?? 0) + (g[t][adj] ?? 0);

    g[s + t][adj] = weight;
    g[adj][s + t] = weight;

    delete g[adj][s];
    delete g[adj][t];
  }
  
  delete g[s];
  delete g[t];
}

function minCutPhase() {
  let supernode = startNode;
  let g = JSON.parse(JSON.stringify(graph));
  let lastNodes = [];

  while (Object.keys(g).length > 1) {
    let max = 0, nextNode = '';

    for (const node of Object.keys(g[supernode])) {
      if (g[supernode][node] > max) {
        max = g[supernode][node]; 
        nextNode = node;
      }
    }

    if (Object.keys(g).length <= 3) {
      lastNodes.push(nextNode);
    }

    shrinkGraph(g, [supernode, nextNode]);
    supernode += nextNode;
  }

  const cutWeight = Object.values(weight)[0];

  return [lastNodes, cutWeight];  
}

while (Object.keys(graph).length > 1) {
  const [nodePair, cutWeight] = minCutPhase();

  if (cutWeight === 3) {
    const groupOneComponents = nodePair[1].length / 3;
    console.log(groupOneComponents * (N - groupOneComponents));
    break;

  }

  shrinkGraph(graph, nodePair);
}
