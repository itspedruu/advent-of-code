import z3

def parseFile():
    file = open("10/input.txt")
    lines = file.read().split("\n")[:-1]

    machines = []

    for line in lines:
        parts = line.split(" ")
        machine = {
            'diagram': [],
            'wiring': [],
            'joltages': []
        }

        for char in parts[0][1:-1]:
            if char == '#':
                machine['diagram'].append(1)
            else:
                machine['diagram'].append(0)

        for part in parts[1:-1]:
            machine['wiring'].append([int(rawButton) for rawButton in part[1:-1].split(",")])

        for rawJoltage in parts[-1][1:-1].split(","):
            machine['joltages'].append(int(rawJoltage))

        machines.append(machine)

    return machines

def solveB(machines):
    out = 0

    for machine in machines:
        # initialize Ab
        matrix = [[0] * (len(machine['wiring']) + 1) for _ in machine['diagram']]

        for i in range(len(matrix)):
            matrix[i][len(machine['wiring'])] = machine['joltages'][i]

        for i, wiring in enumerate(machine['wiring']):
            for button in wiring:
                matrix[button][i] = 1

        # solver
        xs = []
        solver = z3.Optimize()

        for i in range(len(machine['wiring'])):
            xs.append(z3.Int(f'X{i}'))
            solver.add(xs[i] >= 0)

        for row in matrix:
            curXs = []

            for i, value in enumerate(row[:-1]):
                if value != 0:
                    curXs.append(xs[i])

            solver.add(sum(curXs) == row[-1])

        solver.minimize(sum(xs))
        assert solver.check()
        model = solver.model()
        declarations = model.decls()

        for xValue in declarations:
            out += model[xValue].as_long()

    return out

machines = parseFile()
solutionB = solveB(machines)

print(solutionB)
