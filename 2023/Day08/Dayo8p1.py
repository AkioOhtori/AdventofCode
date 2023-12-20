import itertools

f = open("Day08\\data.txt")
full = f.read().strip().splitlines()
f.close()

map = {}
steps = 0
directions = full[0].strip()

for x in range (2, len(full)):
    t = full[x]

    a = t[:t.find("=")].strip()
    b = t[t.find("(")+1:t.find(",")]
    c = t[t.find(",")+2:t.find(")")]

    map[a] = [b,c]

loc = "AAA"

for z in itertools.cycle(directions):
    if z == "R":
        loc = map[loc][1]
    else:
        loc = map[loc][0]
    steps += 1
    if loc == "ZZZ":
        break
print(steps)
    