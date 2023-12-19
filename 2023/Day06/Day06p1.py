# 1 ms hold = 1mm/ms gain

import time
start_time = time.time()
races = []
tot = 1

f = open("Day06\\data.txt")
data = f.read().strip().splitlines()
for x in range(len(data)):
    data[x] = data[x][data[x].find(":")+1:].split()
for x in range(len(data[0])):
    races.append([int(data[0][x]), int(data[1][x])])

for race in races:
    ht = 1
    wins = 0
    while ht <= race[0]:
        d = ht * (race[0] - ht)
        if d > race[1]: wins += 1
        ht += 1
    tot *= wins

print(tot)