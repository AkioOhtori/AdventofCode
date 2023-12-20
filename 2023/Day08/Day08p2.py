import itertools
import time
from math import lcm

start_time = time.time()
f = open("Day08\\data.txt")
full = f.read().strip().splitlines()
f.close()

map = {}
loc = []
steps = 0
directions = full[0].strip()
ans = []

for x in range (2, len(full)):
    t = full[x]

    a = t[:t.find("=")].strip()
    b = t[t.find("(")+1:t.find(",")]
    c = t[t.find(",")+2:t.find(")")]

    map[a] = [b,c]

    if a[2] == "A": loc.append(a)

print(loc)

for z in itertools.cycle(directions):
    ex = 0
    steps += 1
    for x in range(len(loc)):
        if z == "R":
            loc[x] = map[loc[x]][1]
        else:
            loc[x] = map[loc[x]][0]
        
        if loc[x][2] == "Z":
            ans.append(steps)
    if len(ans) == len(loc):
        break
fin = 1
for n in ans:
    fin = lcm(fin, n)
print(ans, fin)
print("\nThis took %s iterations for %s starting locations for a final answer/lcm of %s steps, which took %s seconds to compute.\n" % (steps, len(loc), fin, ("%.5f" % (time.time() - start_time))))